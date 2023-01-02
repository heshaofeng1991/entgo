// Package schema @Author johnny 2022/1/13 3:32 PM:00
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Inbound holds the schema definition for the Inbounds entity.
type Inbound struct {
	ent.Schema
}

// Mixin of the Inbound.
func (Inbound) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Inbounds.
func (Inbound) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_order_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Comment("来自于客户的外部订单ID"),

		field.String("customer_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Comment("客户代码"),

		field.String("tracking_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Comment("订单追踪条码"),

		field.Int64("warehouse_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Comment("入库单对应的仓库ID"),

		field.Text("description").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).Optional().Comment("入库单信息描述"),

		field.Time("estimated_arrival_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional().Comment("预计到达仓库时间").StorageKey("estimated_arrival"),

		field.Time("shipped_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional().Comment("发货时间").StorageKey("ship_date"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("入库单状态 0=>Draft; 1=>New; 2=>InTransit; 10=>StartReceived; 20=>CompleteReceived; 30=>StartShelved; 40=>CompleteShelved;  99=>Complete"),

		field.Int8("type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("发货模式 0 直发转运， 1 仓储模式"),

		field.Int8("is_pickup").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0).Comment("是否上门揽件 0 => false， 1 => true"),

		field.String("shipping_mark_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Optional().Comment("PDF URL"),

		field.Int64("pickup_order_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Comment("揽件信息ID"),

		field.String("carrier_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Comment("物流供应商名称"),

		field.String("order_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").StorageKey("inbound_order_number"),
	}
}

// Edges of the Inbounds.
func (Inbound) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("inbound_items", InboundItem.Type),
	}
}
