// Code generated by ent, DO NOT EDIT.

package userreferral

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/heshaofeng1991/entgo/ent/gen/internal"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// FirstShipmentDate applies equality check predicate on the "first_shipment_date" field. It's identical to FirstShipmentDateEQ.
func FirstShipmentDate(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstShipmentDate), v))
	})
}

// TotalCommission applies equality check predicate on the "total_commission" field. It's identical to TotalCommissionEQ.
func TotalCommission(v float64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalCommission), v))
	})
}

// InvitedByUserID applies equality check predicate on the "invited_by_user_id" field. It's identical to InvitedByUserIDEQ.
func InvitedByUserID(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvitedByUserID), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// FirstShipmentDateEQ applies the EQ predicate on the "first_shipment_date" field.
func FirstShipmentDateEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstShipmentDate), v))
	})
}

// FirstShipmentDateNEQ applies the NEQ predicate on the "first_shipment_date" field.
func FirstShipmentDateNEQ(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstShipmentDate), v))
	})
}

// FirstShipmentDateIn applies the In predicate on the "first_shipment_date" field.
func FirstShipmentDateIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFirstShipmentDate), v...))
	})
}

// FirstShipmentDateNotIn applies the NotIn predicate on the "first_shipment_date" field.
func FirstShipmentDateNotIn(vs ...time.Time) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFirstShipmentDate), v...))
	})
}

// FirstShipmentDateGT applies the GT predicate on the "first_shipment_date" field.
func FirstShipmentDateGT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstShipmentDate), v))
	})
}

// FirstShipmentDateGTE applies the GTE predicate on the "first_shipment_date" field.
func FirstShipmentDateGTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstShipmentDate), v))
	})
}

// FirstShipmentDateLT applies the LT predicate on the "first_shipment_date" field.
func FirstShipmentDateLT(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstShipmentDate), v))
	})
}

// FirstShipmentDateLTE applies the LTE predicate on the "first_shipment_date" field.
func FirstShipmentDateLTE(v time.Time) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstShipmentDate), v))
	})
}

// FirstShipmentDateIsNil applies the IsNil predicate on the "first_shipment_date" field.
func FirstShipmentDateIsNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFirstShipmentDate)))
	})
}

// FirstShipmentDateNotNil applies the NotNil predicate on the "first_shipment_date" field.
func FirstShipmentDateNotNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFirstShipmentDate)))
	})
}

// TotalCommissionEQ applies the EQ predicate on the "total_commission" field.
func TotalCommissionEQ(v float64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionNEQ applies the NEQ predicate on the "total_commission" field.
func TotalCommissionNEQ(v float64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionIn applies the In predicate on the "total_commission" field.
func TotalCommissionIn(vs ...float64) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalCommission), v...))
	})
}

// TotalCommissionNotIn applies the NotIn predicate on the "total_commission" field.
func TotalCommissionNotIn(vs ...float64) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalCommission), v...))
	})
}

// TotalCommissionGT applies the GT predicate on the "total_commission" field.
func TotalCommissionGT(v float64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionGTE applies the GTE predicate on the "total_commission" field.
func TotalCommissionGTE(v float64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionLT applies the LT predicate on the "total_commission" field.
func TotalCommissionLT(v float64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalCommission), v))
	})
}

// TotalCommissionLTE applies the LTE predicate on the "total_commission" field.
func TotalCommissionLTE(v float64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalCommission), v))
	})
}

// InvitedByUserIDEQ applies the EQ predicate on the "invited_by_user_id" field.
func InvitedByUserIDEQ(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvitedByUserID), v))
	})
}

// InvitedByUserIDNEQ applies the NEQ predicate on the "invited_by_user_id" field.
func InvitedByUserIDNEQ(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInvitedByUserID), v))
	})
}

// InvitedByUserIDIn applies the In predicate on the "invited_by_user_id" field.
func InvitedByUserIDIn(vs ...int64) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInvitedByUserID), v...))
	})
}

// InvitedByUserIDNotIn applies the NotIn predicate on the "invited_by_user_id" field.
func InvitedByUserIDNotIn(vs ...int64) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInvitedByUserID), v...))
	})
}

// InvitedByUserIDGT applies the GT predicate on the "invited_by_user_id" field.
func InvitedByUserIDGT(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInvitedByUserID), v))
	})
}

// InvitedByUserIDGTE applies the GTE predicate on the "invited_by_user_id" field.
func InvitedByUserIDGTE(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInvitedByUserID), v))
	})
}

// InvitedByUserIDLT applies the LT predicate on the "invited_by_user_id" field.
func InvitedByUserIDLT(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInvitedByUserID), v))
	})
}

// InvitedByUserIDLTE applies the LTE predicate on the "invited_by_user_id" field.
func InvitedByUserIDLTE(v int64) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInvitedByUserID), v))
	})
}

// InvitedByUserIDIsNil applies the IsNil predicate on the "invited_by_user_id" field.
func InvitedByUserIDIsNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInvitedByUserID)))
	})
}

// InvitedByUserIDNotNil applies the NotNil predicate on the "invited_by_user_id" field.
func InvitedByUserIDNotNil() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInvitedByUserID)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.UserReferral {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// HasTenant applies the HasEdge predicate on the "tenant" edge.
func HasTenant() predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.UserReferral
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTenantWith applies the HasEdge predicate on the "tenant" edge with a given conditions (other predicates).
func HasTenantWith(preds ...predicate.Tenant) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Tenant
		step.Edge.Schema = schemaConfig.UserReferral
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserReferral) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserReferral) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
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
func Not(p predicate.UserReferral) predicate.UserReferral {
	return predicate.UserReferral(func(s *sql.Selector) {
		p(s.Not())
	})
}
