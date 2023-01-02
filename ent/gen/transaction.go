// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
	"github.com/heshaofeng1991/entgo/ent/gen/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
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
	// TransactionType holds the value of the "transaction_type" field.
	TransactionType string `json:"transaction_type,omitempty"`
	// TransactionAmount holds the value of the "transaction_amount" field.
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance float64 `json:"balance,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int8 `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int8 `json:"updated_by,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges              TransactionEdges `json:"edges"`
	transaction_tenant *int64
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Details holds the value of the details edge.
	Details []*TransactionDetail `json:"details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) TenantOrErr() (*Tenant, error) {
	if e.loadedTypes[0] {
		if e.Tenant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tenant.Label}
		}
		return e.Tenant, nil
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// DetailsOrErr returns the Details value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) DetailsOrErr() ([]*TransactionDetail, error) {
	if e.loadedTypes[1] {
		return e.Details, nil
	}
	return nil, &NotLoadedError{edge: "details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldTransactionAmount, transaction.FieldBalance:
			values[i] = new(sql.NullFloat64)
		case transaction.FieldID, transaction.FieldOrderID, transaction.FieldStatus, transaction.FieldCreatedBy, transaction.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case transaction.FieldTransactionType, transaction.FieldRemark:
			values[i] = new(sql.NullString)
		case transaction.FieldCreatedAt, transaction.FieldUpdatedAt, transaction.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case transaction.ForeignKeys[0]: // transaction_tenant
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int64(value.Int64)
		case transaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case transaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case transaction.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = value.Time
			}
		case transaction.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				t.OrderID = value.Int64
			}
		case transaction.FieldTransactionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_type", values[i])
			} else if value.Valid {
				t.TransactionType = value.String
			}
		case transaction.FieldTransactionAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_amount", values[i])
			} else if value.Valid {
				t.TransactionAmount = value.Float64
			}
		case transaction.FieldBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				t.Balance = value.Float64
			}
		case transaction.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				t.Remark = value.String
			}
		case transaction.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = int8(value.Int64)
			}
		case transaction.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				t.CreatedBy = int8(value.Int64)
			}
		case transaction.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				t.UpdatedBy = int8(value.Int64)
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field transaction_tenant", value)
			} else if value.Valid {
				t.transaction_tenant = new(int64)
				*t.transaction_tenant = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryTenant queries the "tenant" edge of the Transaction entity.
func (t *Transaction) QueryTenant() *TenantQuery {
	return (&TransactionClient{config: t.config}).QueryTenant(t)
}

// QueryDetails queries the "details" edge of the Transaction entity.
func (t *Transaction) QueryDetails() *TransactionDetailQuery {
	return (&TransactionClient{config: t.config}).QueryDetails(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return (&TransactionClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("gen: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(t.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", t.OrderID))
	builder.WriteString(", ")
	builder.WriteString("transaction_type=")
	builder.WriteString(t.TransactionType)
	builder.WriteString(", ")
	builder.WriteString("transaction_amount=")
	builder.WriteString(fmt.Sprintf("%v", t.TransactionAmount))
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", t.Balance))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(t.Remark)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatedBy))
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction

func (t Transactions) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}