package errors

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MissingMetadata() error {
	st := status.New(codes.InvalidArgument, "metadata not attached")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "META_UNAVAILABLE",
			Domain: "meta",
			Metadata: map[string]string{
				"reason": "missing metadata",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func MissingAuthMetadata() error {
	st := status.New(codes.InvalidArgument, "auth metadata not attached")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "META_UNAVAILABLE",
			Domain: "meta.auth",
			Metadata: map[string]string{
				"reason": "missing auth metadata",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func BadAuthMetadataFormat() error {
	st := status.New(codes.InvalidArgument, "auth metadata badly formatted")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_META",
			Domain: "meta.auth",
			Metadata: map[string]string{
				"reason": "cannot get auth scheme from meta",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func UnsupportedAuthScheme() error {
	st := status.New(codes.InvalidArgument, "unsupported auth scheme")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_META",
			Domain: "meta.auth",
			Metadata: map[string]string{
				"reason": "called with unsupported auth scheme",
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}

func UserAuthFailure(err error) error {
	st := status.New(codes.Unauthenticated, "bad token")
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: "BAD_META",
			Domain: "meta.auth",
			Metadata: map[string]string{
				"reason": err.Error(),
			},
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}