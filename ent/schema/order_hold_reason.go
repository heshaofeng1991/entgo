package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderHoldReason struct {
	ent.Schema
}

func (OrderHoldReason) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Unique(),
		field.Int64("order_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional(),
		field.Int64("product_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional(),
		field.String("type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),
		field.Int32("code").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}),
		field.String("reason").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),
		field.String("en_reason").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

func (OrderHoldReason) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Order.Type).Ref("order_hold_reasons").Field("order_id").Unique(),
	}
}
