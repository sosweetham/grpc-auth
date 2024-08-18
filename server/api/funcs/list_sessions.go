package funcs

import (
	"context"

	"github.com/sohamjaiswal/grpc-auth/api/errors"
	"github.com/sohamjaiswal/grpc-auth/api/meta"
	"github.com/sohamjaiswal/grpc-auth/global"
	"github.com/sohamjaiswal/grpc-auth/models"
	"github.com/sohamjaiswal/grpc-auth/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ListSessions(ctx context.Context, req *pb.NoParams) (*pb.ListSessionsResponse, error) {
	userAuth, err := meta.AuthorizeUser(ctx)
	if err != nil {
		return nil, err
	}
	db, err := global.GetDBConn(false)
	if err != nil {
		return nil, errors.DbConnectionError()
	}
	var sessions []models.Session
	if err = db.Where("username = ?", userAuth.Username).Find(&sessions).Error; err != nil {
		return nil, errors.SessionsNotFound()
	}
	resSessions := []*pb.Session{}
	for _, currSession := range sessions {
		resSessions = append(resSessions, &pb.Session{
			Id: currSession.ID.String(),
			UserAgent: *currSession.UserAgent,
			ExpiresAt: timestamppb.New(currSession.ExpiresAt),
			IsBlocked: currSession.IsBlocked,
		})
	}
	return &pb.ListSessionsResponse{
		Sessions: resSessions,
	}, nil
}