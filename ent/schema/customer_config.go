package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CustomerConfig holds the schema definition for the CustomerConfig entity.
type CustomerConfig struct {
	ent.Schema
}

// Mixin of the CustomerConfig.
func (CustomerConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the CustomerConfig.
func (CustomerConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10)unsigned", // Override MySQL.
		}),
		field.String("exclude_country_code").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}),
		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("0=>Inactive, 1=>Active"),
	}
}

// Edges of the CustomerConfig.
func (CustomerConfig) Edges() []ent.Edge {
	return nil
}