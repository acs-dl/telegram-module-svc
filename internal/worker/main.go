package worker

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/processor"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const serviceName = data.ModuleName + "-worker"

type Worker interface {
	Run(ctx context.Context)
}

type worker struct {
	logger    *logan.Entry
	processor processor.Processor
	linksQ    data.Links
}

func NewWorker(cfg config.Config, ctx context.Context) Worker {
	return &worker{
		logger:    cfg.Log().WithField("runner", serviceName),
		processor: processor.NewProcessor(cfg, ctx),
		linksQ:    postgres.NewLinksQ(cfg.DB()),
	}
}

func (w *worker) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		w.logger,
		serviceName,
		w.processPermissions,
		15*time.Minute,
		15*time.Minute,
		15*time.Minute,
	)
}

func (w *worker) processPermissions(_ context.Context) error {
	w.logger.Info("fetching links")

	links, err := w.linksQ.Select()
	if err != nil {
		return errors.Wrap(err, "failed to get links")
	}

	reqAmount := len(links)
	if reqAmount == 0 {
		w.logger.Info("no links were found")
		return nil
	}

	w.logger.Infof("found %v links", reqAmount)

	for _, link := range links {
		w.logger.Infof("processing link `%s`", link.Link)

		err = w.createPermissions(link.Link)
		if err != nil {
			w.logger.Infof("failed to create permissions for chat")
			return errors.Wrap(err, "failed to create permissions for chat")
		}

		w.logger.Infof("successfully processed link `%s`", link.Link)
	}

	return nil
}

func (w *worker) createPermissions(link string) error {
	if err := w.processor.HandleNewMessage(data.ModulePayload{
		RequestId: "from-worker",
		Action:    processor.GetUsersAction,
		Link:      link,
	}); err != nil {
		w.logger.Infof("failed to get users for link `%s`", link)
		return errors.Wrap(err, "failed to get users")
	}

	return nil
}
