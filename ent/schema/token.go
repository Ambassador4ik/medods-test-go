package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("token"),
		field.UUID("access_token_id", uuid.UUID{}),
	}
}

func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tokens").
			Unique(),
	}
}
