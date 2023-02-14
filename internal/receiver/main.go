package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/processor"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"time"
)

const serviceName = data.ModuleName + "-receiver"

type Receiver struct {
	subscriber *amqp.Subscriber
	topic      string
	log        *logan.Entry
	processor  processor.Processor
	responseQ  data.Responses
}

func NewReceiver(cfg config.Config) *Receiver {
	return &Receiver{
		subscriber: cfg.Amqp().Subscriber,
		topic:      cfg.Amqp().Topic,
		log:        logan.New().WithField("service", serviceName),
		processor:  processor.NewProcessor(cfg),
		responseQ:  postgres.NewResponsesQ(cfg.DB()),
	}
}

func (r *Receiver) Run(ctx context.Context) {
	go running.WithBackOff(ctx, r.log,
		serviceName,
		r.listenMessages,
		30*time.Second,
		30*time.Second,
		30*time.Second,
	)
}

func (r *Receiver) listenMessages(ctx context.Context) error {
	r.log.Info("started listening messages")
	r.startSubscriber(ctx, r.topic)
	return nil
}

func (r *Receiver) startSubscriber(ctx context.Context, topic string) {
	go running.WithBackOff(ctx, r.log,
		fmt.Sprint(serviceName, "_", topic),
		func(ctx context.Context) error {
			return r.subscribeForTopic(ctx, topic)
		},
		30*time.Second,
		30*time.Second,
		30*time.Second,
	)
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
	err = r.processor.HandleNewMessage(queueOutput)
	if err != nil {
		responseStatus = "failure"
		errMsg = err.Error()
		r.log.WithError(err).Error("failed to process message ", msg.UUID)
	}

	err = r.responseQ.Insert(data.Response{
		ID:     msg.UUID,
		Status: responseStatus,
		Error:  errMsg,
	})
	if err != nil {
		r.log.WithError(err).Errorf("failed to create response", msg.UUID)
		return errors.Wrap(err, "failed to create response "+msg.UUID)
	}

	r.log.Info("finished processing message ", msg.UUID)
	return nil
}
