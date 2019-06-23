package auth

import (
	"google.golang.org/grpc/codes"
)

type Error int

const (
	UnAuthorizationErr Error = iota + 1
)

func (e Error) Error() string {
	switch e {
	case UnAuthorizationErr:
		return "unAuthorization error"
	default:
		return "unknown error"
	}
}

func ToGrpcErr(err Error) codes.Code {
	switch err {
	case UnAuthorizationErr:
		return codes.Unauthenticated
	default:
		return codes.Internal
	}
}
