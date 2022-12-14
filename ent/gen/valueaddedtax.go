// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/valueaddedtax"
)

// ValueAddedTax is the model entity for the ValueAddedTax schema.
type ValueAddedTax struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode string `json:"country_code,omitempty"`
	// 标准税率 12 => 12%
	StandardRate float64 `json:"standard_rate,omitempty"`
	// 没有 IOSS 的税率 14 => 14%
	WithoutIossRate float64 `json:"without_ioss_rate,omitempty"`
	// 免征额,超出才要计算税费
	ExemptionInUsd float64 `json:"exemption_in_usd,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ValueAddedTax) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case valueaddedtax.FieldStandardRate, valueaddedtax.FieldWithoutIossRate, valueaddedtax.FieldExemptionInUsd:
			values[i] = new(sql.NullFloat64)
		case valueaddedtax.FieldID:
			values[i] = new(sql.NullInt64)
		case valueaddedtax.FieldCountryCode:
			values[i] = new(sql.NullString)
		case valueaddedtax.FieldCreatedAt, valueaddedtax.FieldUpdatedAt, valueaddedtax.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ValueAddedTax", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ValueAddedTax fields.
func (vat *ValueAddedTax) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case valueaddedtax.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vat.ID = int64(value.Int64)
		case valueaddedtax.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vat.CreatedAt = value.Time
			}
		case valueaddedtax.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vat.UpdatedAt = value.Time
			}
		case valueaddedtax.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				vat.DeletedAt = value.Time
			}
		case valueaddedtax.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				vat.CountryCode = value.String
			}
		case valueaddedtax.FieldStandardRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field standard_rate", values[i])
			} else if value.Valid {
				vat.StandardRate = value.Float64
			}
		case valueaddedtax.FieldWithoutIossRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field without_ioss_rate", values[i])
			} else if value.Valid {
				vat.WithoutIossRate = value.Float64
			}
		case valueaddedtax.FieldExemptionInUsd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field exemption_in_usd", values[i])
			} else if value.Valid {
				vat.ExemptionInUsd = value.Float64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ValueAddedTax.
// Note that you need to call ValueAddedTax.Unwrap() before calling this method if this ValueAddedTax
// was returned from a transaction, and the transaction was committed or rolled back.
func (vat *ValueAddedTax) Update() *ValueAddedTaxUpdateOne {
	return (&ValueAddedTaxClient{config: vat.config}).UpdateOne(vat)
}

// Unwrap unwraps the ValueAddedTax entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vat *ValueAddedTax) Unwrap() *ValueAddedTax {
	_tx, ok := vat.config.driver.(*txDriver)
	if !ok {
		panic("gen: ValueAddedTax is not a transactional entity")
	}
	vat.config.driver = _tx.drv
	return vat
}

// String implements the fmt.Stringer.
func (vat *ValueAddedTax) String() string {
	var builder strings.Builder
	builder.WriteString("ValueAddedTax(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vat.ID))
	builder.WriteString("created_at=")
	builder.WriteString(vat.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vat.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(vat.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("country_code=")
	builder.WriteString(vat.CountryCode)
	builder.WriteString(", ")
	builder.WriteString("standard_rate=")
	builder.WriteString(fmt.Sprintf("%v", vat.StandardRate))
	builder.WriteString(", ")
	builder.WriteString("without_ioss_rate=")
	builder.WriteString(fmt.Sprintf("%v", vat.WithoutIossRate))
	builder.WriteString(", ")
	builder.WriteString("exemption_in_usd=")
	builder.WriteString(fmt.Sprintf("%v", vat.ExemptionInUsd))
	builder.WriteByte(')')
	return builder.String()
}

// ValueAddedTaxes is a parsable slice of ValueAddedTax.
type ValueAddedTaxes []*ValueAddedTax

func (vat ValueAddedTaxes) config(cfg config) {
	for _i := range vat {
		vat[_i].config = cfg
	}
}
