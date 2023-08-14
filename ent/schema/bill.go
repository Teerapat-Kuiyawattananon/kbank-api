package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Int("biller_id").Optional(),
		field.Int("ref_1").Optional(),
		field.Int("ref_2").Optional(),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("store", Store.Type).
			Ref("bills").
			Field("biller_id").
			Unique(),

		edge.From("customers", Customer.Type).
			Ref("bills").
			Field("ref_1").
			Unique(),

		edge.To("bill_detail", BillDetail.Type).
			Unique(),
	}
}
