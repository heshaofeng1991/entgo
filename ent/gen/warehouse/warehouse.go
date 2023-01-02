// Code generated by ent, DO NOT EDIT.

package warehouse

import (
	"time"
)

const (
	// Label holds the string label denoting the warehouse type in the database.
	Label = "warehouse"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldQuicktronCode holds the string denoting the quicktron_code field in the database.
	FieldQuicktronCode = "quicktron_code"
	// FieldEnableQuicktron holds the string denoting the enable_quicktron field in the database.
	FieldEnableQuicktron = "enable_quicktron"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldAddress1 holds the string denoting the address1 field in the database.
	FieldAddress1 = "address1"
	// FieldAddress2 holds the string denoting the address2 field in the database.
	FieldAddress2 = "address2"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldCountryName holds the string denoting the country_name field in the database.
	FieldCountryName = "country_name"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldZipCode holds the string denoting the zip_code field in the database.
	FieldZipCode = "zip_code"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the warehouse in the database.
	Table = "warehouses"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "warehouse_id"
)

// Columns holds all SQL columns for warehouse fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCode,
	FieldQuicktronCode,
	FieldEnableQuicktron,
	FieldName,
	FieldCompany,
	FieldFirstName,
	FieldLastName,
	FieldAddress1,
	FieldAddress2,
	FieldCountryCode,
	FieldCountryName,
	FieldProvince,
	FieldCity,
	FieldZipCode,
	FieldPhone,
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
)
