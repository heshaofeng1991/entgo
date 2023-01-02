package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// ValueAddedTax holds the schema definition for the ValueAddedTax entity.
type ValueAddedTax struct {
	ent.Schema
}

// Mixin of the ValueAddedTax.
func (ValueAddedTax) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ValueAddedTax.
func (ValueAddedTax) Fields() []ent.Field {
	return []ent.Field{
		field.String("country_code").SchemaType(
			map[string]string{
				dialect.MySQL: "char(2)", // Override MySQL.
			},
		).Default(""),

		field.Float("standard_rate").SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(5,2)", // Override MySQL.
			},
		).Default(0).Comment("标准税率 12 => 12%"),

		field.Float("without_ioss_rate").SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(5,2)", // Override MySQL.
			},
		).Default(0).Comment("没有 IOSS 的税率 14 => 14%"),

		field.Float("exemption_in_usd").SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(15,3)", // Override MySQL.
			},
		).Default(0).Comment("免征额,超出才要计算税费"),
	}
}
