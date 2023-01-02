package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Default(""),

		field.String("sku").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}),

		field.String("barcode").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.String("customer_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.String("declared_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}).Optional(),

		field.String("declared_cn_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}).Optional(),

		field.Float("declared_value_in_usd").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000),

		field.Float("declared_value_in_eur").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000),

		field.String("currency").SchemaType(map[string]string{
			dialect.MySQL: "char(3)", // Override MySQL.
		}).Default(""),

		field.String("hs_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(16)", // Override MySQL.
		}).Default(""),

		field.String("material").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(64)", // Override MySQL.
			},
		).Default(""),

		field.String("purpose").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(64)", // Override MySQL.
			},
		).Default(""),

		field.Int8("with_barcode").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0).Comment("产品自带barcode"),

		field.Int8("barcode_service").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0).Comment("是否代贴barcode服务"),

		field.String("barcode_template").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional().Comment("打印barcode选项"),

		field.String("images").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional(),

		field.String("attributes").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional().Comment("产品属性"),

		field.String("confirmed_attributes").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional().Comment("已审核的产品属性"),

		field.Int("grams").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,0)", // Override MySQL.
		}).Default(0).Comment("来源客户"),

		field.Int("inbound_grams").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,0)", // Override MySQL.
		}).Default(0).Comment("入库实重"),

		field.Int("length").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,0)", // Override MySQL.
		}).Default(0),

		field.Int("width").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,0)", // Override MySQL.
		}).Default(0),

		field.Int("height").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,0)", // Override MySQL.
		}).Default(0),

		field.Int("max_agv_qty").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(100).Comment("AGV最大上架数量"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Default(1).Comment("1=>New, 10=>Normal"),

		field.Int64("created_by").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("inventories", Inventory.Type),
		edge.To("product_mappings", ProductMapping.Type),
	}
}
