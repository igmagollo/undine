package undine

import "strings"

type DeduplicatorPostgresSchemaAdapter struct{}

func (a DeduplicatorPostgresSchemaAdapter) InitializeSchemaQuery(tableName string) string {
	return `CREATE TABLE IF NOT EXISTS ` + tableName + ` (
		id UUID not null,
		message_id UUID NOT NULL,
		topic TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT PK_` + tableName + `_id PRIMARY KEY (id),
		CONSTRAINT UQ_` + tableName + `_message_id_topic UNIQUE (message_id, topic)
	)`
}

func (a DeduplicatorPostgresSchemaAdapter) InsertQuery(tableName string) string {
	return `INSERT INTO ` + tableName + ` (id, message_id, topic) VALUES ($1, $2, $3)`
}

func (a DeduplicatorPostgresSchemaAdapter) IsDuplicationError(tableName string, dberr error) bool {
	if dberr == nil {
		return false
	}

	if strings.Contains(dberr.Error(), `uq_`+tableName+`_message_id_topic`) {
		return true
	}

	return false
}
