package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Users.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default(""),

		field.String("email").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		),

		field.String("type").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(24)", // Override MySQL.
			},
		).Default("customer"),

		field.String("password").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		),

		field.Int64("selected_warehouse_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Optional(),

		field.String("avatar").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default(""),

		field.String("code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Default(""),

		field.Int8("status").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint(1)", // Override MySQL.
			},
		).Default(1),

		field.Bool("guide_finished").Default(false),

		field.Int("guide_status").SchemaType(
			map[string]string{
				dialect.MySQL: "int", // Override MySQL.
			},
		).Default(0),

		field.String("hs_object_id").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(50)", // Override MySQL.
			},
		).Default(""),

		field.String("questions").SchemaType(
			map[string]string{
				dialect.MySQL: "text", // Override MySQL.
			},
		).Optional(),

		field.Time("last_logged_time").SchemaType(
			map[string]string{
				dialect.MySQL: "timestamp", // Override MySQL.
			},
		).Optional(),

		field.String("website").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(512)", // Override MySQL.
			},
		).Default(""),

		field.String("platform").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(512)", // Override MySQL.
			},
		).Default(""),

		field.String("concerns").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(512)", // Override MySQL.
			},
		).Default(""),

		field.String("store_code").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default(""),

		field.String("phone").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default(""),

		field.String("source").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default("").Comment("用户来源"),

		field.String("source_tag").SchemaType(
			map[string]string{
				dialect.MySQL: "varchar(255)", // Override MySQL.
			},
		).Default("").Comment("来源标签"),
	}
}

// Edges of the Users.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("announcements", Announcements.Type),
	}
}
