package api

import (
	"context"

	"github.com/sohamjaiswal/grpc-auth/api/funcs"
	"github.com/sohamjaiswal/grpc-auth/pkg/pb"
)

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return funcs.CreateUser(ctx, req)
}

func (s *server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return funcs.LoginUser(ctx, req)
}

func (s *server) Me(ctx context.Context, req *pb.NoParams) (*pb.MeResponse, error) {
	return funcs.Me(ctx, req)
}

func (s *server) ListSessions(ctx context.Context, req *pb.NoParams) (*pb.ListSessionsResponse, error) {
	return funcs.ListSessions(ctx, req)
}

func (s *server) RenewUserAccessToken(ctx context.Context, req *pb.RenewUserAccessTokenRequest) (*pb.RenewUserAccessTokenResponse, error) {
	return funcs.RenewUserAccessToken(ctx, req)
}

func (s *server) UpdateUserRefreshTokenBlock(ctx context.Context, req *pb.UpdateUserRefreshTokenBlockRequest) (*pb.NoParams, error) {
	return funcs.UpdateUserRefreshTokenBlock(ctx, req)
}
