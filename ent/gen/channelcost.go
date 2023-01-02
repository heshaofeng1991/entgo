// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/channel"
	"github.com/heshaofeng1991/entgo/ent/gen/channelcost"
)

// ChannelCost is the model entity for the ChannelCost schema.
type ChannelCost struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// ChannelCostBatchID holds the value of the "channel_cost_batch_id" field.
	ChannelCostBatchID int64 `json:"channel_cost_batch_id,omitempty"`
	// ChannelID holds the value of the "channel_id" field.
	ChannelID int64 `json:"channel_id,omitempty"`
	// 报价计算模型
	Mode int8 `json:"mode,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode string `json:"country_code,omitempty"`
	// Zone holds the value of the "zone" field.
	Zone string `json:"zone,omitempty"`
	// gram
	StartWeight int `json:"start_weight,omitempty"`
	// gram
	EndWeight int `json:"end_weight,omitempty"`
	// gram
	FirstWeight int `json:"first_weight,omitempty"`
	// FirstWeightFee holds the value of the "first_weight_fee" field.
	FirstWeightFee float64 `json:"first_weight_fee,omitempty"`
	// 单位重量
	UnitWeight int `json:"unit_weight,omitempty"`
	// UnitWeightFee holds the value of the "unit_weight_fee" field.
	UnitWeightFee float64 `json:"unit_weight_fee,omitempty"`
	// FuelFee holds the value of the "fuel_fee" field.
	FuelFee float64 `json:"fuel_fee,omitempty"`
	// ProcessingFee holds the value of the "processing_fee" field.
	ProcessingFee float64 `json:"processing_fee,omitempty"`
	// RegistrationFee holds the value of the "registration_fee" field.
	RegistrationFee float64 `json:"registration_fee,omitempty"`
	// MiscFee holds the value of the "misc_fee" field.
	MiscFee float64 `json:"misc_fee,omitempty"`
	// MinNormalDays holds the value of the "min_normal_days" field.
	MinNormalDays int `json:"min_normal_days,omitempty"`
	// MaxNormalDays holds the value of the "max_normal_days" field.
	MaxNormalDays int `json:"max_normal_days,omitempty"`
	// 0=>启用, 1=>不启用
	Status int8 `json:"status,omitempty"`
	// AverageDays holds the value of the "average_days" field.
	AverageDays int `json:"average_days,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChannelCostQuery when eager-loading is set.
	Edges ChannelCostEdges `json:"edges"`
}

// ChannelCostEdges holds the relations/edges for other nodes in the graph.
type ChannelCostEdges struct {
	// Channels holds the value of the channels edge.
	Channels *Channel `json:"channels,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ChannelsOrErr returns the Channels value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChannelCostEdges) ChannelsOrErr() (*Channel, error) {
	if e.loadedTypes[0] {
		if e.Channels == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: channel.Label}
		}
		return e.Channels, nil
	}
	return nil, &NotLoadedError{edge: "channels"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChannelCost) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case channelcost.FieldFirstWeightFee, channelcost.FieldUnitWeightFee, channelcost.FieldFuelFee, channelcost.FieldProcessingFee, channelcost.FieldRegistrationFee, channelcost.FieldMiscFee:
			values[i] = new(sql.NullFloat64)
		case channelcost.FieldID, channelcost.FieldChannelCostBatchID, channelcost.FieldChannelID, channelcost.FieldMode, channelcost.FieldStartWeight, channelcost.FieldEndWeight, channelcost.FieldFirstWeight, channelcost.FieldUnitWeight, channelcost.FieldMinNormalDays, channelcost.FieldMaxNormalDays, channelcost.FieldStatus, channelcost.FieldAverageDays:
			values[i] = new(sql.NullInt64)
		case channelcost.FieldCountryCode, channelcost.FieldZone:
			values[i] = new(sql.NullString)
		case channelcost.FieldCreatedAt, channelcost.FieldUpdatedAt, channelcost.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChannelCost", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChannelCost fields.
func (cc *ChannelCost) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case channelcost.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = int64(value.Int64)
		case channelcost.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cc.CreatedAt = value.Time
			}
		case channelcost.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cc.UpdatedAt = value.Time
			}
		case channelcost.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cc.DeletedAt = value.Time
			}
		case channelcost.FieldChannelCostBatchID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field channel_cost_batch_id", values[i])
			} else if value.Valid {
				cc.ChannelCostBatchID = value.Int64
			}
		case channelcost.FieldChannelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field channel_id", values[i])
			} else if value.Valid {
				cc.ChannelID = value.Int64
			}
		case channelcost.FieldMode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				cc.Mode = int8(value.Int64)
			}
		case channelcost.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				cc.CountryCode = value.String
			}
		case channelcost.FieldZone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zone", values[i])
			} else if value.Valid {
				cc.Zone = value.String
			}
		case channelcost.FieldStartWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_weight", values[i])
			} else if value.Valid {
				cc.StartWeight = int(value.Int64)
			}
		case channelcost.FieldEndWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_weight", values[i])
			} else if value.Valid {
				cc.EndWeight = int(value.Int64)
			}
		case channelcost.FieldFirstWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field first_weight", values[i])
			} else if value.Valid {
				cc.FirstWeight = int(value.Int64)
			}
		case channelcost.FieldFirstWeightFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field first_weight_fee", values[i])
			} else if value.Valid {
				cc.FirstWeightFee = value.Float64
			}
		case channelcost.FieldUnitWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_weight", values[i])
			} else if value.Valid {
				cc.UnitWeight = int(value.Int64)
			}
		case channelcost.FieldUnitWeightFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_weight_fee", values[i])
			} else if value.Valid {
				cc.UnitWeightFee = value.Float64
			}
		case channelcost.FieldFuelFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fuel_fee", values[i])
			} else if value.Valid {
				cc.FuelFee = value.Float64
			}
		case channelcost.FieldProcessingFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field processing_fee", values[i])
			} else if value.Valid {
				cc.ProcessingFee = value.Float64
			}
		case channelcost.FieldRegistrationFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field registration_fee", values[i])
			} else if value.Valid {
				cc.RegistrationFee = value.Float64
			}
		case channelcost.FieldMiscFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field misc_fee", values[i])
			} else if value.Valid {
				cc.MiscFee = value.Float64
			}
		case channelcost.FieldMinNormalDays:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_normal_days", values[i])
			} else if value.Valid {
				cc.MinNormalDays = int(value.Int64)
			}
		case channelcost.FieldMaxNormalDays:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_normal_days", values[i])
			} else if value.Valid {
				cc.MaxNormalDays = int(value.Int64)
			}
		case channelcost.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cc.Status = int8(value.Int64)
			}
		case channelcost.FieldAverageDays:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field average_days", values[i])
			} else if value.Valid {
				cc.AverageDays = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryChannels queries the "channels" edge of the ChannelCost entity.
func (cc *ChannelCost) QueryChannels() *ChannelQuery {
	return (&ChannelCostClient{config: cc.config}).QueryChannels(cc)
}

// Update returns a builder for updating this ChannelCost.
// Note that you need to call ChannelCost.Unwrap() before calling this method if this ChannelCost
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *ChannelCost) Update() *ChannelCostUpdateOne {
	return (&ChannelCostClient{config: cc.config}).UpdateOne(cc)
}

// Unwrap unwraps the ChannelCost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *ChannelCost) Unwrap() *ChannelCost {
	_tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("gen: ChannelCost is not a transactional entity")
	}
	cc.config.driver = _tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *ChannelCost) String() string {
	var builder strings.Builder
	builder.WriteString("ChannelCost(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(cc.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("channel_cost_batch_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.ChannelCostBatchID))
	builder.WriteString(", ")
	builder.WriteString("channel_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.ChannelID))
	builder.WriteString(", ")
	builder.WriteString("mode=")
	builder.WriteString(fmt.Sprintf("%v", cc.Mode))
	builder.WriteString(", ")
	builder.WriteString("country_code=")
	builder.WriteString(cc.CountryCode)
	builder.WriteString(", ")
	builder.WriteString("zone=")
	builder.WriteString(cc.Zone)
	builder.WriteString(", ")
	builder.WriteString("start_weight=")
	builder.WriteString(fmt.Sprintf("%v", cc.StartWeight))
	builder.WriteString(", ")
	builder.WriteString("end_weight=")
	builder.WriteString(fmt.Sprintf("%v", cc.EndWeight))
	builder.WriteString(", ")
	builder.WriteString("first_weight=")
	builder.WriteString(fmt.Sprintf("%v", cc.FirstWeight))
	builder.WriteString(", ")
	builder.WriteString("first_weight_fee=")
	builder.WriteString(fmt.Sprintf("%v", cc.FirstWeightFee))
	builder.WriteString(", ")
	builder.WriteString("unit_weight=")
	builder.WriteString(fmt.Sprintf("%v", cc.UnitWeight))
	builder.WriteString(", ")
	builder.WriteString("unit_weight_fee=")
	builder.WriteString(fmt.Sprintf("%v", cc.UnitWeightFee))
	builder.WriteString(", ")
	builder.WriteString("fuel_fee=")
	builder.WriteString(fmt.Sprintf("%v", cc.FuelFee))
	builder.WriteString(", ")
	builder.WriteString("processing_fee=")
	builder.WriteString(fmt.Sprintf("%v", cc.ProcessingFee))
	builder.WriteString(", ")
	builder.WriteString("registration_fee=")
	builder.WriteString(fmt.Sprintf("%v", cc.RegistrationFee))
	builder.WriteString(", ")
	builder.WriteString("misc_fee=")
	builder.WriteString(fmt.Sprintf("%v", cc.MiscFee))
	builder.WriteString(", ")
	builder.WriteString("min_normal_days=")
	builder.WriteString(fmt.Sprintf("%v", cc.MinNormalDays))
	builder.WriteString(", ")
	builder.WriteString("max_normal_days=")
	builder.WriteString(fmt.Sprintf("%v", cc.MaxNormalDays))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cc.Status))
	builder.WriteString(", ")
	builder.WriteString("average_days=")
	builder.WriteString(fmt.Sprintf("%v", cc.AverageDays))
	builder.WriteByte(')')
	return builder.String()
}

// ChannelCosts is a parsable slice of ChannelCost.
type ChannelCosts []*ChannelCost

func (cc ChannelCosts) config(cfg config) {
	for _i := range cc {
		cc[_i].config = cfg
	}
}
