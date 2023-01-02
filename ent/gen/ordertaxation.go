// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/order"
	"github.com/heshaofeng1991/entgo/ent/gen/ordertaxation"
)

// OrderTaxation is the model entity for the OrderTaxation schema.
type OrderTaxation struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int64 `json:"order_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode string `json:"country_code,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderTaxationQuery when eager-loading is set.
	Edges OrderTaxationEdges `json:"edges"`
}

// OrderTaxationEdges holds the relations/edges for other nodes in the graph.
type OrderTaxationEdges struct {
	// Orders holds the value of the orders edge.
	Orders *Order `json:"orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderTaxationEdges) OrdersOrErr() (*Order, error) {
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
func (*OrderTaxation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ordertaxation.FieldID, ordertaxation.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case ordertaxation.FieldType, ordertaxation.FieldCountryCode, ordertaxation.FieldNumber:
			values[i] = new(sql.NullString)
		case ordertaxation.FieldCreatedAt, ordertaxation.FieldUpdatedAt, ordertaxation.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderTaxation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderTaxation fields.
func (ot *OrderTaxation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ordertaxation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ot.ID = int64(value.Int64)
		case ordertaxation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ot.CreatedAt = value.Time
			}
		case ordertaxation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ot.UpdatedAt = value.Time
			}
		case ordertaxation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ot.DeletedAt = value.Time
			}
		case ordertaxation.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				ot.OrderID = value.Int64
			}
		case ordertaxation.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ot.Type = value.String
			}
		case ordertaxation.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				ot.CountryCode = value.String
			}
		case ordertaxation.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				ot.Number = value.String
			}
		}
	}
	return nil
}

// QueryOrders queries the "orders" edge of the OrderTaxation entity.
func (ot *OrderTaxation) QueryOrders() *OrderQuery {
	return (&OrderTaxationClient{config: ot.config}).QueryOrders(ot)
}

// Update returns a builder for updating this OrderTaxation.
// Note that you need to call OrderTaxation.Unwrap() before calling this method if this OrderTaxation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ot *OrderTaxation) Update() *OrderTaxationUpdateOne {
	return (&OrderTaxationClient{config: ot.config}).UpdateOne(ot)
}

// Unwrap unwraps the OrderTaxation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ot *OrderTaxation) Unwrap() *OrderTaxation {
	_tx, ok := ot.config.driver.(*txDriver)
	if !ok {
		panic("gen: OrderTaxation is not a transactional entity")
	}
	ot.config.driver = _tx.drv
	return ot
}

// String implements the fmt.Stringer.
func (ot *OrderTaxation) String() string {
	var builder strings.Builder
	builder.WriteString("OrderTaxation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ot.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ot.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ot.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ot.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", ot.OrderID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(ot.Type)
	builder.WriteString(", ")
	builder.WriteString("country_code=")
	builder.WriteString(ot.CountryCode)
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(ot.Number)
	builder.WriteByte(')')
	return builder.String()
}

// OrderTaxations is a parsable slice of OrderTaxation.
type OrderTaxations []*OrderTaxation

func (ot OrderTaxations) config(cfg config) {
	for _i := range ot {
		ot[_i].config = cfg
	}
}