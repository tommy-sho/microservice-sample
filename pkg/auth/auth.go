package auth

import (
	"context"

	"github.com/tommy-sho/microservice-sample/pkg/meta"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DefaultAuthFunc is Function to process authentication in common to all services.
type DefaultAuthFunc func(ctx context.Context) (context.Context, error)

func DefaultAuthentication() DefaultAuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		auth, err := meta.Authorization(ctx)
		if err != nil {
			return nil, UnAuthorizationErr
		}

		err = verify(auth)
		if err != nil {
			return nil, UnAuthorizationErr
		}

		return ctx, nil
	}
}

// TODO: implement authenticate func
func verify(auth string) error {
	return nil
}

// UnaryServerInterceptor return unary server interface to process authentication for each request.
func AuthUnaryServerInterceptor(authFunc DefaultAuthFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := authFunc(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return handler(newCtx, req)
	}
}
