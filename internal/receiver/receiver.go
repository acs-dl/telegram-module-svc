package receiver

import (
	"context"
)

func ReceiverInstance(ctx context.Context) *Receiver {
	return ctx.Value(ServiceName).(*Receiver)
}

func CtxReceiverInstance(entry interface{}, ctx context.Context) context.Context {
	return context.WithValue(ctx, ServiceName, entry)
}

func RunReceiverAsInterface(structure interface{}, ctx context.Context) {
	(structure.(*Receiver)).Run(ctx)
}
