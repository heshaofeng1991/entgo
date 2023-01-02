package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// PickupOrder holds the schema definition for the PickupOrder entity.
type PickupOrder struct {
	ent.Schema
}

// Mixin of the PickupOrder.
func (PickupOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the PickupOrder.
func (PickupOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Time("requested_pickup_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional().Comment("请求上门揽件时间"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Default(0).Comment("揽件状态 0 未取件 1 已取件"),

		field.String("sender_address_info").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional().Comment("地址信息"),
	}
}

// Edges of the PickupOrder.
func (PickupOrder) Edges() []ent.Edge {
	return nil
}
