package ape

import (
	"context"
	"gitlab.com/distributed_lab/lorem"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	requestIDCtxKey
)

func SetContextLog(ctx context.Context, log *logan.Entry) context.Context {
	return context.WithValue(ctx, logCtxKey, log)
}

func getContextLog(ctx context.Context) *logan.Entry {
	return ctx.Value(logCtxKey).(*logan.Entry)
}

func setRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestIDCtxKey, lorem.ULID())
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(requestIDCtxKey).(string)
}

// Log allows to retrieve request fielded logan.Entry,
// useful only with DefaultMiddlewares.
func Log(r *http.Request) *logan.Entry {
	return getContextLog(r.Context())
}
