package auth

import (
	"context"

	sso "github.com/SSSagatov/SaMix/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	sso.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	sso.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *sso.LoginRequest,
) (*sso.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *sso.RegisterRequest,
) (*sso.RegisterResponse, error) {
	panic("implement me")
}
