package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateUserUsernameCharUnderLimit() error {
	st := status.New(codes.InvalidArgument, "username char under req limit")
	ds, err := st.WithDetails(
		&errdetails.BadRequest_FieldViolation{
			Field: "username",
			Description: "the username must be more than 3 chars long",
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func CreateUserUsernameAlreadyExists() error {
	st := status.New(codes.AlreadyExists, "same username already exists")
	ds, err := st.WithDetails(
		&errdetails.BadRequest_FieldViolation{
			Field: "username",
			Description: "please choose a different username",
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func CreateUserConfirmationFetchError() error {
	st := status.New(codes.Internal, "couldn't fetch created user")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "DB_FETCH_FAIL",
			Domain: "psql",
			Metadata: map[string]string{
				"message": "couldn't fetch created user, try again later",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}