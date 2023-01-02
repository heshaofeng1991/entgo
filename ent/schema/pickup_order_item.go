package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// PickupOrderItem holds the schema definition for the PickupOrderItem entity.
type PickupOrderItem struct {
	ent.Schema
}

// Mixin of the PickupOrderItem.
func (PickupOrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the PickupOrderItem.
func (PickupOrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("pickup_order_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Comment("揽件订单ID"),

		field.Int64("product_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Comment("产品ID"),

		field.String("product_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Comment("产品名称"),

		field.String("sku").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Comment("产品SKU"),

		field.String("barcode").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Comment("产品条码"),

		field.Int32("qty").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Comment("产品数量"),

		field.String("customer_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),
	}
}

// Edges of the PickupOrderItem.
func (PickupOrderItem) Edges() []ent.Edge {
	return nil
}
