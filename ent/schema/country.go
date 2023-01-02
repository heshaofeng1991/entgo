package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

type Country struct {
	ent.Schema
}

func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "char(2)",
		}),
		field.String("cn_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)",
		}),
		field.String("en_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)",
		}),
	}
}

func (Country) Edges() []ent.Edge {
	return nil
}
