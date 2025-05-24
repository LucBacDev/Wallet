package usecase

import (
	"context"
	"source-base-go/golang/proto/auth"

	"google.golang.org/grpc"
)

type Service struct {
	auth.UnimplementedAuthServiceServer
}

func NewService(grpc *grpc.Server) *Service{
	s := &Service{}
	auth.RegisterAuthServiceServer(grpc, s)
	return s
}

func (s *Service) VerifyJWT(ctx context.Context, req *auth.TokenRequest) (*auth.TokenResponse, error) {
    return &auth.TokenResponse{
        IsValid: true,
    }, nil
}
