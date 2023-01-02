package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Mixin of the Channel.
func (Channel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("warehouse_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int(10) unsigned", // Override MySQL.
			},
		).Default(0),

		field.String("courier_platform").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(32)", // Override MySQL.
			},
		).Default("").NotEmpty(),

		field.String("name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).NotEmpty(),

		field.String("code").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(64)", // Override MySQL.
			},
		).NotEmpty(),

		field.Int8("type").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(4)", // Override MySQL.
			},
		),

		field.String("quotation_currency").SchemaType(
			map[string]string{
				dialect.MySQL: "char(3)", // Override MySQL.
			},
		).Default("RMB").NotEmpty(),

		field.Int32("volume_factor").SchemaType(
			map[string]string{
				dialect.MySQL: "int(11)", // Override MySQL.
			},
		).Default(0).Comment("材积系数"),

		field.String("en_name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default("").NotEmpty(),

		field.String("display_name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default("").NotEmpty(),

		field.Int8("has_tracking_number").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(4)", // Override MySQL.
			},
		).Default(1),

		field.Int32("min_normal_days").SchemaType(
			map[string]string{
				dialect.MySQL: "int(11)", // Override MySQL.
			},
		).Default(0),

		field.Int32("max_normal_days").SchemaType(
			map[string]string{
				dialect.MySQL: "int(11)", // Override MySQL.
			},
		).Default(0),

		field.Int("max_weight").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("Gram"),

		field.Int("max_length").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0),

		field.Int("min_length").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0),

		field.Int("max_three_side_sum").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0),

		field.String("description").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(512)", // Override MySQL.
			},
		).Default(""),

		field.Int32("sorting_port").SchemaType(
			map[string]string{
				dialect.MySQL: "int(11)", // Override MySQL.
			},
		).Default(0).Comment("流水线出口"),

		field.Bool("prepay_tariff").Default(false),

		field.Int8("status").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(4)", // Override MySQL.
			},
		).Default(0).Comment("1=>启用；0=>禁用"),

		field.Bool("test").Default(false),

		field.String("options").SchemaType(
			map[string]string{
				dialect.MySQL: "json", // Override MySQL.
			},
		).Optional(),

		field.String("exclude_attributes").SchemaType(
			map[string]string{
				dialect.MySQL: "json", // Override MySQL.
			},
		).Optional(),
		
		field.Int8("battery").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(1)", // Override MySQL.
			},
		),
		
		field.Int8("virtual").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(1)", // Override MySQL.
			},
		).Default(1),
		
		field.Int8("channel_type").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(4)", // Override MySQL.
			},
		).Default(1),
		
		field.String("deliver_duty").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(20)", // Override MySQL.
			},
		).Default("ddu"),
		
		field.Int8("special").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(1)", // Override MySQL.
			},
		).Default(0),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("channel_costs", ChannelCost.Type),
		edge.To("orders", Order.Type),
	}
}
