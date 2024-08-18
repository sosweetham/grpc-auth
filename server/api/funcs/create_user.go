package funcs

import (
	"context"

	"github.com/sohamjaiswal/grpc-auth/api/errors"
	"github.com/sohamjaiswal/grpc-auth/global"
	"github.com/sohamjaiswal/grpc-auth/models"
	"github.com/sohamjaiswal/grpc-auth/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	db, err := global.GetDBConn(false)
	if err != nil {
		return nil, err
	}
	if len([]rune(req.Username)) < 3 {
		return nil, errors.CreateUserUsernameCharUnderLimit()
	}
	user := &models.User{
		Username: &req.Username,
		Password: &req.Password,
	}
	err = db.Create(&user).Error
	if err != nil {
		return nil, errors.CreateUserUsernameAlreadyExists()
	}
	var createdUser models.User
	err = db.Where("username = ?", user.Username).First(&createdUser).Error
	if err != nil {
		return nil, errors.CreateUserConfirmationFetchError()
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Username: *createdUser.Username,
			CreatedAt: timestamppb.New(createdUser.CreatedAt),
			UpdatedAt: timestamppb.New(createdUser.UpdatedAt),
		},
	}, nil
}