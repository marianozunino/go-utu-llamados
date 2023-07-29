package schema

// import (
// 	"entgo.io/ent"
// 	// "entgo.io/ent/schema/edge"
// 	"entgo.io/ent/schema/edge"
// 	"entgo.io/ent/schema/field"
// )
//
// type SearchResult struct {
// 	ent.Schema
// }
//
// func (SearchResult) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.Int("search_input_id"),
// 		field.Int("job_offer_id"),
// 	}
// }
//
// func (SearchResult) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.From("search_input", SearchInput.Type).Ref("search_results").Field("search_input_id").Required(),
// 		edge.From("job_offer", JobOffer.Type).Ref("search_results").Field("job_offer_id").Required(),
// 	}
// }
