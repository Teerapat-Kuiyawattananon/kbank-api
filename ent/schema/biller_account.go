package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Biller_account holds the schema definition for the Biller_account entity.
type Biller_account struct {
	ent.Schema
}

// Fields of the Biller_account.
func (Biller_account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").Default(""),
		field.String("service_name").Default(""),
	}
}

// Edges of the Biller_account.
func (Biller_account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bills", Bill.Type),
	}
}
