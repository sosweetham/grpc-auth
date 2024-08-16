package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DbConnectionError() error {
	st := status.New(codes.Internal, "database connection failure")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "DB_CONN_FAIL",
			Domain: "psql",
			Metadata: map[string]string{
				"reason": "backend failed to connect to db",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func PasswordCompFailure() error {
	st := status.New(codes.Internal, "argon compare failure")
		ds, err := st.WithDetails(
			&errdetails.ErrorInfo{
				Reason: "HASH_COMP_FAIL",
				Domain: "tools.argon",
				Metadata: map[string]string{
					"reason": "err occurred while comparing hash",
				},
			},
		)
		if err != nil {
			return st.Err()
		}
		return ds.Err()
}