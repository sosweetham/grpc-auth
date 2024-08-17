package funcs

import (
	"context"

	"github.com/sohamjaiswal/grpc-ftp/api/errors"
	"github.com/sohamjaiswal/grpc-ftp/global"
	"github.com/sohamjaiswal/grpc-ftp/models"
	"github.com/sohamjaiswal/grpc-ftp/pkg/pb"
	"github.com/sohamjaiswal/grpc-ftp/tools"
)

func CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordResponse, error) {
	db, err := global.GetDBConn(false)
	if err != nil {
		return nil, errors.DbConnectionError()
	}
	var checkingUser models.User
	if err = db.Where("username = ?", req.Username).First(&checkingUser).Error; err != nil {
		return nil, errors.UserNotFound()
	}
	comp, err := tools.ComparePasswordAndHash(req.Password, *checkingUser.Password)
	if err != nil {
		return nil, errors.PasswordCompFailure()
	}
	return &pb.CheckPasswordResponse{
		Correct: comp,
	}, nil
}