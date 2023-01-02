package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Mixin of the OrderItem.
func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int unsigned", // Override MySQL.
			},
		).Optional(),

		field.Int64("product_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int unsigned", // Override MySQL.
			},
		).Default(0).Optional(),

		field.Int64("platform_product_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int unsigned", // Override MySQL.
			},
		).Default(0).Optional(),

		field.String("barcode").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(32)", // Override MySQL.
			},
		).Default(""),

		field.String("fulfillment_service").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(64)", // Override MySQL.
			},
		).Default(""),

		field.String("ext_order_item_id").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(64)", // Override MySQL.
			},
		).Default(""),

		field.String("ext_product_id").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(64)", // Override MySQL.
			},
		).Default(""),

		field.Bool("is_custom_item").Default(false),

		field.String("name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(512)", // Override MySQL.
			},
		).Default(""),

		field.String("declared_cn_name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(256)", // Override MySQL.
			},
		).Optional().Default(""),

		field.String("hs_code").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(16)", // Override MySQL.
			},
		).Default(""),

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

		field.String("images").SchemaType(
			map[string]string{
				dialect.MySQL: "json", // Override MySQL.
			},
		).Optional(),

		field.String("attributes").SchemaType(
			map[string]string{
				dialect.MySQL: "json", // Override MySQL.
			},
		).Optional(),

		field.Int("grams").SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(9,0)", // Override MySQL.
			},
		).Default(0),

		field.Int("length").SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(9,0)", // Override MySQL.
			},
		).Default(0),

		field.Int("width").SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(9,0)", // Override MySQL.
			},
		).Default(0),

		field.Int("height").SchemaType(map[string]string{
			dialect.MySQL: "decimal(9,0)", // Override MySQL.
		}).Default(0),

		field.Int("qty").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}),

		field.Float("unit_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000),

		field.Float("declared_value_in_usd").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000),

		field.Float("declared_value_in_eur").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000),

		field.String("currency").SchemaType(map[string]string{
			dialect.MySQL: "char(3)", // Override MySQL.
		}).Default("USD"),

		field.Int("fulfill_qty").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0),

		field.Bool("requires_shipping").Default(true),
		field.Bool("gift_card").Default(false),
		field.Bool("taxable").Default(false),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Default(1),

		field.String("sku").SchemaType(map[string]string{
			dialect.MySQL: "varchar(256)", // Override MySQL.
		}).Optional(),

		field.String("listing_sku").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(256)", // Override MySQL.
			},
		).Optional(),

		field.String("declared_en_name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(256)", // Override MySQL.
			},
		).Optional(),

		field.String("product_name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(256)", // Override MySQL.
			},
		).Optional(),

		field.String("customer_code").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(64)", // Override MySQL.
			},
		),
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order_items").Field("order_id").Unique(),
	}
}
