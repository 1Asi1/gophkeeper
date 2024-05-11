package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gophkeeper/internal/server/models"
	"gophkeeper/internal/server/oops"
	"gophkeeper/internal/server/services"
	proto "gophkeeper/rpc/gen"
)

type GophkeeperGrpcService struct {
	proto.UnsafeGophkeeperGrpcServer
	authS services.AuthService
	itemS services.ItemService
}

func NewGophkeeperGrpcService(authS services.AuthService, itemS services.ItemService) *GophkeeperGrpcService {
	return &GophkeeperGrpcService{authS: authS, itemS: itemS}
}

func (s *GophkeeperGrpcService) Register(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	model := models.User{Email: req.Email, Password: req.Password}
	id, err := s.authS.Create(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("s.userS.Create: %w", err)
	}

	expireAt := time.Now().UTC().Add(time.Hour)
	token, err := s.authS.Generate(id, expireAt)
	if err != nil {
		return nil, fmt.Errorf("s.authS.Generate: %w", err)
	}

	return &proto.AuthResponse{
		Token:    token,
		ExpireAt: timestamppb.New(expireAt),
	}, nil
}

func (s *GophkeeperGrpcService) Login(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	user, err := s.authS.Get(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("s.userS.Get: %w", err)
	}

	ok, err := s.authS.ComparePassword(user, req.Password)
	if err != nil {
		return nil, fmt.Errorf("s.userS.ComparePassword: %w", err)
	}
	if !ok {
		return nil, fmt.Errorf("s.userS.ComparePassword: %w", oops.ErrPasswordInvalid)
	}

	expireAt := time.Now().UTC().Add(time.Hour)
	token, err := s.authS.Generate(user.ID, expireAt)
	if err != nil {
		return nil, fmt.Errorf("s.authS.Generate: %w", err)
	}

	return &proto.AuthResponse{
		Token:    token,
		ExpireAt: timestamppb.New(expireAt),
	}, nil
}

func (s *GophkeeperGrpcService) Create(ctx context.Context, req *proto.ItemRequest) (*proto.ItemIDResponse, error) {
	idCtx := ctx.Value(UserID)
	id, err := uuid.Parse(idCtx.(string))
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	model := models.Item{
		UserID: id,
		Data:   req.Request.Data,
		Meta:   req.Request.Meta,
		Type:   req.Request.Type,
	}

	itemID, err := s.itemS.Create(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("s.itemS.Create: %w", err)
	}

	return &proto.ItemIDResponse{
		Id: itemID.NodeID(),
	}, nil
}

func (s *GophkeeperGrpcService) Update(ctx context.Context, req *proto.ItemRequest) (*emptypb.Empty, error) {
	idCtx := ctx.Value(UserID)
	id, err := uuid.Parse(idCtx.(string))
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	model := models.Item{
		UserID: id,
		Data:   req.Request.Data,
		Meta:   req.Request.Meta,
		Type:   req.Request.Type,
	}

	err = s.itemS.Update(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("s.itemS.Update: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *GophkeeperGrpcService) GetAllByType(ctx context.Context, req *proto.GetByTypeRequest) (*proto.GetResponse, error) {
	idCtx := ctx.Value(UserID)
	id, err := uuid.Parse(idCtx.(string))
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	data, err := s.itemS.GetAllByType(ctx, id, req.Type)
	if err != nil {
		return nil, fmt.Errorf("s.itemS.GetAllByType: %w", err)
	}

	res := make([]*proto.Item, len(data))
	for i, v := range data {
		res[i] = &proto.Item{
			Id:   v.ID.NodeID(),
			Data: v.Data,
			Meta: v.Meta,
			Type: v.Type,
		}
	}

	return &proto.GetResponse{Items: res}, nil
}

func (s *GophkeeperGrpcService) Get(ctx context.Context, req *proto.ItemIDRequest) (*proto.Item, error) {
	id, err := uuid.ParseBytes(req.Id)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	data, err := s.itemS.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.itemS.Get: %w", err)
	}

	return &proto.Item{
		Id:   data.ID.NodeID(),
		Type: data.Type,
		Meta: data.Meta,
		Data: data.Data,
	}, nil
}

func (s *GophkeeperGrpcService) GetAll(ctx context.Context, req *proto.ItemIDRequest) (*proto.GetResponse, error) {
	idCtx := ctx.Value(UserID)
	id, err := uuid.Parse(idCtx.(string))
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	data, err := s.itemS.GetAll(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.itemS.GetAll: %w", err)
	}

	res := make([]*proto.Item, len(data))
	for i, v := range data {
		res[i] = &proto.Item{
			Id:   v.ID.NodeID(),
			Data: v.Data,
			Meta: v.Meta,
			Type: v.Type,
		}
	}

	return &proto.GetResponse{Items: res}, nil
}

func (s *GophkeeperGrpcService) Delete(ctx context.Context, req *proto.ItemIDRequest) (*emptypb.Empty, error) {
	id, err := uuid.ParseBytes(req.Id)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	err = s.itemS.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.itemS.Delete: %w", err)
	}

	return &emptypb.Empty{}, nil
}
