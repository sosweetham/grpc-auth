package funcs

import (
	"context"

	"github.com/sohamjaiswal/grpc-ftp/api/errors"
	"github.com/sohamjaiswal/grpc-ftp/api/meta"
	"github.com/sohamjaiswal/grpc-ftp/global"
	"github.com/sohamjaiswal/grpc-ftp/models"
	"github.com/sohamjaiswal/grpc-ftp/pkg/pb"
	"github.com/sohamjaiswal/grpc-ftp/tools"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	db, err := global.GetDBConn(false)
	if err != nil {
		return nil, errors.DbConnectionError()
	}
	var user models.User
	if err = db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, errors.UserNotFound()
	}
	isPassCorrect, err := tools.ComparePasswordAndHash(req.Password, *user.Password)
	if err != nil {
		return nil, errors.PasswordCompFailure()
	}
	if !isPassCorrect {
		return nil, errors.LoginUserIncorrectPassword()
	}
	accessToken, accessPayload, err := global.GlobalUserAuthTokenizer.CreateAuthToken(
		req.Username,
		global.GlobalUserAccessTokenDuration,
	)
	if err != nil {
		return nil, errors.LoginUserTokenCreationError(err)
	}
	refreshToken, refreshPayload, err := global.GlobalUserAuthTokenizer.CreateAuthToken(
		req.Username,
		global.GlobalUserRefreshTokenDuration,
	)
	if err != nil {
		return nil, errors.LoginUserTokenCreationError(err)
	}
	mtdt := meta.ExtractMetadata(ctx)
	session := &models.Session{
		ID: refreshPayload.ID,
		Username: user.Username,
		RefreshToken: refreshToken,
		UserAgent: &mtdt.UserAgent,
		ClientIp: &mtdt.ClientIp,
		ExpiresAt: refreshPayload.ExpiresAt,
	}
	if err = db.Create(&session).Error; err != nil {
		return nil, errors.LoginUserSessionSaveError(err)
	}
	return &pb.LoginUserResponse{
		User: &pb.User{
			Username: *user.Username,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
		SessionId: session.ID.String(),
		AccessToken: *accessToken,
		RefreshToken: *refreshToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
	}, nil
}