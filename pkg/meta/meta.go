package meta

import (
	"context"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/metadata"
)

const (
	AuthKey = "authorization"
)

// Authorization return authorization key string from context.
func Authorization(ctx context.Context) (string, error) {
	return fromMeta(ctx, AuthKey)
}

// fromMeta return string from meta data with specified string key.
func fromMeta(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", xerrors.Errorf("not found meta data")
	}

	vs, ok := md[key]
	if !ok {
		return "", xerrors.Errorf("not found %d in meta", key)
	}
	return vs[0], nil
}
