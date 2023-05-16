package receiver

import (
	"context"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"time"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/acs-dl/telegram-module-svc/internal/config"
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/data/postgres"
	"github.com/acs-dl/telegram-module-svc/internal/processor"
	"github.com/acs-dl/telegram-module-svc/internal/worker"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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

var handleActions = map[string]func(r *Receiver, msg data.ModulePayload) (string, error){
	AddUserAction: func(r *Receiver, msg data.ModulePayload) (string, error) {
		return r.processor.HandleAddUserAction(msg)
	},
	UpdateUserAction: func(r *Receiver, msg data.ModulePayload) (string, error) {
		return r.processor.HandleUpdateUserAction(msg)
	},
	RemoveUserAction: func(r *Receiver, msg data.ModulePayload) (string, error) {
		return r.processor.HandleRemoveUserAction(msg)
	},
	DeleteUserAction: func(r *Receiver, msg data.ModulePayload) (string, error) {
		return r.processor.HandleDeleteUserAction(msg)
	},
	VerifyUserAction: func(r *Receiver, msg data.ModulePayload) (string, error) {
		return r.processor.HandleVerifyUserAction(msg)
	},
	RefreshModuleAction: func(r *Receiver, msg data.ModulePayload) (string, error) {
		return r.worker.RefreshModule()
	},
	RefreshSubmoduleAction: func(r *Receiver, msg data.ModulePayload) (string, error) {
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

func (r *Receiver) HandleNewMessage(msg data.ModulePayload) (string, error) {
	r.log.Infof("handling message with id `%s`", msg.RequestId)

	err := validation.Errors{
		"action": validation.Validate(msg.Action, validation.Required),
	}.Filter()
	if err != nil {
		r.log.WithError(err).Error("no such action to handle for message with id `%s`", msg.RequestId)
		return data.FAILURE, errors.New("no such action " + msg.Action + " to handle for message with id " + msg.RequestId)
	}

	requestHandler := handleActions[msg.Action]
	requestStatus, err := requestHandler(r, msg)
	if err != nil {
		r.log.WithError(err).Errorf("failed to handle message with id `%s`", msg.RequestId)
		return requestStatus, err
	}

	r.log.Infof("finish handling message with id `%s`", msg.RequestId)
	return requestStatus, nil
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

	var errMsg *string = nil
	requestStatus, err := r.HandleNewMessage(queueOutput)
	if err != nil {
		requestError := err.Error()
		errMsg = &requestError
		r.log.WithError(err).Error("failed to process message ", msg.UUID)
	}

	err = r.responseQ.Insert(data.Response{
		ID:          msg.UUID,
		Status:      requestStatus,
		Error:       errMsg,
		Description: getDescription(requestStatus),
		Payload:     json.RawMessage(msg.Payload),
	})
	if err != nil {
		r.log.WithError(err).Errorf("failed to create response", msg.UUID)
		return errors.Wrap(err, "failed to create response "+msg.UUID)
	}

	r.log.Info("finished processing message ", msg.UUID)
	return nil
}

func getDescription(status string) *string {
	switch status {
	case data.FAILURE:
		return nil
	case data.SUCCESS:
		return nil
	case data.INVITED:
		description := "Message with invite link was sent to user"
		return &description
	default:
		return nil
	}
}
