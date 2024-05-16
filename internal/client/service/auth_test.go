package service

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"gophkeeper/internal/client/models"
	proto "gophkeeper/rpc/gen"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		client proto.GophkeeperGrpcClient
		token  *models.Token
		log    zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.client, tt.args.token, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auth_Login(t *testing.T) {
	type fields struct {
		log              zerolog.Logger
		client           proto.GophkeeperGrpcClient
		refreshTokenOnce sync.Once
		token            *models.Token
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
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
			s := &auth{
				log:              tt.fields.log,
				client:           tt.fields.client,
				refreshTokenOnce: tt.fields.refreshTokenOnce,
				token:            tt.fields.token,
			}
			if err := s.Login(tt.args.ctx, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_auth_Register(t *testing.T) {
	type fields struct {
		log              zerolog.Logger
		client           proto.GophkeeperGrpcClient
		refreshTokenOnce sync.Once
		token            *models.Token
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
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
			s := &auth{
				log:              tt.fields.log,
				client:           tt.fields.client,
				refreshTokenOnce: tt.fields.refreshTokenOnce,
				token:            tt.fields.token,
			}
			if err := s.Register(tt.args.ctx, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
