package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// TransactionDetail holds the schema definition for the TransactionDetail entity.
type TransactionDetail struct {
	ent.Schema
}

// Mixin of the TransactionDetail schema.
func (TransactionDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the TransactionDetail.
func (TransactionDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}),
		field.Int64("transaction_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}),
		field.String("transaction_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),

		field.Float("delivery_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("misc_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("fuel_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("registration_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("processing_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("package_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("handling_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("vat").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,2)", // Override MySQL.
		}).Default(0),
		field.Int("weight").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Default(0),
	}
}

// Policy defines the privacy policy of the User.
func (TransactionDetail) Policy() ent.Policy {
	return nil
}
