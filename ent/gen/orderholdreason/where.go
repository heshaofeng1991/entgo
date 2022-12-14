// Code generated by ent, DO NOT EDIT.

package orderholdreason

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// ProductID applies equality check predicate on the "product_id" field. It's identical to ProductIDEQ.
func ProductID(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductID), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v int32) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// EnReason applies equality check predicate on the "en_reason" field. It's identical to EnReasonEQ.
func EnReason(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnReason), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// ProductIDEQ applies the EQ predicate on the "product_id" field.
func ProductIDEQ(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductID), v))
	})
}

// ProductIDNEQ applies the NEQ predicate on the "product_id" field.
func ProductIDNEQ(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductID), v))
	})
}

// ProductIDIn applies the In predicate on the "product_id" field.
func ProductIDIn(vs ...int64) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProductID), v...))
	})
}

// ProductIDNotIn applies the NotIn predicate on the "product_id" field.
func ProductIDNotIn(vs ...int64) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProductID), v...))
	})
}

// ProductIDGT applies the GT predicate on the "product_id" field.
func ProductIDGT(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductID), v))
	})
}

// ProductIDGTE applies the GTE predicate on the "product_id" field.
func ProductIDGTE(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductID), v))
	})
}

// ProductIDLT applies the LT predicate on the "product_id" field.
func ProductIDLT(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductID), v))
	})
}

// ProductIDLTE applies the LTE predicate on the "product_id" field.
func ProductIDLTE(v int64) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductID), v))
	})
}

// ProductIDIsNil applies the IsNil predicate on the "product_id" field.
func ProductIDIsNil() predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProductID)))
	})
}

// ProductIDNotNil applies the NotNil predicate on the "product_id" field.
func ProductIDNotNil() predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProductID)))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v int32) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v int32) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...int32) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...int32) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v int32) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v int32) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v int32) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v int32) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReason), v))
	})
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReason), v))
	})
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReason), v...))
	})
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReason), v...))
	})
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReason), v))
	})
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReason), v))
	})
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReason), v))
	})
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReason), v))
	})
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReason), v))
	})
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReason), v))
	})
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReason), v))
	})
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReason), v))
	})
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReason), v))
	})
}

// EnReasonEQ applies the EQ predicate on the "en_reason" field.
func EnReasonEQ(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnReason), v))
	})
}

// EnReasonNEQ applies the NEQ predicate on the "en_reason" field.
func EnReasonNEQ(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnReason), v))
	})
}

// EnReasonIn applies the In predicate on the "en_reason" field.
func EnReasonIn(vs ...string) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEnReason), v...))
	})
}

// EnReasonNotIn applies the NotIn predicate on the "en_reason" field.
func EnReasonNotIn(vs ...string) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEnReason), v...))
	})
}

// EnReasonGT applies the GT predicate on the "en_reason" field.
func EnReasonGT(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnReason), v))
	})
}

// EnReasonGTE applies the GTE predicate on the "en_reason" field.
func EnReasonGTE(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnReason), v))
	})
}

// EnReasonLT applies the LT predicate on the "en_reason" field.
func EnReasonLT(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnReason), v))
	})
}

// EnReasonLTE applies the LTE predicate on the "en_reason" field.
func EnReasonLTE(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnReason), v))
	})
}

// EnReasonContains applies the Contains predicate on the "en_reason" field.
func EnReasonContains(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEnReason), v))
	})
}

// EnReasonHasPrefix applies the HasPrefix predicate on the "en_reason" field.
func EnReasonHasPrefix(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEnReason), v))
	})
}

// EnReasonHasSuffix applies the HasSuffix predicate on the "en_reason" field.
func EnReasonHasSuffix(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEnReason), v))
	})
}

// EnReasonEqualFold applies the EqualFold predicate on the "en_reason" field.
func EnReasonEqualFold(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEnReason), v))
	})
}

// EnReasonContainsFold applies the ContainsFold predicate on the "en_reason" field.
func EnReasonContainsFold(v string) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEnReason), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrderHoldReason {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrdersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrdersTable, OrdersColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Order
		step.Edge.Schema = schemaConfig.OrderHoldReason
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrdersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrdersTable, OrdersColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Order
		step.Edge.Schema = schemaConfig.OrderHoldReason
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderHoldReason) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderHoldReason) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
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
func Not(p predicate.OrderHoldReason) predicate.OrderHoldReason {
	return predicate.OrderHoldReason(func(s *sql.Selector) {
		p(s.Not())
	})
}
