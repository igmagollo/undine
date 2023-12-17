// Code generated by ent, DO NOT EDIT.

package processedmessage

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the processedmessage type in the database.
	Label = "processed_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMessageID holds the string denoting the message_id field in the database.
	FieldMessageID = "message_id"
	// FieldMessageTopic holds the string denoting the message_topic field in the database.
	FieldMessageTopic = "message_topic"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the processedmessage in the database.
	Table = "processed_messages"
)

// Columns holds all SQL columns for processedmessage fields.
var Columns = []string{
	FieldID,
	FieldMessageID,
	FieldMessageTopic,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the ProcessedMessage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMessageID orders the results by the message_id field.
func ByMessageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessageID, opts...).ToFunc()
}

// ByMessageTopic orders the results by the message_topic field.
func ByMessageTopic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessageTopic, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
