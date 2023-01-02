/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    channel_volume_factor
	@Date    2022/5/30 17:57
	@Desc
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// ChannelVolumeFactor holds the schema definition for the ChannelVolumeFactor entity.
type ChannelVolumeFactor struct {
	ent.Schema
}

// Mixin of the ChannelVolumeFactor.
func (ChannelVolumeFactor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ChannelVolumeFactor.
func (ChannelVolumeFactor) Fields() []ent.Field {
	return []ent.Field{
		field.String("country_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10)unsigned", // Override MySQL.
		}),

		field.Int("volume_factor").SchemaType(map[string]string{
			dialect.MySQL: "int", // Override MySQL.
		}).Default(0).Comment("体积因子"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("0=>Inactive, 1=>Active"),
	}
}

// Edges of the ChannelVolumeFactor.
func (ChannelVolumeFactor) Edges() []ent.Edge {
	return nil
}
