package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SessionsNotFound() error {
	st := status.New(codes.NotFound, "sessions not found")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_SESSION",
			Domain: "psql.session",
			Metadata: map[string]string{
				"reason": "sessions for the user do not exist in db",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}