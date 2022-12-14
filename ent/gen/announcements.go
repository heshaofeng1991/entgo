// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/announcements"
	"github.com/heshaofeng1991/entgo/ent/gen/user"
)

// Announcements is the model entity for the Announcements schema.
type Announcements struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Index holds the value of the "index" field.
	Index int `json:"index,omitempty"`
	// CreateBy holds the value of the "create_by" field.
	CreateBy int64 `json:"create_by,omitempty"`
	// Expiration holds the value of the "expiration" field.
	Expiration time.Time `json:"expiration,omitempty"`
	// EffectiveTime holds the value of the "effective_time" field.
	EffectiveTime time.Time `json:"effective_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnnouncementsQuery when eager-loading is set.
	Edges AnnouncementsEdges `json:"edges"`
}

// AnnouncementsEdges holds the relations/edges for other nodes in the graph.
type AnnouncementsEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AnnouncementsEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Announcements) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case announcements.FieldID, announcements.FieldStatus, announcements.FieldIndex, announcements.FieldCreateBy:
			values[i] = new(sql.NullInt64)
		case announcements.FieldTitle, announcements.FieldContent:
			values[i] = new(sql.NullString)
		case announcements.FieldCreatedAt, announcements.FieldUpdatedAt, announcements.FieldDeletedAt, announcements.FieldExpiration, announcements.FieldEffectiveTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Announcements", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Announcements fields.
func (a *Announcements) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case announcements.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int64(value.Int64)
		case announcements.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case announcements.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case announcements.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case announcements.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case announcements.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				a.Content = value.String
			}
		case announcements.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = int(value.Int64)
			}
		case announcements.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				a.Index = int(value.Int64)
			}
		case announcements.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				a.CreateBy = value.Int64
			}
		case announcements.FieldExpiration:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiration", values[i])
			} else if value.Valid {
				a.Expiration = value.Time
			}
		case announcements.FieldEffectiveTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field effective_time", values[i])
			} else if value.Valid {
				a.EffectiveTime = value.Time
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Announcements entity.
func (a *Announcements) QueryUsers() *UserQuery {
	return (&AnnouncementsClient{config: a.config}).QueryUsers(a)
}

// Update returns a builder for updating this Announcements.
// Note that you need to call Announcements.Unwrap() before calling this method if this Announcements
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Announcements) Update() *AnnouncementsUpdateOne {
	return (&AnnouncementsClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Announcements entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Announcements) Unwrap() *Announcements {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("gen: Announcements is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Announcements) String() string {
	var builder strings.Builder
	builder.WriteString("Announcements(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(a.Content)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", a.Index))
	builder.WriteString(", ")
	builder.WriteString("create_by=")
	builder.WriteString(fmt.Sprintf("%v", a.CreateBy))
	builder.WriteString(", ")
	builder.WriteString("expiration=")
	builder.WriteString(a.Expiration.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("effective_time=")
	builder.WriteString(a.EffectiveTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AnnouncementsSlice is a parsable slice of Announcements.
type AnnouncementsSlice []*Announcements

func (a AnnouncementsSlice) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
