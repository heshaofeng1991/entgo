/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    inventory
	@Date    2022/2/24 7:08 下午:00
	@Desc
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Inventory holds the schema definition for the Inventory entity.
type Inventory struct {
	ent.Schema
}

// Mixin of the Inventory.
func (Inventory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Inventory.
func (Inventory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("product_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int unsigned", // Override MySQL.
			},
		).Optional(),

		field.Int64("warehouse_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int unsigned", // Override MySQL.
			},
		),

		field.Int32("storage_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("总数量=可用数量+准备发货数量+准备上架数量"),

		field.Int32("available_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("可用数量"),

		field.Int32("prepare_ship_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("准备发货数量"),

		field.Int32("prepare_shelve_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("准备上架数量"),

		field.Int32("quicktron_storage_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("快仓总数量=可用数量+准备下架数量+准备上架数量"),

		field.Int32("quicktron_available_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("快仓可用数量"),

		field.Int32("quicktron_prepare_outbound_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("快仓准备下架数量"),

		field.Int32("quicktron_prepare_shelve_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("快仓准备上架数量"),

		field.Int32("normal_storage_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("传统区域总数量=可用数量+准备下架数量+准备上架数量"),

		field.Int32("normal_available_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("传统区域可用数量"),

		field.Int32("normal_prepare_outbound_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("传统区域准备下架数量"),

		field.Int32("normal_prepare_shelve_qty").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0).Comment("传统区域准备上架数量"),

		field.Int8("status").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint", // Override MySQL.
			},
		).Default(1),
	}
}

// Edges of the Inventory.
func (Inventory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).Ref("inventories").Field("product_id").Unique(),
	}
}
