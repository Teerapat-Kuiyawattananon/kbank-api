package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("first_name").Default("Unknown"),
		field.String("last_name").Default(""),
		field.String("title_name").Default(""),
		field.String("mobile_number").Default(""),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("bill_details", BillDetail.Type),
		// edge.From("bill_detail", BillDetail.Type).
		// 	Ref("customer").
		// 	Unique().
		// 	Required(),

		edge.To("bills", Bill.Type),
	}
}
