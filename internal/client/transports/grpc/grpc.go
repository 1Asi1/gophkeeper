package grpc

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gophkeeper/internal/client/models"
	"gophkeeper/internal/client/transports/interceptors"
)

func CreateGrpcConnection(targetAddr string, tokenHolder *models.Token) (*grpc.ClientConn, error) {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		return nil, err
	}
	tokenProcessor := interceptors.NewRequestTokenProcessor(tokenHolder)
	return grpc.Dial(
		targetAddr,
		grpc.WithTransportCredentials(tlsCredentials),
		grpc.WithUnaryInterceptor(tokenProcessor.TokenInterceptor()),
		grpc.WithStreamInterceptor(tokenProcessor.TokenStreamInterceptor()),
	)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewClientTLSFromFile("cert/server-cert.pem", "")
	if err != nil {
		return nil, errors.Wrap(err, "tls-error")
	}

	return creds, nil
}
