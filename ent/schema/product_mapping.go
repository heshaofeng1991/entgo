package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductMapping holds the schema definition for the ProductMapping entity.
type ProductMapping struct {
	ent.Schema
}

// Mixin of the ProductMapping.
func (ProductMapping) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ProductMapping.
func (ProductMapping) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("platform_product_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int(10) unsigned", // Override MySQL.
			},
		).Default(0).Optional(),

		field.Int64("product_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int(10) unsigned", // Override MySQL.
			},
		).Default(0).Optional(),

		field.Int("qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		),

		field.Int64("created_by").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		),
	}
}

// Edges of the ProductMapping.
func (ProductMapping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("platform_products", PlatformProduct.Type).
			Ref("product_mappings").Field("platform_product_id").Unique(),
		edge.From("products", Product.Type).
			Ref("product_mappings").Field("product_id").Unique(),
	}
}
