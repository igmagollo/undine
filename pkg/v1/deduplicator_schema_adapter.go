package undine

type DeduplicatorSchemaAdapter interface {
	InitializeSchemaQuery(tableName string) string
	InsertQuery(tableName string) string
	IsDuplicationError(tableName string, dberr error) bool
}
