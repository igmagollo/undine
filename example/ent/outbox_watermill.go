// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type outboxStorerPublisher struct {
	storer *OutboxStorer
	ctx    context.Context // this struct is ephemeral
}

func (*outboxStorerPublisher) Close() error {
	return nil
}

func (o *outboxStorerPublisher) Publish(topic string, messages ...*message.Message) error {
	outboxMessages := make([]*OutboxStorerMessage, len(messages))
	for i, watermillMsg := range messages {
		outboxMessages[i] = &OutboxStorerMessage{
			Topic:   topic,
			Payload: watermillMsg.Payload,
			Headers: watermillMsg.Metadata,
		}
	}

	return o.storer.Store(o.ctx, outboxMessages...)
}

func NewOutboxStorerPublisher(ctx context.Context, storer *OutboxStorer) message.Publisher {
	return &outboxStorerPublisher{
		storer: storer,
		ctx:    ctx,
	}
}
