// Code generated by ent, DO NOT EDIT.

package channelvolumefactor

import (
	"time"
)

const (
	// Label holds the string label denoting the channelvolumefactor type in the database.
	Label = "channel_volume_factor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldVolumeFactor holds the string denoting the volume_factor field in the database.
	FieldVolumeFactor = "volume_factor"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the channelvolumefactor in the database.
	Table = "channel_volume_factors"
)

// Columns holds all SQL columns for channelvolumefactor fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCountryCode,
	FieldChannelID,
	FieldVolumeFactor,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultVolumeFactor holds the default value on creation for the "volume_factor" field.
	DefaultVolumeFactor int
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
)
