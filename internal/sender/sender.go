package sender

import (
	"context"
)

func SenderInstance(ctx context.Context) *Sender {
	return ctx.Value(ServiceName).(*Sender)
}

func CtxSenderInstance(entry interface{}, ctx context.Context) context.Context {
	return context.WithValue(ctx, ServiceName, entry)
}

func RunSenderAsInterface(structure interface{}, ctx context.Context) {
	(structure.(*Sender)).Run(ctx)
}
