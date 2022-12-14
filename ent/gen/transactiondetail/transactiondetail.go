// Code generated by ent, DO NOT EDIT.

package transactiondetail

import (
	"time"
)

const (
	// Label holds the string label denoting the transactiondetail type in the database.
	Label = "transaction_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldTransactionType holds the string denoting the transaction_type field in the database.
	FieldTransactionType = "transaction_type"
	// FieldDeliveryCost holds the string denoting the delivery_cost field in the database.
	FieldDeliveryCost = "delivery_cost"
	// FieldMiscFee holds the string denoting the misc_fee field in the database.
	FieldMiscFee = "misc_fee"
	// FieldFuelFee holds the string denoting the fuel_fee field in the database.
	FieldFuelFee = "fuel_fee"
	// FieldRegistrationFee holds the string denoting the registration_fee field in the database.
	FieldRegistrationFee = "registration_fee"
	// FieldProcessingFee holds the string denoting the processing_fee field in the database.
	FieldProcessingFee = "processing_fee"
	// FieldPackageFee holds the string denoting the package_fee field in the database.
	FieldPackageFee = "package_fee"
	// FieldHandlingFee holds the string denoting the handling_fee field in the database.
	FieldHandlingFee = "handling_fee"
	// FieldVat holds the string denoting the vat field in the database.
	FieldVat = "vat"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// Table holds the table name of the transactiondetail in the database.
	Table = "transaction_details"
)

// Columns holds all SQL columns for transactiondetail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldOrderID,
	FieldTransactionID,
	FieldTransactionType,
	FieldDeliveryCost,
	FieldMiscFee,
	FieldFuelFee,
	FieldRegistrationFee,
	FieldProcessingFee,
	FieldPackageFee,
	FieldHandlingFee,
	FieldVat,
	FieldAmount,
	FieldWeight,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transaction_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"transaction_details",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultTransactionType holds the default value on creation for the "transaction_type" field.
	DefaultTransactionType string
	// DefaultDeliveryCost holds the default value on creation for the "delivery_cost" field.
	DefaultDeliveryCost float64
	// DefaultMiscFee holds the default value on creation for the "misc_fee" field.
	DefaultMiscFee float64
	// DefaultFuelFee holds the default value on creation for the "fuel_fee" field.
	DefaultFuelFee float64
	// DefaultRegistrationFee holds the default value on creation for the "registration_fee" field.
	DefaultRegistrationFee float64
	// DefaultProcessingFee holds the default value on creation for the "processing_fee" field.
	DefaultProcessingFee float64
	// DefaultPackageFee holds the default value on creation for the "package_fee" field.
	DefaultPackageFee float64
	// DefaultHandlingFee holds the default value on creation for the "handling_fee" field.
	DefaultHandlingFee float64
	// DefaultVat holds the default value on creation for the "vat" field.
	DefaultVat float64
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount float64
	// DefaultWeight holds the default value on creation for the "weight" field.
	DefaultWeight int
)
