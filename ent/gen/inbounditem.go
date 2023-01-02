// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/inbound"
	"github.com/heshaofeng1991/entgo/ent/gen/inbounditem"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
)

// InboundItem is the model entity for the InboundItem schema.
type InboundItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 入库单ID
	InboundID int64 `json:"inbound_id,omitempty"`
	// 产品ID
	ProductID int64 `json:"product_id,omitempty"`
	// 产品名称
	ProductName string `json:"product_name,omitempty"`
	// 产品SKU
	Sku string `json:"sku,omitempty"`
	// 产品条码
	Barcode string `json:"barcode,omitempty"`
	// 产品数量
	Qty int `json:"qty,omitempty"`
	// 状态
	Status int8 `json:"status,omitempty"`
	// CustomerCode holds the value of the "customer_code" field.
	CustomerCode string `json:"customer_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InboundItemQuery when eager-loading is set.
	Edges               InboundItemEdges `json:"edges"`
	inbound_item_tenant *int64
}

// InboundItemEdges holds the relations/edges for other nodes in the graph.
type InboundItemEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Inbounds holds the value of the inbounds edge.
	Inbounds *Inbound `json:"inbounds,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InboundItemEdges) TenantOrErr() (*Tenant, error) {
	if e.loadedTypes[0] {
		if e.Tenant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tenant.Label}
		}
		return e.Tenant, nil
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// InboundsOrErr returns the Inbounds value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InboundItemEdges) InboundsOrErr() (*Inbound, error) {
	if e.loadedTypes[1] {
		if e.Inbounds == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: inbound.Label}
		}
		return e.Inbounds, nil
	}
	return nil, &NotLoadedError{edge: "inbounds"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InboundItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case inbounditem.FieldID, inbounditem.FieldInboundID, inbounditem.FieldProductID, inbounditem.FieldQty, inbounditem.FieldStatus:
			values[i] = new(sql.NullInt64)
		case inbounditem.FieldProductName, inbounditem.FieldSku, inbounditem.FieldBarcode, inbounditem.FieldCustomerCode:
			values[i] = new(sql.NullString)
		case inbounditem.FieldCreatedAt, inbounditem.FieldUpdatedAt, inbounditem.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case inbounditem.ForeignKeys[0]: // inbound_item_tenant
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type InboundItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InboundItem fields.
func (ii *InboundItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case inbounditem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ii.ID = int64(value.Int64)
		case inbounditem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ii.CreatedAt = value.Time
			}
		case inbounditem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ii.UpdatedAt = value.Time
			}
		case inbounditem.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ii.DeletedAt = value.Time
			}
		case inbounditem.FieldInboundID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field inbound_id", values[i])
			} else if value.Valid {
				ii.InboundID = value.Int64
			}
		case inbounditem.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				ii.ProductID = value.Int64
			}
		case inbounditem.FieldProductName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_name", values[i])
			} else if value.Valid {
				ii.ProductName = value.String
			}
		case inbounditem.FieldSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sku", values[i])
			} else if value.Valid {
				ii.Sku = value.String
			}
		case inbounditem.FieldBarcode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field barcode", values[i])
			} else if value.Valid {
				ii.Barcode = value.String
			}
		case inbounditem.FieldQty:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field qty", values[i])
			} else if value.Valid {
				ii.Qty = int(value.Int64)
			}
		case inbounditem.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ii.Status = int8(value.Int64)
			}
		case inbounditem.FieldCustomerCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_code", values[i])
			} else if value.Valid {
				ii.CustomerCode = value.String
			}
		case inbounditem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field inbound_item_tenant", value)
			} else if value.Valid {
				ii.inbound_item_tenant = new(int64)
				*ii.inbound_item_tenant = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryTenant queries the "tenant" edge of the InboundItem entity.
func (ii *InboundItem) QueryTenant() *TenantQuery {
	return (&InboundItemClient{config: ii.config}).QueryTenant(ii)
}

// QueryInbounds queries the "inbounds" edge of the InboundItem entity.
func (ii *InboundItem) QueryInbounds() *InboundQuery {
	return (&InboundItemClient{config: ii.config}).QueryInbounds(ii)
}

// Update returns a builder for updating this InboundItem.
// Note that you need to call InboundItem.Unwrap() before calling this method if this InboundItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (ii *InboundItem) Update() *InboundItemUpdateOne {
	return (&InboundItemClient{config: ii.config}).UpdateOne(ii)
}

// Unwrap unwraps the InboundItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ii *InboundItem) Unwrap() *InboundItem {
	_tx, ok := ii.config.driver.(*txDriver)
	if !ok {
		panic("gen: InboundItem is not a transactional entity")
	}
	ii.config.driver = _tx.drv
	return ii
}

// String implements the fmt.Stringer.
func (ii *InboundItem) String() string {
	var builder strings.Builder
	builder.WriteString("InboundItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ii.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ii.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ii.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ii.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("inbound_id=")
	builder.WriteString(fmt.Sprintf("%v", ii.InboundID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", ii.ProductID))
	builder.WriteString(", ")
	builder.WriteString("product_name=")
	builder.WriteString(ii.ProductName)
	builder.WriteString(", ")
	builder.WriteString("sku=")
	builder.WriteString(ii.Sku)
	builder.WriteString(", ")
	builder.WriteString("barcode=")
	builder.WriteString(ii.Barcode)
	builder.WriteString(", ")
	builder.WriteString("qty=")
	builder.WriteString(fmt.Sprintf("%v", ii.Qty))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ii.Status))
	builder.WriteString(", ")
	builder.WriteString("customer_code=")
	builder.WriteString(ii.CustomerCode)
	builder.WriteByte(')')
	return builder.String()
}

// InboundItems is a parsable slice of InboundItem.
type InboundItems []*InboundItem

func (ii InboundItems) config(cfg config) {
	for _i := range ii {
		ii[_i].config = cfg
	}
}
