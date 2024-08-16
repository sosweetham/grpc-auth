package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckPassUserNotFound() error {
	st := status.New(codes.NotFound, "username not found")
		ds, err := st.WithDetails(
			&errdetails.ErrorInfo{
				Reason: "BAD_USERNAME",
				Domain: "psql.users",
				Metadata: map[string]string{
					"reason": "passed username does not exist in db",
				},
			},
		)
		if err != nil {
			return st.Err()
		}
		return ds.Err()
}
