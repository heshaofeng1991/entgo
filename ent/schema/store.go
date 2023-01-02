package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

// Mixin of the Store.
func (Store) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		field.String("store_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}).Default(""),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}).Default(""),

		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),

		field.String("email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}).Default(""),

		field.String("access_token").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),

		field.Int8("timezone_offset").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0),

		field.String("scope").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),

		field.String("location_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),

		field.String("locations").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional().Comment("店铺的地址列表"),

		field.String("store_currency").SchemaType(map[string]string{
			dialect.MySQL: "varchar(10)", // Override MySQL.
		}).Default(""),

		field.Int8("initial_status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0),

		field.String("state").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),

		field.Int32("timestamp").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),

		field.String("nonce").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("0"),

		field.Float("balance").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,2)", // Override MySQL.
		}).Default(0.00),

		field.String("currency").SchemaType(map[string]string{
			dialect.MySQL: "varchar(12)", // Override MySQL.
		}).Default("USD"),

		field.Float("handling_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,2)", // Override MySQL.
		}).Default(0.99),

		field.Int8("shipping_option").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0),

		field.Int64("default_warehouse").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0),

		field.Int8("prepay_tariff").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0),

		field.String("ioss_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),

		field.String("ioss_country_code").SchemaType(map[string]string{
			dialect.MySQL: "char(2)", // Override MySQL.
		}).Default(""),

		field.String("uk_vat_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),

		field.String("store_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),

		field.String("preset_channel_ids").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional(),

		field.String("test_channel_ids").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional(),

		field.Time("first_inbound_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional().Comment("第一次入库时间"),

		field.Float("storage_unit_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,3)", // Override MySQL.
		}).Default(1.250).Comment("$1.25/day/立方米"),

		field.Time("integration_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional().Comment("集成时间"),
	}
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
		edge.To("platform_products", PlatformProduct.Type),
	}
}
