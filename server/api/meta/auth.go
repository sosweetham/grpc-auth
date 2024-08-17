package meta

import (
	"context"
	"strings"

	"github.com/sohamjaiswal/grpc-ftp/api/errors"
	"github.com/sohamjaiswal/grpc-ftp/global"
	"github.com/sohamjaiswal/grpc-ftp/pkg/payload"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func AuthorizeUser(ctx context.Context) (*payload.AuthPayload, error) {
	md, ok := metadata.FromIncomingContext(ctx); 
	if !ok {
		return nil, errors.MissingMetadata()
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, errors.MissingAuthMetadata()
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, errors.BadAuthMetadataFormat()
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, errors.UnsupportedAuthScheme()
	}

	accessToken := fields[1]
	payload, err := global.GlobalUserAuthTokenizer.VerifyAuthToken(accessToken)
	if err != nil {
		return nil, errors.UserAuthFailure(err)
	}

	return payload, nil
}