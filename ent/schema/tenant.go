package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Mixin of the Tenant schema.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).Default(""),

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

		field.Int64("default_warehouse").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0),

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

		field.Int64("cs_user_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional().Comment("客服用户 ID"),

		field.Int64("sales_user_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional().Comment("销售用户 ID"),

		field.Int64("inviter_user_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional().Comment("推荐人客户 ID"),

		field.String("platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Comment("平台"),
	}
}

// Policy defines the privacy policy of the User.
func (Tenant) Policy() ent.Policy {
	return nil
}
