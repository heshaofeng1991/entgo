package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InboundItem holds the schema definition for the InboundItem entity.
type InboundItem struct {
	ent.Schema
}

// Mixin of the InboundItem.
func (InboundItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the InboundItem.
func (InboundItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("inbound_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Comment("入库单ID").Optional(),

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

		field.Int("qty").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Comment("产品数量"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("状态"),

		field.String("customer_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),
	}
}

// Edges of the InboundItem.
func (InboundItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("inbounds", Inbound.Type).Ref("inbound_items").Field("inbound_id").Unique(),
	}
}
