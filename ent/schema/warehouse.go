package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Warehouse struct {
	ent.Schema
}

// Mixin of the Channel.
func (Warehouse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Warehouse) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)",
		}),
		field.String("quicktron_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)",
		}),
		field.Int("enable_quicktron").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)",
		}),
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("company").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("first_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("last_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("address1").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("address2").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("country_code").SchemaType(map[string]string{
			dialect.MySQL: "char(2)",
		}),
		field.String("country_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("province").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.String("city").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)",
		}),
		field.String("zip_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)",
		}),
		field.String("phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)",
		}),
		field.Int("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)",
		}),
	}
}

// Edges of the Users.
func (Warehouse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
	}
}
