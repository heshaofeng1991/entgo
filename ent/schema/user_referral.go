/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    user_referral
	@Date    2022/3/19 6:12 下午
	@Desc
*/

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// UserReferral holds the schema definition for the UserReferral entity.
type UserReferral struct {
	ent.Schema
}

// Mixin of the UserReferral.
func (UserReferral) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the UserReferral.
func (UserReferral) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int unsigned", // Override MySQL.
			},
		).Optional().Comment("被邀请人ID"),

		field.Time("first_shipment_date").SchemaType(
			map[string]string{
				dialect.MySQL: "date", // Override MySQL.
			},
		).Optional().Comment("First Shipment Date"),

		field.Float("total_commission").SchemaType(
			map[string]string{
				dialect.MySQL: "decimal(15,2)", // Override MySQL.
			},
		).Default(0.00).Comment("Total Commission"),

		field.Int64("invited_by_user_id").SchemaType(
			map[string]string{
				dialect.MySQL: "int unsigned", // Override MySQL.
			},
		).Optional().Comment("Invited by user ID"),

		field.Int8("status").SchemaType(
			map[string]string{
				dialect.MySQL: "tinyint", // Override MySQL.
			},
		).Default(1).Comment("1=>Pending, 2=>Qualified"),
	}
}

// Edges of the UserReferral.
func (UserReferral) Edges() []ent.Edge {
	return nil
}
