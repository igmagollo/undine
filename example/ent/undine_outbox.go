// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill-sql/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/components/forwarder"
	"github.com/ThreeDotsLabs/watermill/message"
)

const undineOutboxSQLTopic = "undine_outbox_messages"

func (tx *Tx) OutboxPublisher() (message.Publisher, error) {
	var publisher message.Publisher
	publisher, err := sql.NewPublisher(
		tx,
		sql.PublisherConfig{
			SchemaAdapter: tx.OutboxSchemaAdapter,
		},
		tx.WatermillLogger,
	)
	if err != nil {
		return nil, fmt.Errorf("creating outbox publisher: %w", err)
	}

	// Decorate publisher so it wraps an event in an envelope understood by the Forwarder component.
	publisher = forwarder.NewPublisher(publisher, forwarder.PublisherConfig{
		ForwarderTopic: undineOutboxSQLTopic,
	})

	return publisher, nil
}

func (c *Client) Forwarder(consumerGroup string) (*forwarder.Forwarder, error) {
	begginer, err := c.DB()
	if err != nil {
		return nil, fmt.Errorf("creating forwarder subscriber: %w", err)
	}

	subscriber, err := sql.NewSubscriber(begginer, sql.SubscriberConfig{
		SchemaAdapter:    c.OutboxSchemaAdapter,
		OffsetsAdapter:   c.OutboxOffsetsAdapter,
		ConsumerGroup:    consumerGroup,
		InitializeSchema: true,
	},
		c.WatermillLogger,
	)
	if err != nil {
		return nil, fmt.Errorf("creating forwarder subscriber: %w", err)
	}

	f, err := forwarder.NewForwarder(subscriber, c.Publisher, c.WatermillLogger, forwarder.Config{
		ForwarderTopic: undineOutboxSQLTopic,
	})
	if err != nil {
		return nil, fmt.Errorf("creating forwarder: %w", err)
	}

	return f, nil
}