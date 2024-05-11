package service

import (
	"context"
	"crypto/rsa"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gophkeeper/internal/client/models"
	proto "gophkeeper/rpc/gen"
)

func TestNewItemService(t *testing.T) {
	type args struct {
		client proto.GophkeeperGrpcClient
		key    *rsa.PrivateKey
		log    zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want ItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemService(tt.args.client, tt.args.key, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decrypt(t *testing.T) {
	type args struct {
		privateKey *rsa.PrivateKey
		data       []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decrypt(tt.args.privateKey, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		privateKey *rsa.PrivateKey
		data       []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encrypt(tt.args.privateKey, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_item_Create(t *testing.T) {
	type fields struct {
		log    zerolog.Logger
		client proto.GophkeeperGrpcClient
		key    *rsa.PrivateKey
	}
	type args struct {
		ctx context.Context
		req models.Item
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
			s := &item{
				log:    tt.fields.log,
				client: tt.fields.client,
				key:    tt.fields.key,
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

func Test_item_Delete(t *testing.T) {
	type fields struct {
		log    zerolog.Logger
		client proto.GophkeeperGrpcClient
		key    *rsa.PrivateKey
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &item{
				log:    tt.fields.log,
				client: tt.fields.client,
				key:    tt.fields.key,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_item_Get(t *testing.T) {
	type fields struct {
		log    zerolog.Logger
		client proto.GophkeeperGrpcClient
		key    *rsa.PrivateKey
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &item{
				log:    tt.fields.log,
				client: tt.fields.client,
				key:    tt.fields.key,
			}
			got, err := s.Get(tt.args.ctx, tt.args.id)
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

func Test_item_GetAll(t *testing.T) {
	type fields struct {
		log    zerolog.Logger
		client proto.GophkeeperGrpcClient
		key    *rsa.PrivateKey
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &item{
				log:    tt.fields.log,
				client: tt.fields.client,
				key:    tt.fields.key,
			}
			got, err := s.GetAll(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_item_GetAllByType(t *testing.T) {
	type fields struct {
		log    zerolog.Logger
		client proto.GophkeeperGrpcClient
		key    *rsa.PrivateKey
	}
	type args struct {
		ctx   context.Context
		id    uuid.UUID
		types string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &item{
				log:    tt.fields.log,
				client: tt.fields.client,
				key:    tt.fields.key,
			}
			got, err := s.GetAllByType(tt.args.ctx, tt.args.id, tt.args.types)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllByType() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_item_Update(t *testing.T) {
	type fields struct {
		log    zerolog.Logger
		client proto.GophkeeperGrpcClient
		key    *rsa.PrivateKey
	}
	type args struct {
		ctx context.Context
		req models.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &item{
				log:    tt.fields.log,
				client: tt.fields.client,
				key:    tt.fields.key,
			}
			if err := s.Update(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
