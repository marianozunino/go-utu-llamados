package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type JobOffer struct {
	ent.Schema
}

func (JobOffer) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description"),
		field.String("url"),
		field.String("file_url"),
		field.Time("published_at"),
		field.Time("updated_at"),
	}
}

func (JobOffer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("search_inputs", SearchInput.Type).Ref("search_results"),
	}
}
