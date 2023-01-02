/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    sequence
	@Date    2022/2/7 3:02 下午:00
	@Desc
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Sequence holds the schema definition for the Sequence entity.
type Sequence struct {
	ent.Schema
}

// Mixin of the Sequence.
func (Sequence) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Sequence.
func (Sequence) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.String("prefix").SchemaType(map[string]string{
			dialect.MySQL: "varchar(8)", // Override MySQL.
		}).Default(""),

		field.Int64("value").SchemaType(map[string]string{
			dialect.MySQL: "bigint(10) unsigned zerofill", // Override MySQL.
		}).Default(0o000000000),

		field.String("display_value").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Default(""),
	}
}

// Edges of the Sequence.
func (Sequence) Edges() []ent.Edge {
	return nil
}
