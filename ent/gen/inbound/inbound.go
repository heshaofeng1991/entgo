// Code generated by ent, DO NOT EDIT.

package inbound

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the inbound type in the database.
	Label = "inbound"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCustomerOrderID holds the string denoting the customer_order_id field in the database.
	FieldCustomerOrderID = "customer_order_id"
	// FieldCustomerCode holds the string denoting the customer_code field in the database.
	FieldCustomerCode = "customer_code"
	// FieldTrackingNumber holds the string denoting the tracking_number field in the database.
	FieldTrackingNumber = "tracking_number"
	// FieldWarehouseID holds the string denoting the warehouse_id field in the database.
	FieldWarehouseID = "warehouse_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldEstimatedArrivalAt holds the string denoting the estimated_arrival_at field in the database.
	FieldEstimatedArrivalAt = "estimated_arrival"
	// FieldShippedAt holds the string denoting the shipped_at field in the database.
	FieldShippedAt = "ship_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsPickup holds the string denoting the is_pickup field in the database.
	FieldIsPickup = "is_pickup"
	// FieldShippingMarkURL holds the string denoting the shipping_mark_url field in the database.
	FieldShippingMarkURL = "shipping_mark_url"
	// FieldPickupOrderID holds the string denoting the pickup_order_id field in the database.
	FieldPickupOrderID = "pickup_order_id"
	// FieldCarrierName holds the string denoting the carrier_name field in the database.
	FieldCarrierName = "carrier_name"
	// FieldOrderNumber holds the string denoting the order_number field in the database.
	FieldOrderNumber = "inbound_order_number"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeInboundItems holds the string denoting the inbound_items edge name in mutations.
	EdgeInboundItems = "inbound_items"
	// Table holds the table name of the inbound in the database.
	Table = "inbounds"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "inbounds"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "inbound_tenant"
	// InboundItemsTable is the table that holds the inbound_items relation/edge.
	InboundItemsTable = "inbound_items"
	// InboundItemsInverseTable is the table name for the InboundItem entity.
	// It exists in this package in order to avoid circular dependency with the "inbounditem" package.
	InboundItemsInverseTable = "inbound_items"
	// InboundItemsColumn is the table column denoting the inbound_items relation/edge.
	InboundItemsColumn = "inbound_id"
)

// Columns holds all SQL columns for inbound fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCustomerOrderID,
	FieldCustomerCode,
	FieldTrackingNumber,
	FieldWarehouseID,
	FieldDescription,
	FieldEstimatedArrivalAt,
	FieldShippedAt,
	FieldStatus,
	FieldType,
	FieldIsPickup,
	FieldShippingMarkURL,
	FieldPickupOrderID,
	FieldCarrierName,
	FieldOrderNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "inbounds"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"inbound_tenant",
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/heshaofeng1991/entgo/ent/gen/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCustomerOrderID holds the default value on creation for the "customer_order_id" field.
	DefaultCustomerOrderID string
	// DefaultCustomerCode holds the default value on creation for the "customer_code" field.
	DefaultCustomerCode string
	// DefaultTrackingNumber holds the default value on creation for the "tracking_number" field.
	DefaultTrackingNumber string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int8
	// DefaultIsPickup holds the default value on creation for the "is_pickup" field.
	DefaultIsPickup int8
	// DefaultCarrierName holds the default value on creation for the "carrier_name" field.
	DefaultCarrierName string
	// DefaultOrderNumber holds the default value on creation for the "order_number" field.
	DefaultOrderNumber string
)