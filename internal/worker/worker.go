package worker

import (
	"context"
)

func WorkerInstance(ctx context.Context) *worker {
	return ctx.Value(ServiceName).(*worker)
}

func CtxWorkerInstance(entry interface{}, ctx context.Context) context.Context {
	return context.WithValue(ctx, ServiceName, entry)
}

func RunWorkerAsInterface(structure interface{}, ctx context.Context) {
	(structure.(Worker)).Run(ctx)
}
