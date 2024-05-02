package repository

import (
	"embed"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

const (
	countStep      = 2
	retryStopCount = 5
)

var (
	count = 1
)

// Config структура с полями для подключения к базе данных.
type Config struct {
	// строка подключения с базой данных.
	ConnDSN string
	// максимальное количество открытых соединений с базой данных.
	MaxConn int
	// максимальное количество времени, в течение которого соединение может быть использовано повторно.
	MaxConnLifeTime time.Duration
	// максимальное количество времени, в течение которого соединение может простаивать.
	MaxConnIdleTime time.Duration
	// логгер.
	Logger zerolog.Logger
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

// Store структура базы данных.
type Store struct {
	db  *sqlx.DB
	log zerolog.Logger
}

// New инициализация подключения к базе данных.
func New(cfg Config, log zerolog.Logger) (*Store, error) {
	l := cfg.Logger.With().Str("postgres", "New").Logger()
	var db *sqlx.DB
	var err error
	for ; ; count += countStep {
		ticker := time.NewTicker(time.Duration(count) * time.Second)
		db, err = sqlx.Connect("pgx", cfg.ConnDSN)
		if err != nil {
			pgErr, ok := (err).(pgx.PgError)
			if !ok {
				return &Store{}, fmt.Errorf("postgres connection error: %w", err)
			}

			if pgErr.Code == pgerrcode.InvalidAuthorizationSpecification {
				l.Info().Msgf("try connection sec: %d", count)
				<-ticker.C
				l.Err(err).Msg("sqlx.Connect try agan...")
				if count == retryStopCount {
					l.Error().Msg("sqlx.Connect try cancel")
					return &Store{}, fmt.Errorf("postgres connection error: %w", err)
				}
				continue
			}

			return &Store{}, fmt.Errorf("postgres connection error: %w", err)
		}
		break
	}

	l.Info().Msg("succeeded in connecting to postgres")

	db.SetConnMaxLifetime(cfg.MaxConnLifeTime)
	db.SetConnMaxIdleTime(cfg.MaxConnIdleTime)
	db.SetMaxOpenConns(cfg.MaxConn)

	if err = runMigrations(cfg.ConnDSN); err != nil {
		return &Store{}, fmt.Errorf("runMigrations error: %w", err)
	}

	return &Store{db: db, log: log}, nil
}

func runMigrations(dsn string) error {
	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return fmt.Errorf("failed to return an iofs driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return fmt.Errorf("failed to get a new migrate instance: %w", err)
	}
	if err = m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations to the DB: %w", err)
		}
	}
	return nil
}
