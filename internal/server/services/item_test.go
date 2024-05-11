package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gophkeeper/internal/server/models"
	"gophkeeper/internal/server/repository"
)

func TestItem_Create(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
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
			s := &Item{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
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

func TestItem_Delete(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
	}
	type args struct {
		ctx    context.Context
		itemID uuid.UUID
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
			s := &Item{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
			}
			if err := s.Delete(tt.args.ctx, tt.args.itemID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItem_Get(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
	}
	type args struct {
		ctx    context.Context
		itemID uuid.UUID
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
			s := &Item{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
			}
			got, err := s.Get(tt.args.ctx, tt.args.itemID)
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

func TestItem_GetAll(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
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
			s := &Item{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
			}
			got, err := s.GetAll(tt.args.ctx, tt.args.userID)
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

func TestItem_GetAllByType(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
	}
	type args struct {
		ctx      context.Context
		userID   uuid.UUID
		itemType string
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
			s := &Item{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
			}
			got, err := s.GetAllByType(tt.args.ctx, tt.args.userID, tt.args.itemType)
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

func TestItem_Update(t *testing.T) {
	type fields struct {
		log     zerolog.Logger
		storage *repository.Store
		ctx     context.Context
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
			s := &Item{
				log:     tt.fields.log,
				storage: tt.fields.storage,
				ctx:     tt.fields.ctx,
			}
			if err := s.Update(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewItemService(t *testing.T) {
	type args struct {
		ctx     context.Context
		storage *repository.Store
		log     zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want *Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemService(tt.args.ctx, tt.args.storage, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}
