package interceptors

import (
	"context"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gophkeeper/internal/client/models"
	"gophkeeper/internal/logger"
)

type RequestTokenProcessor interface {
	TokenInterceptor() grpc.UnaryClientInterceptor
	TokenStreamInterceptor() grpc.StreamClientInterceptor
}

type requestTokenProcessor struct {
	log         zerolog.Logger
	tokenHolder *models.Token
}

func NewRequestTokenProcessor(tokenHolder *models.Token) RequestTokenProcessor {
	return &requestTokenProcessor{
		log:         logger.NewLogger(),
		tokenHolder: tokenHolder,
	}
}

func (tp *requestTokenProcessor) TokenInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(tp.ctxWithToken(ctx), method, req, reply, cc, opts...)
	}
}

func (tp *requestTokenProcessor) TokenStreamInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		return streamer(tp.ctxWithToken(ctx), desc, cc, method, opts...)
	}
}

func (tp *requestTokenProcessor) ctxWithToken(ctx context.Context) context.Context {
	token := tp.tokenHolder.Get()
	if token != "" {
		return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"token": token}))
	}
	return ctx
}
