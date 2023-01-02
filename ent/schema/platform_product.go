/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    platform_product
	@Date    2022/3/19 7:20 下午
	@Desc
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlatformProduct holds the schema definition for the PlatformProduct entity.
type PlatformProduct struct {
	ent.Schema
}

// Mixin of the PlatformProduct.
func (PlatformProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the PlatformProduct.
func (PlatformProduct) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("store_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Default(0).Optional(),

		field.String("listing_id").SchemaType(map[string]string{
			dialect.MySQL: "char(256)", // Override MySQL.
		}).NotEmpty(),

		field.String("collection_id").SchemaType(map[string]string{
			dialect.MySQL: "char(256)", // Override MySQL.
		}).Default(""),

		field.String("listing_sku").SchemaType(map[string]string{
			dialect.MySQL: "char(256)", // Override MySQL.
		}).Default(""),

		field.String("barcode").SchemaType(map[string]string{
			dialect.MySQL: "char(255)", // Override MySQL.
		}).Default(""),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "char(512)", // Override MySQL.
		}).Default(""),

		field.String("images").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional(),

		field.String("vendor").SchemaType(map[string]string{
			dialect.MySQL: "char(128)", // Override MySQL.
		}).Optional(),

		field.Float("selling_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000).Comment("售卖价 货币单位跟随store_currency"),

		field.String("currency").SchemaType(map[string]string{
			dialect.MySQL: "char(3)", // Override MySQL.
		}).Default(""),

		field.Int("grams").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,0)", // Override MySQL.
		}).Default(0),

		field.Int8("platform_status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Default(1).Comment("1=>使用中；0=>删除"),
	}
}

// Edges of the PlatformProduct.
func (PlatformProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_mappings", ProductMapping.Type),
		edge.From("stores", Store.Type).Ref("platform_products").Field("store_id").Unique(),
	}
}
