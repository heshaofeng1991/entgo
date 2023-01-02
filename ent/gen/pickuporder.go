// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/pickuporder"
)

// PickupOrder is the model entity for the PickupOrder schema.
type PickupOrder struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 请求上门揽件时间
	RequestedPickupAt time.Time `json:"requested_pickup_at,omitempty"`
	// 揽件状态 0 未取件 1 已取件
	Status int8 `json:"status,omitempty"`
	// 地址信息
	SenderAddressInfo string `json:"sender_address_info,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PickupOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pickuporder.FieldID, pickuporder.FieldStatus:
			values[i] = new(sql.NullInt64)
		case pickuporder.FieldSenderAddressInfo:
			values[i] = new(sql.NullString)
		case pickuporder.FieldCreatedAt, pickuporder.FieldUpdatedAt, pickuporder.FieldDeletedAt, pickuporder.FieldRequestedPickupAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PickupOrder", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PickupOrder fields.
func (po *PickupOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pickuporder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int64(value.Int64)
		case pickuporder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case pickuporder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case pickuporder.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				po.DeletedAt = value.Time
			}
		case pickuporder.FieldRequestedPickupAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field requested_pickup_at", values[i])
			} else if value.Valid {
				po.RequestedPickupAt = value.Time
			}
		case pickuporder.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				po.Status = int8(value.Int64)
			}
		case pickuporder.FieldSenderAddressInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender_address_info", values[i])
			} else if value.Valid {
				po.SenderAddressInfo = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PickupOrder.
// Note that you need to call PickupOrder.Unwrap() before calling this method if this PickupOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *PickupOrder) Update() *PickupOrderUpdateOne {
	return (&PickupOrderClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the PickupOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *PickupOrder) Unwrap() *PickupOrder {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("gen: PickupOrder is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *PickupOrder) String() string {
	var builder strings.Builder
	builder.WriteString("PickupOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(po.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("requested_pickup_at=")
	builder.WriteString(po.RequestedPickupAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", po.Status))
	builder.WriteString(", ")
	builder.WriteString("sender_address_info=")
	builder.WriteString(po.SenderAddressInfo)
	builder.WriteByte(')')
	return builder.String()
}

// PickupOrders is a parsable slice of PickupOrder.
type PickupOrders []*PickupOrder

func (po PickupOrders) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
