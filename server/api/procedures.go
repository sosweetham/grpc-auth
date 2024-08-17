package api

import (
	"context"

	"github.com/sohamjaiswal/grpc-ftp/api/funcs"
	"github.com/sohamjaiswal/grpc-ftp/pkg/pb"
)

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return funcs.CreateUser(ctx, req)
}

func (s *server) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordResponse, error) {
	return funcs.CheckPassword(ctx, req)
}

func (s *server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return funcs.LoginUser(ctx, req)
}

func (s *server) Me(ctx context.Context, req *pb.NoParam) (*pb.MeResponse, error) {
	return funcs.Me(ctx, req)
}