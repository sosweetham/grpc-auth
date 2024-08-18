package funcs

import (
	"context"
	"time"

	"github.com/sohamjaiswal/grpc-auth/api/errors"
	"github.com/sohamjaiswal/grpc-auth/global"
	"github.com/sohamjaiswal/grpc-auth/models"
	"github.com/sohamjaiswal/grpc-auth/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RenewUserAccessToken(ctx context.Context, req *pb.RenewUserAccessTokenRequest) (*pb.RenewUserAccessTokenResponse, error) {
	db, err := global.GetDBConn(false)
	if err != nil {
		return nil, errors.DbConnectionError()
	}

	refreshPayload, err := global.GlobalUserAuthTokenizer.VerifyAuthToken(req.RefreshToken)
	if err != nil {
		return nil, errors.UserAuthFailure(err)
	}

	var session models.Session
	if err = db.Where("id = ?", refreshPayload.ID).First(&session).Error; err != nil {
		return nil, errors.SessionNotFound()
	}

	if session.IsBlocked {
		return nil, errors.BlockedSession()
	}

	if *session.Username != refreshPayload.Username {
		return nil, errors.IncorrectSessionUser()
	}

	if *session.RefreshToken != req.RefreshToken {
		return nil, errors.MismatchedSessionToken()
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, errors.ExpiredSession()
	}

	accessToken, accessPayload, err := global.GlobalUserAuthTokenizer.CreateAuthToken(
		refreshPayload.Username,
		global.GlobalUserAccessTokenDuration,
	)
	if err != nil {
		return nil, errors.LoginUserTokenCreationError(err)
	}

	return &pb.RenewUserAccessTokenResponse{
		AccessToken: *accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiresAt),
	}, nil
}