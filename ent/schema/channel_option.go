/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_option
	@Date    2022/5/20 10:23
	@Desc
*/

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// ChannelOption holds the schema definition for the ChannelOption entity.
type ChannelOption struct {
	ent.Schema
}

// Fields of the ChannelOption.
func (ChannelOption) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Unique(),

		field.Int64("order_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}),

		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}),

		field.String("country_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Default(time.Now),

		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now),

		field.Time("deleted_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),
	}

}

// Edges of the ChannelOption.
func (ChannelOption) Edges() []ent.Edge {
	return nil
}
