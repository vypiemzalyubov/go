package token

import (
	"context"

	"github.com/samber/lo"
	"google.golang.org/grpc/metadata"
)

// nolint
const (
	HeaderKey = "Authorization"
	Key       = "authorization"
)

func Check(ctx context.Context, sessions []string) bool {
	session, ok := FromCtx(ctx)
	if !ok {
		return false
	}

	return lo.Contains(sessions, session)
}

func ToCtx(ctx context.Context, session string) context.Context {
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		Key: session,
	}))
	ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
		Key: session,
	}))
	return ctx
}

func FromCtx(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	res := md.Get(Key)
	if len(res) == 0 {
		return "", false
	}
	return res[0], true
}

func FromCtxToRequest(ctx context.Context) (string, bool) {
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
