// Code generated by ent, DO NOT EDIT.

package inventory

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the inventory type in the database.
	Label = "inventory"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldWarehouseID holds the string denoting the warehouse_id field in the database.
	FieldWarehouseID = "warehouse_id"
	// FieldStorageQty holds the string denoting the storage_qty field in the database.
	FieldStorageQty = "storage_qty"
	// FieldAvailableQty holds the string denoting the available_qty field in the database.
	FieldAvailableQty = "available_qty"
	// FieldPrepareShipQty holds the string denoting the prepare_ship_qty field in the database.
	FieldPrepareShipQty = "prepare_ship_qty"
	// FieldPrepareShelveQty holds the string denoting the prepare_shelve_qty field in the database.
	FieldPrepareShelveQty = "prepare_shelve_qty"
	// FieldQuicktronStorageQty holds the string denoting the quicktron_storage_qty field in the database.
	FieldQuicktronStorageQty = "quicktron_storage_qty"
	// FieldQuicktronAvailableQty holds the string denoting the quicktron_available_qty field in the database.
	FieldQuicktronAvailableQty = "quicktron_available_qty"
	// FieldQuicktronPrepareOutboundQty holds the string denoting the quicktron_prepare_outbound_qty field in the database.
	FieldQuicktronPrepareOutboundQty = "quicktron_prepare_outbound_qty"
	// FieldQuicktronPrepareShelveQty holds the string denoting the quicktron_prepare_shelve_qty field in the database.
	FieldQuicktronPrepareShelveQty = "quicktron_prepare_shelve_qty"
	// FieldNormalStorageQty holds the string denoting the normal_storage_qty field in the database.
	FieldNormalStorageQty = "normal_storage_qty"
	// FieldNormalAvailableQty holds the string denoting the normal_available_qty field in the database.
	FieldNormalAvailableQty = "normal_available_qty"
	// FieldNormalPrepareOutboundQty holds the string denoting the normal_prepare_outbound_qty field in the database.
	FieldNormalPrepareOutboundQty = "normal_prepare_outbound_qty"
	// FieldNormalPrepareShelveQty holds the string denoting the normal_prepare_shelve_qty field in the database.
	FieldNormalPrepareShelveQty = "normal_prepare_shelve_qty"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the inventory in the database.
	Table = "inventories"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "inventories"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "inventory_tenant"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "inventories"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "product_id"
)

// Columns holds all SQL columns for inventory fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldProductID,
	FieldWarehouseID,
	FieldStorageQty,
	FieldAvailableQty,
	FieldPrepareShipQty,
	FieldPrepareShelveQty,
	FieldQuicktronStorageQty,
	FieldQuicktronAvailableQty,
	FieldQuicktronPrepareOutboundQty,
	FieldQuicktronPrepareShelveQty,
	FieldNormalStorageQty,
	FieldNormalAvailableQty,
	FieldNormalPrepareOutboundQty,
	FieldNormalPrepareShelveQty,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "inventories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"inventory_tenant",
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
	// DefaultStorageQty holds the default value on creation for the "storage_qty" field.
	DefaultStorageQty int32
	// DefaultAvailableQty holds the default value on creation for the "available_qty" field.
	DefaultAvailableQty int32
	// DefaultPrepareShipQty holds the default value on creation for the "prepare_ship_qty" field.
	DefaultPrepareShipQty int32
	// DefaultPrepareShelveQty holds the default value on creation for the "prepare_shelve_qty" field.
	DefaultPrepareShelveQty int32
	// DefaultQuicktronStorageQty holds the default value on creation for the "quicktron_storage_qty" field.
	DefaultQuicktronStorageQty int32
	// DefaultQuicktronAvailableQty holds the default value on creation for the "quicktron_available_qty" field.
	DefaultQuicktronAvailableQty int32
	// DefaultQuicktronPrepareOutboundQty holds the default value on creation for the "quicktron_prepare_outbound_qty" field.
	DefaultQuicktronPrepareOutboundQty int32
	// DefaultQuicktronPrepareShelveQty holds the default value on creation for the "quicktron_prepare_shelve_qty" field.
	DefaultQuicktronPrepareShelveQty int32
	// DefaultNormalStorageQty holds the default value on creation for the "normal_storage_qty" field.
	DefaultNormalStorageQty int32
	// DefaultNormalAvailableQty holds the default value on creation for the "normal_available_qty" field.
	DefaultNormalAvailableQty int32
	// DefaultNormalPrepareOutboundQty holds the default value on creation for the "normal_prepare_outbound_qty" field.
	DefaultNormalPrepareOutboundQty int32
	// DefaultNormalPrepareShelveQty holds the default value on creation for the "normal_prepare_shelve_qty" field.
	DefaultNormalPrepareShelveQty int32
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
)