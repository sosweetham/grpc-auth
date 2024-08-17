package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SessionNotFound() error {
	st := status.New(codes.NotFound, "session not found")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_SESSION",
			Domain: "psql.session",
			Metadata: map[string]string{
				"reason": "passed session does not exist in db",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func BlockedSession() error {
	st := status.New(codes.PermissionDenied, "session is blocked")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_SESSION",
			Domain: "psql.session",
			Metadata: map[string]string{
				"reason": "passed session has been blocked",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func IncorrectSessionUser() error {
	st := status.New(codes.PermissionDenied, "session belongs to someone else")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_SESSION",
			Domain: "psql.session",
			Metadata: map[string]string{
				"reason": "passed session is owned by someone else",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func MismatchedSessionToken() error {
	st := status.New(codes.Internal, "mismatched session")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "SESSION_MISMATCH",
			Domain: "psql.session",
			Metadata: map[string]string{
				"reason": "requested session couldnt be found",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func ExpiredSession() error {
	st := status.New(codes.PermissionDenied, "expired session")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_SESSION",
			Domain: "psql.session",
			Metadata: map[string]string{
				"reason": "passed session is already expired",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}