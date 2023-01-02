// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/channeloption"
)

// ChannelOption is the model entity for the ChannelOption schema.
type ChannelOption struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int64 `json:"order_id,omitempty"`
	// ChannelID holds the value of the "channel_id" field.
	ChannelID int64 `json:"channel_id,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode string `json:"country_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChannelOption) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case channeloption.FieldID, channeloption.FieldOrderID, channeloption.FieldChannelID:
			values[i] = new(sql.NullInt64)
		case channeloption.FieldCountryCode:
			values[i] = new(sql.NullString)
		case channeloption.FieldCreatedAt, channeloption.FieldUpdatedAt, channeloption.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChannelOption", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChannelOption fields.
func (co *ChannelOption) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case channeloption.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			co.ID = int64(value.Int64)
		case channeloption.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				co.OrderID = value.Int64
			}
		case channeloption.FieldChannelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field channel_id", values[i])
			} else if value.Valid {
				co.ChannelID = value.Int64
			}
		case channeloption.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				co.CountryCode = value.String
			}
		case channeloption.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				co.CreatedAt = value.Time
			}
		case channeloption.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				co.UpdatedAt = value.Time
			}
		case channeloption.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				co.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ChannelOption.
// Note that you need to call ChannelOption.Unwrap() before calling this method if this ChannelOption
// was returned from a transaction, and the transaction was committed or rolled back.
func (co *ChannelOption) Update() *ChannelOptionUpdateOne {
	return (&ChannelOptionClient{config: co.config}).UpdateOne(co)
}

// Unwrap unwraps the ChannelOption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (co *ChannelOption) Unwrap() *ChannelOption {
	_tx, ok := co.config.driver.(*txDriver)
	if !ok {
		panic("gen: ChannelOption is not a transactional entity")
	}
	co.config.driver = _tx.drv
	return co
}

// String implements the fmt.Stringer.
func (co *ChannelOption) String() string {
	var builder strings.Builder
	builder.WriteString("ChannelOption(")
	builder.WriteString(fmt.Sprintf("id=%v, ", co.ID))
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", co.OrderID))
	builder.WriteString(", ")
	builder.WriteString("channel_id=")
	builder.WriteString(fmt.Sprintf("%v", co.ChannelID))
	builder.WriteString(", ")
	builder.WriteString("country_code=")
	builder.WriteString(co.CountryCode)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(co.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(co.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(co.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ChannelOptions is a parsable slice of ChannelOption.
type ChannelOptions []*ChannelOption

func (co ChannelOptions) config(cfg config) {
	for _i := range co {
		co[_i].config = cfg
	}
}
