package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoginUserIncorrectPassword() error {
	st := status.New(codes.Unauthenticated, "incorrect password or username")
	ds, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field: "username",
					Description: "doesnt match password",
				},
				{
					Field: "password",
					Description: "doesnt match username",
				},
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func LoginUserTokenCreationError(err error) error {
	st := status.New(codes.Internal, "token issuance error")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "TOKEN_CREATION_ERR",
			Domain: "tokenizer",
			Metadata: map[string]string{
				"error": err.Error(),
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func LoginUserSessionSaveError(err error) error {
	st := status.New(codes.Internal, "session save error")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "SESSION_SAVE_ERR",
			Domain: "tokenizer,psql",
			Metadata: map[string]string{
				"error": err.Error(),
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}
