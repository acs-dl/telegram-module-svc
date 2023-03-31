package handlers

import (
	"context"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	permissionsCtxKey
	usersCtxKey
	linksCtxKey
	paramsCtxKey
	subsCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxParams(entry *config.GitLabCfg) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, paramsCtxKey, entry)
	}
}

func Params(r *http.Request) *config.GitLabCfg {
	return r.Context().Value(paramsCtxKey).(*config.GitLabCfg)
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

func SubsQ(r *http.Request) data.Subs {
	return r.Context().Value(subsCtxKey).(data.Subs).New()
}

func CtxSubsQ(entry data.Subs) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, subsCtxKey, entry)
	}
}
