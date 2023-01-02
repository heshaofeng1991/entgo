package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Announcements struct {
	ent.Schema
}

func (Announcements) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Announcements) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),

		field.String("content").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).Default(""),

		field.Int("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint", // Override MySQL.
		}).Default(1),

		field.Int("index").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Default(0),

		field.Int64("create_by").SchemaType(map[string]string{
			dialect.MySQL: "int(10)", // Override MySQL.
		}).Default(0).Positive().Optional(),

		field.Time("expiration"),

		field.Time("effective_time"),
	}
}

// Edges of the Announcements.
func (Announcements) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("announcements").Field("create_by").Unique(),
	}
}
