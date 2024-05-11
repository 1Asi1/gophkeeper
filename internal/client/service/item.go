package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gophkeeper/internal/client/models"
	proto "gophkeeper/rpc/gen"
)

const block = 128

type ItemService interface {
	Create(context.Context, models.Item) (uuid.UUID, error)
	Update(context.Context, models.Item) error
	Get(context.Context, uuid.UUID) (models.Item, error)
	GetAll(context.Context, uuid.UUID) ([]models.Item, error)
	GetAllByType(context.Context, uuid.UUID, string) ([]models.Item, error)
	Delete(context.Context, uuid.UUID) error
}

type item struct {
	log    zerolog.Logger
	client proto.GophkeeperGrpcClient
	key    *rsa.PrivateKey
}

func NewItemService(client proto.GophkeeperGrpcClient, key *rsa.PrivateKey, log zerolog.Logger) ItemService {
	return &item{
		client: client,
		log:    log,
		key:    key,
	}
}

func (s *item) Create(ctx context.Context, req models.Item) (uuid.UUID, error) {
	encryptedData, err := encrypt(s.key, req.Data)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("encrypt: %w", err)
	}

	res, err := s.client.Create(ctx, &proto.ItemRequest{
		Request: &proto.Item{
			Type: req.Type,
			Meta: req.Meta,
			Data: encryptedData,
		},
	})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("s.client.Create: %w", err)
	}

	id, err := uuid.ParseBytes(res.GetId())
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("uuid.ParseBytes: %w", err)
	}

	return id, nil
}

func (s *item) Update(ctx context.Context, req models.Item) error {
	encryptedData, err := encrypt(s.key, req.Data)
	if err != nil {
		return fmt.Errorf("encrypt: %w", err)
	}

	_, err = s.client.Update(ctx, &proto.ItemRequest{
		Request: &proto.Item{
			Id:   req.ID.NodeID(),
			Type: req.Type,
			Meta: req.Meta,
			Data: encryptedData,
		},
	})
	if err != nil {
		return fmt.Errorf("s.client.Update: %w", err)
	}

	return err
}

func (s *item) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.client.Delete(ctx, &proto.ItemIDRequest{Id: id.NodeID()})
	if err != nil {
		return fmt.Errorf("s.client.Delete: %w", err)
	}

	return err
}

func (s *item) GetAll(ctx context.Context, id uuid.UUID) ([]models.Item, error) {
	data, err := s.client.GetAll(ctx, &proto.ItemIDRequest{Id: id.NodeID()})
	if err != nil {
		return nil, fmt.Errorf("s.client.GetAll: %w", err)
	}
	res := make([]models.Item, len(data.Items))
	for i, v := range data.Items {
		resID, err := uuid.ParseBytes(v.Id)
		if err != nil {
			return nil, fmt.Errorf("uuid.ParseBytes: %w", err)
		}

		resData, err := decrypt(s.key, v.Data)
		if err != nil {
			return nil, fmt.Errorf("decrypt: %w", err)
		}

		res[i] = models.Item{
			ID:   resID,
			Type: v.Type,
			Data: resData,
			Meta: v.Meta,
		}
	}

	return res, nil
}

func (s *item) GetAllByType(ctx context.Context, id uuid.UUID, types string) ([]models.Item, error) {
	data, err := s.client.GetAllByType(ctx, &proto.GetByTypeRequest{
		Id:   id.NodeID(),
		Type: types,
	})
	if err != nil {
		return nil, fmt.Errorf("s.client.GetAllByType: %w", err)
	}

	res := make([]models.Item, len(data.Items))
	for i, v := range data.Items {
		resID, err := uuid.ParseBytes(v.Id)
		if err != nil {
			return nil, fmt.Errorf("uuid.ParseBytes: %w", err)
		}

		resData, err := decrypt(s.key, v.Data)
		if err != nil {
			return nil, fmt.Errorf("decrypt: %w", err)
		}

		res[i] = models.Item{
			ID:   resID,
			Type: v.Type,
			Data: resData,
			Meta: v.Meta,
		}
	}

	return res, nil
}

func (s *item) Get(ctx context.Context, id uuid.UUID) (models.Item, error) {
	data, err := s.client.Get(ctx, &proto.ItemIDRequest{
		Id: id.NodeID(),
	})
	if err != nil {
		return models.Item{}, fmt.Errorf("s.client.Get: %w", err)
	}

	resID, err := uuid.ParseBytes(data.Id)
	if err != nil {
		return models.Item{}, fmt.Errorf("uuid.ParseBytes: %w", err)
	}

	resData, err := decrypt(s.key, data.Data)
	if err != nil {
		return models.Item{}, fmt.Errorf("decrypt: %w", err)
	}

	return models.Item{
		ID:   resID,
		Type: data.Type,
		Data: resData,
		Meta: data.Meta,
	}, nil
}

func decrypt(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	if privateKey == nil {
		return data, nil
	}
	decryptedData := make([]byte, 0, len(data))
	var nextBlockLength int
	for i := 0; i < len(data); i += privateKey.PublicKey.Size() {
		nextBlockLength = i + privateKey.PublicKey.Size()
		if nextBlockLength > len(data) {
			nextBlockLength = len(data)
		}
		block, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data[i:nextBlockLength], []byte("practicum"))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt data: %v", err)
		}
		decryptedData = append(decryptedData, block...)
	}
	return decryptedData, nil
}

func encrypt(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	if privateKey == nil {
		return data, nil
	}
	encryptedData := make([]byte, 0, len(data))
	var nextBlockLength int
	for i := 0; i < len(data); i += block {
		nextBlockLength = i + block
		if nextBlockLength > len(data) {
			nextBlockLength = len(data)
		}
		block, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, data[i:nextBlockLength], []byte("practicum"))
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt data '%s': %v", data, err)
		}
		encryptedData = append(encryptedData, block...)
	}
	return encryptedData, nil
}
