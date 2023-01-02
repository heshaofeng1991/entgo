// Code generated by ent, DO NOT EDIT.

package channeloption

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// CountryCode applies equality check predicate on the "country_code" field. It's identical to CountryCodeEQ.
func CountryCode(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountryCode), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChannelID), v))
	})
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...int64) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChannelID), v...))
	})
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...int64) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChannelID), v...))
	})
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChannelID), v))
	})
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChannelID), v))
	})
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChannelID), v))
	})
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v int64) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChannelID), v))
	})
}

// CountryCodeEQ applies the EQ predicate on the "country_code" field.
func CountryCodeEQ(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountryCode), v))
	})
}

// CountryCodeNEQ applies the NEQ predicate on the "country_code" field.
func CountryCodeNEQ(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCountryCode), v))
	})
}

// CountryCodeIn applies the In predicate on the "country_code" field.
func CountryCodeIn(vs ...string) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCountryCode), v...))
	})
}

// CountryCodeNotIn applies the NotIn predicate on the "country_code" field.
func CountryCodeNotIn(vs ...string) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCountryCode), v...))
	})
}

// CountryCodeGT applies the GT predicate on the "country_code" field.
func CountryCodeGT(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCountryCode), v))
	})
}

// CountryCodeGTE applies the GTE predicate on the "country_code" field.
func CountryCodeGTE(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCountryCode), v))
	})
}

// CountryCodeLT applies the LT predicate on the "country_code" field.
func CountryCodeLT(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCountryCode), v))
	})
}

// CountryCodeLTE applies the LTE predicate on the "country_code" field.
func CountryCodeLTE(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCountryCode), v))
	})
}

// CountryCodeContains applies the Contains predicate on the "country_code" field.
func CountryCodeContains(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCountryCode), v))
	})
}

// CountryCodeHasPrefix applies the HasPrefix predicate on the "country_code" field.
func CountryCodeHasPrefix(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCountryCode), v))
	})
}

// CountryCodeHasSuffix applies the HasSuffix predicate on the "country_code" field.
func CountryCodeHasSuffix(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCountryCode), v))
	})
}

// CountryCodeEqualFold applies the EqualFold predicate on the "country_code" field.
func CountryCodeEqualFold(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCountryCode), v))
	})
}

// CountryCodeContainsFold applies the ContainsFold predicate on the "country_code" field.
func CountryCodeContainsFold(v string) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCountryCode), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ChannelOption {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChannelOption) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChannelOption) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
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
func Not(p predicate.ChannelOption) predicate.ChannelOption {
	return predicate.ChannelOption(func(s *sql.Selector) {
		p(s.Not())
	})
}
