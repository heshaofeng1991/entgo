// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/order"
	"github.com/heshaofeng1991/entgo/ent/gen/orderholdreason"
)

// OrderHoldReason is the model entity for the OrderHoldReason schema.
type OrderHoldReason struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int64 `json:"order_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int64 `json:"product_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Code holds the value of the "code" field.
	Code int32 `json:"code,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// EnReason holds the value of the "en_reason" field.
	EnReason string `json:"en_reason,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderHoldReasonQuery when eager-loading is set.
	Edges OrderHoldReasonEdges `json:"edges"`
}

// OrderHoldReasonEdges holds the relations/edges for other nodes in the graph.
type OrderHoldReasonEdges struct {
	// Orders holds the value of the orders edge.
	Orders *Order `json:"orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderHoldReasonEdges) OrdersOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Orders == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderHoldReason) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderholdreason.FieldID, orderholdreason.FieldOrderID, orderholdreason.FieldProductID, orderholdreason.FieldCode:
			values[i] = new(sql.NullInt64)
		case orderholdreason.FieldType, orderholdreason.FieldReason, orderholdreason.FieldEnReason:
			values[i] = new(sql.NullString)
		case orderholdreason.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderHoldReason", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderHoldReason fields.
func (ohr *OrderHoldReason) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderholdreason.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ohr.ID = int64(value.Int64)
		case orderholdreason.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				ohr.OrderID = value.Int64
			}
		case orderholdreason.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				ohr.ProductID = value.Int64
			}
		case orderholdreason.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ohr.Type = value.String
			}
		case orderholdreason.FieldCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ohr.Code = int32(value.Int64)
			}
		case orderholdreason.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				ohr.Reason = value.String
			}
		case orderholdreason.FieldEnReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field en_reason", values[i])
			} else if value.Valid {
				ohr.EnReason = value.String
			}
		case orderholdreason.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ohr.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryOrders queries the "orders" edge of the OrderHoldReason entity.
func (ohr *OrderHoldReason) QueryOrders() *OrderQuery {
	return (&OrderHoldReasonClient{config: ohr.config}).QueryOrders(ohr)
}

// Update returns a builder for updating this OrderHoldReason.
// Note that you need to call OrderHoldReason.Unwrap() before calling this method if this OrderHoldReason
// was returned from a transaction, and the transaction was committed or rolled back.
func (ohr *OrderHoldReason) Update() *OrderHoldReasonUpdateOne {
	return (&OrderHoldReasonClient{config: ohr.config}).UpdateOne(ohr)
}

// Unwrap unwraps the OrderHoldReason entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ohr *OrderHoldReason) Unwrap() *OrderHoldReason {
	_tx, ok := ohr.config.driver.(*txDriver)
	if !ok {
		panic("gen: OrderHoldReason is not a transactional entity")
	}
	ohr.config.driver = _tx.drv
	return ohr
}

// String implements the fmt.Stringer.
func (ohr *OrderHoldReason) String() string {
	var builder strings.Builder
	builder.WriteString("OrderHoldReason(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ohr.ID))
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", ohr.OrderID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", ohr.ProductID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(ohr.Type)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(fmt.Sprintf("%v", ohr.Code))
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(ohr.Reason)
	builder.WriteString(", ")
	builder.WriteString("en_reason=")
	builder.WriteString(ohr.EnReason)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ohr.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OrderHoldReasons is a parsable slice of OrderHoldReason.
type OrderHoldReasons []*OrderHoldReason

func (ohr OrderHoldReasons) config(cfg config) {
	for _i := range ohr {
		ohr[_i].config = cfg
	}
}
