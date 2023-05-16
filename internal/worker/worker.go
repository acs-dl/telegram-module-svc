package worker

import (
	"context"
)

func WorkerInstance(ctx context.Context) *Worker {
	return ctx.Value(ServiceName).(*Worker)
}

func CtxWorkerInstance(entry interface{}, ctx context.Context) context.Context {
	return context.WithValue(ctx, ServiceName, entry)
}

func RunWorkerAsInterface(structure interface{}, ctx context.Context) {
	(structure.(IWorker)).Run(ctx)
}
