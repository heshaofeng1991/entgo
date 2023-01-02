package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Mixin of the Transaction schema.
func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}),
		field.String("transaction_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),
		field.Float("transaction_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("balance").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1),
		field.Int8("created_by").SchemaType(map[string]string{
			dialect.MySQL: "bigint(4)", // Override MySQL.
		}).Default(0),
		field.Int8("updated_by").SchemaType(map[string]string{
			dialect.MySQL: "bigint(4)", // Override MySQL.
		}).Default(0),
	}
}

// Policy defines the privacy policy of the User.
func (Transaction) Policy() ent.Policy {
	return nil
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("details", TransactionDetail.Type),
	}
}
