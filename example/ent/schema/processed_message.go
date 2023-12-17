package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type ProcessedMessage struct {
	ent.Schema
}

func (ProcessedMessage) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}).Immutable(), field.UUID("message_id", uuid.UUID{}).Immutable(), field.String("message_topic").Immutable(), field.Time("created_at")}
}
func (ProcessedMessage) Edges() []ent.Edge {
	return nil
}
func (ProcessedMessage) Annotations() []schema.Annotation {
	return nil
}
func (ProcessedMessage) Indexes() []ent.Index {
	return []ent.Index{index.Fields("message_id", "message_topic").Unique()}
}
