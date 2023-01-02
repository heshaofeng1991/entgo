package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderTaxation struct {
	ent.Schema
}

// Mixin of the OrderTaxation.
func (OrderTaxation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (OrderTaxation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional(),
		field.String("type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(8)", // Override MySQL.
		}),
		field.String("country_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(2)", // Override MySQL.
		}).Optional().Default(""),
		field.String("number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}),
	}
}

func (OrderTaxation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Order.Type).Ref("order_taxations").Field("order_id").Unique(),
	}
}
