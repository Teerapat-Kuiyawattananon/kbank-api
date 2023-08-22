package schema

import (
	"time"

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
		field.Int("id"),
		field.Int("biller_id").Optional(),
		field.Int("reference_1").Optional(),
		field.Int("reference_2").Optional(),
		field.String("transaction_id").
			Default(""),
		field.Float("tran_amount").
			Default(0.00),
		field.String("channel_code").
			Default(""),
		field.String("sender_bank_code").
			Default(""),
		field.String("status").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("biller_account", Biller_account.Type).
			Ref("bills").
			Field("biller_id").
			Unique(),

		edge.From("customer", Customer.Type).
			Ref("bills").
			Field("reference_1").
			Unique(),
	}
}
