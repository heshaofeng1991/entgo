package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// ChannelCostBatche holds the schema definition for the Channel Cost Batches entity.
type ChannelCostBatche struct {
	ent.Schema
}

// Fields of the ChannelCostBatche.
func (ChannelCostBatche) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Unique(),

		field.Time("created_at").Immutable().Default(time.Now),

		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Positive(),

		field.Time("effective_date").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),

		field.Time("expiry_date").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),

		field.Bool("status").Default(true),
	}
}

// Edges of the ChannelCostBatche.
func (ChannelCostBatche) Edges() []ent.Edge {
	return nil
}
