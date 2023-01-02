// Code generated by ent, DO NOT EDIT.

package courierorder

import (
	"time"
)

const (
	// Label holds the string label denoting the courierorder type in the database.
	Label = "courier_order"
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
	// FieldOrderNumber holds the string denoting the order_number field in the database.
	FieldOrderNumber = "order_number"
	// FieldCourierPlatform holds the string denoting the courier_platform field in the database.
	FieldCourierPlatform = "courier_platform"
	// FieldShippingMethodCode holds the string denoting the shipping_method_code field in the database.
	FieldShippingMethodCode = "shipping_method_code"
	// FieldShippingMethodName holds the string denoting the shipping_method_name field in the database.
	FieldShippingMethodName = "shipping_method_name"
	// FieldTrackingURL holds the string denoting the tracking_url field in the database.
	FieldTrackingURL = "tracking_url"
	// FieldTrackingNumber holds the string denoting the tracking_number field in the database.
	FieldTrackingNumber = "tracking_number"
	// FieldWaybillNumber holds the string denoting the waybill_number field in the database.
	FieldWaybillNumber = "waybill_number"
	// FieldCourierOrderNumber holds the string denoting the courier_order_number field in the database.
	FieldCourierOrderNumber = "courier_order_number"
	// FieldShippingLabelURL holds the string denoting the shipping_label_url field in the database.
	FieldShippingLabelURL = "shipping_label_url"
	// FieldTotalItemsPrice holds the string denoting the total_items_price field in the database.
	FieldTotalItemsPrice = "total_items_price"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldPackageCount holds the string denoting the package_count field in the database.
	FieldPackageCount = "package_count"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldReceiverAddress holds the string denoting the receiver_address field in the database.
	FieldReceiverAddress = "receiver_address"
	// FieldSenderAddress holds the string denoting the sender_address field in the database.
	FieldSenderAddress = "sender_address"
	// FieldItems holds the string denoting the items field in the database.
	FieldItems = "items"
	// FieldRequestData holds the string denoting the request_data field in the database.
	FieldRequestData = "request_data"
	// FieldResponseData holds the string denoting the response_data field in the database.
	FieldResponseData = "response_data"
	// FieldResultCode holds the string denoting the result_code field in the database.
	FieldResultCode = "result_code"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldEnMessage holds the string denoting the en_message field in the database.
	FieldEnMessage = "en_message"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the courierorder in the database.
	Table = "courier_orders"
)

// Columns holds all SQL columns for courierorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldOrderID,
	FieldOrderNumber,
	FieldCourierPlatform,
	FieldShippingMethodCode,
	FieldShippingMethodName,
	FieldTrackingURL,
	FieldTrackingNumber,
	FieldWaybillNumber,
	FieldCourierOrderNumber,
	FieldShippingLabelURL,
	FieldTotalItemsPrice,
	FieldCurrency,
	FieldPackageCount,
	FieldWeight,
	FieldReceiverAddress,
	FieldSenderAddress,
	FieldItems,
	FieldRequestData,
	FieldResponseData,
	FieldResultCode,
	FieldMessage,
	FieldEnMessage,
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
	// DefaultCourierPlatform holds the default value on creation for the "courier_platform" field.
	DefaultCourierPlatform string
	// DefaultShippingMethodCode holds the default value on creation for the "shipping_method_code" field.
	DefaultShippingMethodCode string
	// DefaultShippingMethodName holds the default value on creation for the "shipping_method_name" field.
	DefaultShippingMethodName string
	// DefaultTrackingURL holds the default value on creation for the "tracking_url" field.
	DefaultTrackingURL string
	// DefaultTrackingNumber holds the default value on creation for the "tracking_number" field.
	DefaultTrackingNumber string
	// DefaultWaybillNumber holds the default value on creation for the "waybill_number" field.
	DefaultWaybillNumber string
	// DefaultCourierOrderNumber holds the default value on creation for the "courier_order_number" field.
	DefaultCourierOrderNumber string
	// DefaultShippingLabelURL holds the default value on creation for the "shipping_label_url" field.
	DefaultShippingLabelURL string
	// DefaultTotalItemsPrice holds the default value on creation for the "total_items_price" field.
	DefaultTotalItemsPrice float64
	// DefaultPackageCount holds the default value on creation for the "package_count" field.
	DefaultPackageCount int
	// DefaultWeight holds the default value on creation for the "weight" field.
	DefaultWeight int
	// DefaultResultCode holds the default value on creation for the "result_code" field.
	DefaultResultCode string
	// DefaultMessage holds the default value on creation for the "message" field.
	DefaultMessage string
	// DefaultEnMessage holds the default value on creation for the "en_message" field.
	DefaultEnMessage string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
)