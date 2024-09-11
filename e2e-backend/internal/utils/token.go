package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const HeaderKey = "Authorization"
const Key = "authorization"

// ToCtx - кладет переданный токен в контекст
func ToCtx(ctx context.Context, token string) context.Context {
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		Key: token,
	}))
	return ctx
}

// FromCtxToRequest - достает токен из контекса
func FromCtx(ctx context.Context) (string, bool) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return "", false
	}
	res := md.Get(Key)
	if len(res) == 0 {
		return "", false
	}
	return res[0], true
}
