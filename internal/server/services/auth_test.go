package services

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gophkeeper/internal/server/models"
	"gophkeeper/internal/server/repository"
)

func TestAuth_ComparePassword(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
		Key     string
	}
	type args struct {
		req      models.User
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Auth{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
				Key:     tt.fields.Key,
			}
			got, err := s.ComparePassword(tt.args.req, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComparePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ComparePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_Create(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
		Key     string
	}
	type args struct {
		ctx context.Context
		req models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Auth{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
				Key:     tt.fields.Key,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_Generate(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
		Key     string
	}
	type args struct {
		id       uuid.UUID
		expireAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Auth{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
				Key:     tt.fields.Key,
			}
			got, err := s.Generate(tt.args.id, tt.args.expireAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_Get(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
		Key     string
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Auth{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
				Key:     tt.fields.Key,
			}
			got, err := s.Get(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_GetUser(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
		Key     string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Auth{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
				Key:     tt.fields.Key,
			}
			got, err := s.GetUser(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthService(t *testing.T) {
	type args struct {
		ctx     context.Context
		storage *repository.Store
		key     string
		log     zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want *Auth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.ctx, tt.args.storage, tt.args.key, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}
