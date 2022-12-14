// Code generated by ent, DO NOT EDIT.

package channeloption

import (
	"time"
)

const (
	// Label holds the string label denoting the channeloption type in the database.
	Label = "channel_option"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the channeloption in the database.
	Table = "channel_options"
)

// Columns holds all SQL columns for channeloption fields.
var Columns = []string{
	FieldID,
	FieldOrderID,
	FieldChannelID,
	FieldCountryCode,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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
)
