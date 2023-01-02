// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/taskschedule"
)

// TaskSchedule is the model entity for the TaskSchedule schema.
type TaskSchedule struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Platform holds the value of the "platform" field.
	Platform string `json:"platform,omitempty"`
	// StoreID holds the value of the "store_id" field.
	StoreID int64 `json:"store_id,omitempty"`
	// FuncName holds the value of the "func_name" field.
	FuncName string `json:"func_name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked bool `json:"locked,omitempty"`
	// LockedTimes holds the value of the "locked_times" field.
	LockedTimes int64 `json:"locked_times,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// LastAccessAt holds the value of the "last_access_at" field.
	LastAccessAt time.Time `json:"last_access_at,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskSchedule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case taskschedule.FieldLocked:
			values[i] = new(sql.NullBool)
		case taskschedule.FieldID, taskschedule.FieldStoreID, taskschedule.FieldLockedTimes, taskschedule.FieldStatus:
			values[i] = new(sql.NullInt64)
		case taskschedule.FieldPlatform, taskschedule.FieldFuncName, taskschedule.FieldDescription, taskschedule.FieldRemark:
			values[i] = new(sql.NullString)
		case taskschedule.FieldCreatedAt, taskschedule.FieldUpdatedAt, taskschedule.FieldDeletedAt, taskschedule.FieldLastAccessAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaskSchedule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskSchedule fields.
func (ts *TaskSchedule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case taskschedule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int64(value.Int64)
		case taskschedule.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ts.CreatedAt = value.Time
			}
		case taskschedule.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ts.UpdatedAt = value.Time
			}
		case taskschedule.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ts.DeletedAt = value.Time
			}
		case taskschedule.FieldPlatform:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform", values[i])
			} else if value.Valid {
				ts.Platform = value.String
			}
		case taskschedule.FieldStoreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field store_id", values[i])
			} else if value.Valid {
				ts.StoreID = value.Int64
			}
		case taskschedule.FieldFuncName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field func_name", values[i])
			} else if value.Valid {
				ts.FuncName = value.String
			}
		case taskschedule.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ts.Description = value.String
			}
		case taskschedule.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				ts.Locked = value.Bool
			}
		case taskschedule.FieldLockedTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field locked_times", values[i])
			} else if value.Valid {
				ts.LockedTimes = value.Int64
			}
		case taskschedule.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				ts.Remark = value.String
			}
		case taskschedule.FieldLastAccessAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_access_at", values[i])
			} else if value.Valid {
				ts.LastAccessAt = value.Time
			}
		case taskschedule.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ts.Status = int8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TaskSchedule.
// Note that you need to call TaskSchedule.Unwrap() before calling this method if this TaskSchedule
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *TaskSchedule) Update() *TaskScheduleUpdateOne {
	return (&TaskScheduleClient{config: ts.config}).UpdateOne(ts)
}

// Unwrap unwraps the TaskSchedule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *TaskSchedule) Unwrap() *TaskSchedule {
	_tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("gen: TaskSchedule is not a transactional entity")
	}
	ts.config.driver = _tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *TaskSchedule) String() string {
	var builder strings.Builder
	builder.WriteString("TaskSchedule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ts.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ts.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ts.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ts.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("platform=")
	builder.WriteString(ts.Platform)
	builder.WriteString(", ")
	builder.WriteString("store_id=")
	builder.WriteString(fmt.Sprintf("%v", ts.StoreID))
	builder.WriteString(", ")
	builder.WriteString("func_name=")
	builder.WriteString(ts.FuncName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ts.Description)
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", ts.Locked))
	builder.WriteString(", ")
	builder.WriteString("locked_times=")
	builder.WriteString(fmt.Sprintf("%v", ts.LockedTimes))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(ts.Remark)
	builder.WriteString(", ")
	builder.WriteString("last_access_at=")
	builder.WriteString(ts.LastAccessAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ts.Status))
	builder.WriteByte(')')
	return builder.String()
}

// TaskSchedules is a parsable slice of TaskSchedule.
type TaskSchedules []*TaskSchedule

func (ts TaskSchedules) config(cfg config) {
	for _i := range ts {
		ts[_i].config = cfg
	}
}
