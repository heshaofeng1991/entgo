package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

type CourierOrder struct {
	ent.Schema
}

// Mixin of the InboundItem.
func (CourierOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (CourierOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned",
		}),
		field.String("order_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("courier_platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)",
		}).Default(""),
		field.String("shipping_method_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default(""),
		field.String("shipping_method_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default(""),
		field.String("tracking_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}).Default(""),
		field.String("tracking_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default(""),
		field.String("waybill_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default(""),
		field.String("courier_order_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default(""),
		field.String("shipping_label_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}).Default(""),
		field.Float("total_items_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)",
		}).Default(0),
		field.String("currency").SchemaType(map[string]string{
			dialect.MySQL: "char(3)",
		}).Optional(),
		field.Int("package_count").SchemaType(map[string]string{
			dialect.MySQL: "int",
		}).Default(1),
		field.Int("weight").SchemaType(map[string]string{
			dialect.MySQL: "float(9,0)",
		}).Default(0),
		field.String("receiver_address").SchemaType(map[string]string{
			dialect.MySQL: "json",
		}).Nillable().Optional(),
		field.String("sender_address").SchemaType(map[string]string{
			dialect.MySQL: "json",
		}).Nillable().Optional(),
		field.String("items").SchemaType(map[string]string{
			dialect.MySQL: "json",
		}).Nillable().Optional(),
		field.String("request_data").SchemaType(map[string]string{
			dialect.MySQL: "json",
		}).Nillable().Optional(),
		field.String("response_data").SchemaType(map[string]string{
			dialect.MySQL: "json",
		}).Nillable().Optional(),
		field.String("result_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default("").Optional(),
		field.String("message").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default("").Optional(),
		field.String("en_message").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}).Default("").Optional(),
		field.String("status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(2)",
		}).Default("N"),
	}
}

func (CourierOrder) Edges() []ent.Edge {
	return nil
}
