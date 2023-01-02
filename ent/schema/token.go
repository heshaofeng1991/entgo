/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    token
	@Date    2022/3/9 3:17 下午
	@Desc
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Mixin of the Tenant schema.
func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)", // Override MySQL.
		}),

		field.String("token").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("refresh_token").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("pin").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("trace_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("token_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(30)", // Override MySQL.
		}),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(1),

		field.String("expired_at").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("refresh_expired_at").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("response_at").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return nil
}
