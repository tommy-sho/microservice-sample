package auth

import (
	"context"

	"golang.org/x/xerrors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

/*
example:

https://qiita.com/yoheimuta/items/72d4b75f72d8913adc10

at normal service package
func (*UserService) Authorize(ctx context.Context, fullMethodName string) error {
    switch fullMethodName {
    // 以下の RPC メソッドはすべてのユーザーが行うことができるメソッドなので無条件に許可する {
    case
        "/userpb.UserService/ReportTrouble":
        return nil
    // }
    default:
        return userAuthFunc(ctx)
    }
}

at admin service package
// Authorize はユーザーの認可を行う。
func (*AdminService) Authorize(ctx context.Context, fullMethodName string) error {
    return adminAuthFunc(ctx)
}


at main package
func main() {
    grpcServer := grpc.NewServer(
        grpc.UnaryInterceptor(grpcauthorization.UnaryServerInterceptor()),
    )
...
}


*/

// ServiceAuthorize is an interface that implement function for authorization of each services.
type ServiceAuthorize interface {
	Authorize(context.Context, string) error
}

func CertUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		srv, ok := info.Server.(ServiceAuthorize)
		if !ok {
			return nil, xerrors.Errorf("eaxh service should implement an certification func")
		}

		err := srv.Authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return handler(ctx, req)
	}
}
