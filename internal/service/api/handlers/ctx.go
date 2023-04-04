package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	permissionsCtxKey
	usersCtxKey
	linksCtxKey
	paramsCtxKey
	pqueueCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxParams(entry *config.TelegramCfg) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, paramsCtxKey, entry)
	}
}

func Params(r *http.Request) *config.TelegramCfg {
	return r.Context().Value(paramsCtxKey).(*config.TelegramCfg)
}

func PermissionsQ(r *http.Request) data.Permissions {
	return r.Context().Value(permissionsCtxKey).(data.Permissions).New()
}

func CtxPermissionsQ(entry data.Permissions) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, permissionsCtxKey, entry)
	}
}

func UsersQ(r *http.Request) data.Users {
	return r.Context().Value(usersCtxKey).(data.Users).New()
}

func CtxUsersQ(entry data.Users) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersCtxKey, entry)
	}
}

func LinksQ(r *http.Request) data.Links {
	return r.Context().Value(linksCtxKey).(data.Links).New()
}

func CtxLinksQ(entry data.Links) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, linksCtxKey, entry)
	}
}

func PQueue(ctx context.Context) *pqueue.PriorityQueue {
	return ctx.Value(pqueueCtxKey).(*pqueue.PriorityQueue)
}

func CtxPQueue(entry *pqueue.PriorityQueue, ctx context.Context) context.Context {
	return context.WithValue(ctx, pqueueCtxKey, entry)
}
