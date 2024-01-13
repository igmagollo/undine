package undine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type DuplicationError struct {
	Err error
}

func (e DuplicationError) Error() string {
	return e.Err.Error()
}

func (e DuplicationError) Unwrap() error {
	return e.Err
}

func IsDuplicationError(err error) bool {
	var dupErr *DuplicationError
	if errors.As(err, &dupErr) {
		return true
	}
	return false
}

type Deduplicator struct {
	db            ContextExecutor
	tableName     string
	schemaAdapter DeduplicatorSchemaAdapter
}

func (d *Deduplicator) InitializeSchema(ctx context.Context) error {
	if isTx(d.db) {
		return errors.New("cannot initialize schema in a transaction")
	}

	query := d.schemaAdapter.InitializeSchemaQuery(d.tableName)
	if _, err := d.db.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}

func (d *Deduplicator) Deduplicate(ctx context.Context, topic string, id uuid.UUID) error {
	query := d.schemaAdapter.InsertQuery(d.tableName)
	_, err := d.db.ExecContext(ctx, query, uuid.New(), id, topic)
	if err != nil {
		if d.schemaAdapter.IsDuplicationError(d.tableName, err) {
			return &DuplicationError{Err: err}
		}

		return fmt.Errorf("inserting messaga into deduplication table: %w", err)
	}

	return nil
}

func NewDeduplicator(db ContextExecutor, tableName string, schemaAdapter DeduplicatorSchemaAdapter) *Deduplicator {
	return &Deduplicator{
		db:            db,
		tableName:     tableName,
		schemaAdapter: schemaAdapter,
	}
}

type ContextExecutor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func isTx(db ContextExecutor) bool {
	_, dbIsTx := db.(interface {
		Commit() error
		Rollback() error
	})
	return dbIsTx
}
