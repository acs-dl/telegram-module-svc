package config

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AmqpConfig struct {
	Topic      string `fig:"topic,required"`
	Publisher  string `fig:"publisher,required"`
	Subscriber string `fig:"subscriber,required"`
}

type AmqpData struct {
	Topic      string
	Publisher  *amqp.Publisher
	Subscriber *amqp.Subscriber
}

func (c *config) Amqp() *AmqpData {
	return c.amqp.Do(func() interface{} {
		var cfg AmqpConfig

		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "amqp")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out publisher config"))
		}

		return &AmqpData{
			Topic:      cfg.Topic,
			Subscriber: createSubscriber(cfg.Subscriber),
			Publisher:  createPublisher(cfg.Publisher),
		}
	}).(*AmqpData)
}

func createSubscriber(url string) *amqp.Subscriber {
	amqpConfig := amqp.NewDurablePubSubConfig(url, amqp.GenerateQueueNameTopicNameWithSuffix("SUFFIX"))
	watermillLogger := watermill.NewStdLogger(false, false)

	subscriber, err := amqp.NewSubscriber(amqpConfig, watermillLogger)
	if err != nil {
		panic(errors.Wrap(err, "failed to create subscriber"))
	}

	return subscriber
}

func createPublisher(url string) *amqp.Publisher {
	amqpConfig := amqp.NewDurablePubSubConfig(url, nil)
	watermillLogger := watermill.NewStdLogger(false, false)

	publisher, err := amqp.NewPublisher(amqpConfig, watermillLogger)
	if err != nil {
		panic(errors.Wrap(err, "failed to create publisher"))
	}

	return publisher
}
