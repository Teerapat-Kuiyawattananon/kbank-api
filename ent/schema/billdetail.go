package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BillDetail holds the schema definition for the BillDetail entity.
type BillDetail struct {
	ent.Schema
}

// Fields of the BillDetail.
func (BillDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("channel_code").
			Default(""),
		field.String("sender_bank_code").
			Default(""),
		field.String("status").
			NotEmpty(),
		field.Int("customer_id").
			Optional(),
		field.Float("tran_amount").
			Default(0.00),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the BillDetail.
func (BillDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bills", Bill.Type).
			Ref("bill_detail").
			Unique(),
			// Required(),
		
		edge.From("customer", Customer.Type).
			Ref("bill_details").
			Field("customer_id").
			Unique(),
	}
}
