package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// TaskSchedule holds the schema definition for the TaskSchedule entity.
type TaskSchedule struct {
	ent.Schema
}

// Mixin of the TaskSchedule.
func (TaskSchedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the TaskSchedule.
func (TaskSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.String("platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),

		field.Int64("store_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),

		field.String("func_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),

		field.Bool("locked").Default(false),

		field.Int64("locked_times").SchemaType(map[string]string{
			dialect.MySQL: "int(4)", // Override MySQL.
		}).Default(0),

		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).Optional().Optional(),

		field.Time("last_access_at").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(nil).Optional(),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(1),
	}
}

// Edges of the TaskSchedule.
func (TaskSchedule) Edges() []ent.Edge {
	return nil
}
