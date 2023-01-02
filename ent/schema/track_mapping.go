package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TrackMapping holds the schema definition for the TrackMapping entity.
type TrackMapping struct {
	ent.Schema
}

// Mixin of the TrackMapping.
func (TrackMapping) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the TrackMapping.
func (TrackMapping) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional(),

		field.String("tracking_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.String("tracking_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Default(""),

		field.String("ext_tracking_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.String("track_details").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional(),

		field.Time("last_updated_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),

		field.String("courier_platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),

		field.Int32("status").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}),

		field.Int8("flag").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(1),
	}
}

// Edges of the TrackMapping.
func (TrackMapping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Order.Type).Ref("track_mappings").Field("order_id").Unique(),
	}
}
