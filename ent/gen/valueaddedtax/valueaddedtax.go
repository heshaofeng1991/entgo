// Code generated by ent, DO NOT EDIT.

package valueaddedtax

import (
	"time"
)

const (
	// Label holds the string label denoting the valueaddedtax type in the database.
	Label = "value_added_tax"
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
	// FieldStandardRate holds the string denoting the standard_rate field in the database.
	FieldStandardRate = "standard_rate"
	// FieldWithoutIossRate holds the string denoting the without_ioss_rate field in the database.
	FieldWithoutIossRate = "without_ioss_rate"
	// FieldExemptionInUsd holds the string denoting the exemption_in_usd field in the database.
	FieldExemptionInUsd = "exemption_in_usd"
	// Table holds the table name of the valueaddedtax in the database.
	Table = "value_added_taxes"
)

// Columns holds all SQL columns for valueaddedtax fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCountryCode,
	FieldStandardRate,
	FieldWithoutIossRate,
	FieldExemptionInUsd,
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
	// DefaultCountryCode holds the default value on creation for the "country_code" field.
	DefaultCountryCode string
	// DefaultStandardRate holds the default value on creation for the "standard_rate" field.
	DefaultStandardRate float64
	// DefaultWithoutIossRate holds the default value on creation for the "without_ioss_rate" field.
	DefaultWithoutIossRate float64
	// DefaultExemptionInUsd holds the default value on creation for the "exemption_in_usd" field.
	DefaultExemptionInUsd float64
)
