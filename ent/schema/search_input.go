package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SearchInput struct {
	ent.Schema
}

func (SearchInput) Fields() []ent.Field {
	return []ent.Field{
		field.String("search"),
	}
}

func (SearchInput) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("search_results", JobOffer.Type),
	}
}
