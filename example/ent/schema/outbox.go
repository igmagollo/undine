package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Outbox struct {
	ent.Schema
}

func (Outbox) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}).Immutable(), field.String("topic").Immutable(), field.Bytes("payload").Immutable(), field.JSON("headers", map[string]string{}).Optional(), field.Time("created_at")}
}
func (Outbox) Edges() []ent.Edge {
	return nil
}
func (Outbox) Annotations() []schema.Annotation {
	return nil
}
func (Outbox) Indexes() []ent.Index {
	return []ent.Index{index.Fields("created_at", "topic")}
}
