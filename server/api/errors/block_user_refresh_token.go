package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CouldntBlockSession() error {
	st := status.New(codes.Internal, "couldnt block session")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "DB_ERR",
			Domain: "psql.session",
			Metadata: map[string]string{
				"reason": "couldnt block session",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}