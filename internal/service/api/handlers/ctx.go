package handlers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/config"
	"github.com/acs-dl/telegram-module-svc/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	permissionsCtxKey
	usersCtxKey
	chatsCtxKey
	linksCtxKey
	configCtxKey
	parentContextCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
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

func ChatsQ(r *http.Request) data.Chats {
	return r.Context().Value(chatsCtxKey).(data.Chats).New()
}

func CtxChatsQ(entry data.Chats) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, chatsCtxKey, entry)
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

func Config(ctx context.Context) config.Config {
	return ctx.Value(configCtxKey).(config.Config)
}

func CtxConfig(entry config.Config, ctx context.Context) context.Context {
	return context.WithValue(ctx, configCtxKey, entry)
}

func ParentContext(ctx context.Context) context.Context {
	return ctx.Value(parentContextCtxKey).(context.Context)
}

func CtxParentContext(entry context.Context) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, parentContextCtxKey, entry)
	}
}
