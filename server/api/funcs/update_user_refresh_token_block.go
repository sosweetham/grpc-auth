package funcs

import (
	"context"

	"github.com/sohamjaiswal/grpc-ftp/api/errors"
	"github.com/sohamjaiswal/grpc-ftp/api/meta"
	"github.com/sohamjaiswal/grpc-ftp/global"
	"github.com/sohamjaiswal/grpc-ftp/models"
	"github.com/sohamjaiswal/grpc-ftp/pkg/pb"
)

func UpdateUserRefreshTokenBlock(ctx context.Context, req *pb.UpdateUserRefreshTokenBlockRequest) (*pb.NoParams, error) {
	userAuth, err := meta.AuthorizeUser(ctx)
	if err != nil {
		return nil, err
	}
	db, err := global.GetDBConn(false)
	if err != nil {
		return nil, errors.DbConnectionError()
	}
	refreshPayload, err := global.GlobalUserAuthTokenizer.VerifyAuthToken(req.RefreshToken)
	if err != nil {
		return nil, errors.UserAuthFailure(err)
	}
	var session models.Session
	if err = db.Where("username = ? AND id = ?", userAuth.Username, refreshPayload.ID).First(&session).Error; err != nil {
		return nil, errors.SessionNotFound()
	}
	if err = db.Model(&session).Update("is_blocked", req.Block).Error; err != nil {
		return nil, errors.CouldntBlockSession()
	}
	return &pb.NoParams{}, nil
}