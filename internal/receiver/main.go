package receiver

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/processor"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/worker"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const (
	ServiceName = data.ModuleName + "-receiver"

	AddUserAction    = "add_user"
	UpdateUserAction = "update_user"
	RemoveUserAction = "remove_user"
	VerifyUserAction = "verify_user"
	DeleteUserAction = "delete_user"

	RefreshModuleAction    = "refresh_module"
	RefreshSubmoduleAction = "refresh_submodule"
)

type Receiver struct {
	subscriber  *amqp.Subscriber
	topic       string
	log         *logan.Entry
	processor   processor.Processor
	worker      *worker.Worker
	responseQ   data.Responses
	runnerDelay time.Duration
}

var handleActions = map[string]func(r *Receiver, msg data.ModulePayload) error{
	AddUserAction: func(r *Receiver, msg data.ModulePayload) error {
		return r.processor.HandleAddUserAction(msg)
	},
	UpdateUserAction: func(r *Receiver, msg data.ModulePayload) error {
		return r.processor.HandleUpdateUserAction(msg)
	},
	RemoveUserAction: func(r *Receiver, msg data.ModulePayload) error {
		return r.processor.HandleRemoveUserAction(msg)
	},
	DeleteUserAction: func(r *Receiver, msg data.ModulePayload) error {
		return r.processor.HandleDeleteUserAction(msg)
	},
	VerifyUserAction: func(r *Receiver, msg data.ModulePayload) error {
		return r.processor.HandleVerifyUserAction(msg)
	},
	RefreshModuleAction: func(r *Receiver, msg data.ModulePayload) error {
		return r.worker.ProcessPermissions(context.Background())
	},
	RefreshSubmoduleAction: func(r *Receiver, msg data.ModulePayload) error {
		return r.worker.RefreshSubmodules(msg)
	},
}

func NewReceiverAsInterface(cfg config.Config, ctx context.Context) interface{} {
	return interface{}(&Receiver{
		subscriber:  cfg.Amqp().Subscriber,
		topic:       cfg.Amqp().Topic,
		log:         logan.New().WithField("service", ServiceName),
		processor:   processor.ProcessorInstance(ctx),
		responseQ:   postgres.NewResponsesQ(cfg.DB()),
		worker:      worker.WorkerInstance(ctx),
		runnerDelay: cfg.Runners().Receiver,
	})
}

func (r *Receiver) Run(ctx context.Context) {
	go running.WithBackOff(ctx, r.log,
		ServiceName,
		r.listenMessages,
		r.runnerDelay,
		r.runnerDelay,
		r.runnerDelay,
	)
}

func (r *Receiver) listenMessages(ctx context.Context) error {
	r.log.Info("started listening messages")
	return r.subscribeForTopic(ctx, r.topic)
}

func (r *Receiver) subscribeForTopic(ctx context.Context, topic string) error {
	msgChan, err := r.subscriber.Subscribe(ctx, topic)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe for topic "+topic)
	}
	r.log.Info("successfully subscribed for topic ", topic)

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgChan:
			r.log.Info("received message ", msg.UUID)
			err = r.processMessage(msg)
			if err != nil {
				r.log.WithError(err).Error("failed to process message ", msg.UUID)
			} else {
				msg.Ack()
			}
		}
	}
}

func (r *Receiver) HandleNewMessage(msg data.ModulePayload) error {
	r.log.Infof("handling message with id `%s`", msg.RequestId)

	err := validation.Errors{
		"action": validation.Validate(msg.Action, validation.Required),
	}.Filter()
	if err != nil {
		r.log.WithError(err).Error("no such action to handle for message with id `%s`", msg.RequestId)
		return errors.New("no such action " + msg.Action + " to handle for message with id " + msg.RequestId)
	}

	requestHandler := handleActions[msg.Action]
	if err = requestHandler(r, msg); err != nil {
		r.log.WithError(err).Errorf("failed to handle message with id `%s`", msg.RequestId)
		return err
	}

	r.log.Infof("finish handling message with id `%s`", msg.RequestId)
	return nil
}

func (r *Receiver) processMessage(msg *message.Message) error {
	r.log.Info("started processing message ", msg.UUID)

	var queueOutput data.ModulePayload
	err := json.Unmarshal(msg.Payload, &queueOutput)
	if err != nil {
		r.log.WithError(err).Errorf("failed to unmarshal message", msg.UUID)
		return errors.Wrap(err, "failed to unmarshal message "+msg.UUID)
	}
	queueOutput.RequestId = msg.UUID

	var responseStatus = "success"
	var errMsg = ""
	err = r.HandleNewMessage(queueOutput)
	if err != nil {
		responseStatus = "failure"
		errMsg = err.Error()
		r.log.WithError(err).Error("failed to process message ", msg.UUID)
	}

	err = r.responseQ.Insert(data.Response{
		ID:      msg.UUID,
		Status:  responseStatus,
		Error:   errMsg,
		Payload: json.RawMessage(msg.Payload),
	})
	if err != nil {
		r.log.WithError(err).Errorf("failed to create response", msg.UUID)
		return errors.Wrap(err, "failed to create response "+msg.UUID)
	}

	r.log.Info("finished processing message ", msg.UUID)
	return nil
}
