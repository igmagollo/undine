// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	undine "github.com/igmagollo/undine/pkg/v1"
)

const tableName = "undine_deduplicator_messages"

func (tx *Tx) Deduplicator() *undine.Deduplicator {
	return undine.NewDeduplicator(tx, tableName, tx.DeduplicatorSchemaAdapter)
}

func (c *Client) InitializeDeduplicatorSchema(ctx context.Context) error {
	dedup := undine.NewDeduplicator(c, tableName, c.DeduplicatorSchemaAdapter)

	return dedup.InitializeSchema(ctx)
}

func DeduplicatorFromContext(ctx context.Context) *undine.Deduplicator {
	tx := TxFromContext(ctx)
	if tx == nil {
		return nil
	}

	return tx.Deduplicator()
}
