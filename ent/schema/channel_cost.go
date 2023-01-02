package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChannelCost holds the schema definition for the Channel Cost entity.
type ChannelCost struct {
	ent.Schema
}

// Mixin of the ChannelCost.
func (ChannelCost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ChannelCost.
func (ChannelCost) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("channel_cost_batch_id").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Positive(),

		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Positive().Optional(),

		field.Int8("mode").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("报价计算模型"),

		field.String("country_code").SchemaType(map[string]string{
			dialect.MySQL: "char(2)", // Override MySQL.
		}).NotEmpty(),

		field.String("zone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),

		field.Int("start_weight").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Comment("gram"),

		field.Int("end_weight").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Comment("gram"),

		field.Int("first_weight").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Default(0).Comment("gram"),

		field.Float("first_weight_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00),

		field.Int("unit_weight").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Comment("单位重量"),

		field.Float("unit_weight_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,5)", // Override MySQL.
		}).Optional().Default(0.00000),

		field.Float("fuel_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00),

		field.Float("processing_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00),

		field.Float("registration_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00),

		field.Float("misc_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)", // Override MySQL.
		}).Default(0.00),
		
		field.Int("min_normal_days").SchemaType(
			map[string]string{
				dialect.MySQL: "int(11)", // Override MySQL.
			},
		).Default(0),
		
		field.Int("max_normal_days").SchemaType(
			map[string]string{
				dialect.MySQL: "int(11)", // Override MySQL.
			},
		).Default(0),
		
		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("0=>启用, 1=>不启用"),
		
		field.Int("average_days").SchemaType(
			map[string]string{
				dialect.MySQL: "int(11)", // Override MySQL.
			},
		).Default(0),
	}
}

// Edges of the ChannelCost.
func (ChannelCost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("channels", Channel.Type).Ref("channel_costs").Field("channel_id").Unique(),
	}
}
