// Code generated by ent, DO NOT EDIT.

package customerconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// ExcludeCountryCode applies equality check predicate on the "exclude_country_code" field. It's identical to ExcludeCountryCodeEQ.
func ExcludeCountryCode(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExcludeCountryCode), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...int64) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChannelID), v...))
	})
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...int64) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChannelID), v...))
	})
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChannelID), v))
	})
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChannelID), v))
	})
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChannelID), v))
	})
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v int64) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChannelID), v))
	})
}

// ExcludeCountryCodeEQ applies the EQ predicate on the "exclude_country_code" field.
func ExcludeCountryCodeEQ(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeNEQ applies the NEQ predicate on the "exclude_country_code" field.
func ExcludeCountryCodeNEQ(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeIn applies the In predicate on the "exclude_country_code" field.
func ExcludeCountryCodeIn(vs ...string) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExcludeCountryCode), v...))
	})
}

// ExcludeCountryCodeNotIn applies the NotIn predicate on the "exclude_country_code" field.
func ExcludeCountryCodeNotIn(vs ...string) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExcludeCountryCode), v...))
	})
}

// ExcludeCountryCodeGT applies the GT predicate on the "exclude_country_code" field.
func ExcludeCountryCodeGT(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeGTE applies the GTE predicate on the "exclude_country_code" field.
func ExcludeCountryCodeGTE(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeLT applies the LT predicate on the "exclude_country_code" field.
func ExcludeCountryCodeLT(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeLTE applies the LTE predicate on the "exclude_country_code" field.
func ExcludeCountryCodeLTE(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeContains applies the Contains predicate on the "exclude_country_code" field.
func ExcludeCountryCodeContains(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeHasPrefix applies the HasPrefix predicate on the "exclude_country_code" field.
func ExcludeCountryCodeHasPrefix(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeHasSuffix applies the HasSuffix predicate on the "exclude_country_code" field.
func ExcludeCountryCodeHasSuffix(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeEqualFold applies the EqualFold predicate on the "exclude_country_code" field.
func ExcludeCountryCodeEqualFold(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExcludeCountryCode), v))
	})
}

// ExcludeCountryCodeContainsFold applies the ContainsFold predicate on the "exclude_country_code" field.
func ExcludeCountryCodeContainsFold(v string) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExcludeCountryCode), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.CustomerConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// HasTenant applies the HasEdge predicate on the "tenant" edge.
func HasTenant() predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.CustomerConfig
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTenantWith applies the HasEdge predicate on the "tenant" edge with a given conditions (other predicates).
func HasTenantWith(preds ...predicate.Tenant) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.CustomerConfig
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CustomerConfig) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CustomerConfig) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CustomerConfig) predicate.CustomerConfig {
	return predicate.CustomerConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
