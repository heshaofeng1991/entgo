/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_recommend
	@Date    2022/5/24 10:21
	@Desc
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// ChannelRecommend holds the schema definition for the ChannelRecommend entity.
type ChannelRecommend struct {
	ent.Schema
}

// Mixin of the ChannelRecommend.
func (ChannelRecommend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ChannelRecommend.
func (ChannelRecommend) Fields() []ent.Field {
	return []ent.Field{
		field.String("country_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),
		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10)unsigned", // Override MySQL.
		}),
		field.Int8("is_recommended").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0).Comment("0=>不推荐, 1=>推荐"),
		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("0=>Inactive, 1=>Active"),
		field.String("value").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),
	}
}

// Edges of the ChannelRecommend.
func (ChannelRecommend) Edges() []ent.Edge {
	return nil
}

