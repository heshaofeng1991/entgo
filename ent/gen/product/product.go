// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSku holds the string denoting the sku field in the database.
	FieldSku = "sku"
	// FieldBarcode holds the string denoting the barcode field in the database.
	FieldBarcode = "barcode"
	// FieldCustomerCode holds the string denoting the customer_code field in the database.
	FieldCustomerCode = "customer_code"
	// FieldDeclaredName holds the string denoting the declared_name field in the database.
	FieldDeclaredName = "declared_name"
	// FieldDeclaredCnName holds the string denoting the declared_cn_name field in the database.
	FieldDeclaredCnName = "declared_cn_name"
	// FieldDeclaredValueInUsd holds the string denoting the declared_value_in_usd field in the database.
	FieldDeclaredValueInUsd = "declared_value_in_usd"
	// FieldDeclaredValueInEur holds the string denoting the declared_value_in_eur field in the database.
	FieldDeclaredValueInEur = "declared_value_in_eur"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldHsCode holds the string denoting the hs_code field in the database.
	FieldHsCode = "hs_code"
	// FieldMaterial holds the string denoting the material field in the database.
	FieldMaterial = "material"
	// FieldPurpose holds the string denoting the purpose field in the database.
	FieldPurpose = "purpose"
	// FieldWithBarcode holds the string denoting the with_barcode field in the database.
	FieldWithBarcode = "with_barcode"
	// FieldBarcodeService holds the string denoting the barcode_service field in the database.
	FieldBarcodeService = "barcode_service"
	// FieldBarcodeTemplate holds the string denoting the barcode_template field in the database.
	FieldBarcodeTemplate = "barcode_template"
	// FieldImages holds the string denoting the images field in the database.
	FieldImages = "images"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// FieldConfirmedAttributes holds the string denoting the confirmed_attributes field in the database.
	FieldConfirmedAttributes = "confirmed_attributes"
	// FieldGrams holds the string denoting the grams field in the database.
	FieldGrams = "grams"
	// FieldInboundGrams holds the string denoting the inbound_grams field in the database.
	FieldInboundGrams = "inbound_grams"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldMaxAgvQty holds the string denoting the max_agv_qty field in the database.
	FieldMaxAgvQty = "max_agv_qty"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// EdgeTenant holds the string denoting the tenant edge name in mutations.
	EdgeTenant = "tenant"
	// EdgeInventories holds the string denoting the inventories edge name in mutations.
	EdgeInventories = "inventories"
	// EdgeProductMappings holds the string denoting the product_mappings edge name in mutations.
	EdgeProductMappings = "product_mappings"
	// Table holds the table name of the product in the database.
	Table = "products"
	// TenantTable is the table that holds the tenant relation/edge.
	TenantTable = "products"
	// TenantInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantInverseTable = "tenants"
	// TenantColumn is the table column denoting the tenant relation/edge.
	TenantColumn = "product_tenant"
	// InventoriesTable is the table that holds the inventories relation/edge.
	InventoriesTable = "inventories"
	// InventoriesInverseTable is the table name for the Inventory entity.
	// It exists in this package in order to avoid circular dependency with the "inventory" package.
	InventoriesInverseTable = "inventories"
	// InventoriesColumn is the table column denoting the inventories relation/edge.
	InventoriesColumn = "product_id"
	// ProductMappingsTable is the table that holds the product_mappings relation/edge.
	ProductMappingsTable = "product_mappings"
	// ProductMappingsInverseTable is the table name for the ProductMapping entity.
	// It exists in this package in order to avoid circular dependency with the "productmapping" package.
	ProductMappingsInverseTable = "product_mappings"
	// ProductMappingsColumn is the table column denoting the product_mappings relation/edge.
	ProductMappingsColumn = "product_id"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldSku,
	FieldBarcode,
	FieldCustomerCode,
	FieldDeclaredName,
	FieldDeclaredCnName,
	FieldDeclaredValueInUsd,
	FieldDeclaredValueInEur,
	FieldCurrency,
	FieldHsCode,
	FieldMaterial,
	FieldPurpose,
	FieldWithBarcode,
	FieldBarcodeService,
	FieldBarcodeTemplate,
	FieldImages,
	FieldAttributes,
	FieldConfirmedAttributes,
	FieldGrams,
	FieldInboundGrams,
	FieldLength,
	FieldWidth,
	FieldHeight,
	FieldMaxAgvQty,
	FieldStatus,
	FieldCreatedBy,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_tenant",
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultDeclaredValueInUsd holds the default value on creation for the "declared_value_in_usd" field.
	DefaultDeclaredValueInUsd float64
	// DefaultDeclaredValueInEur holds the default value on creation for the "declared_value_in_eur" field.
	DefaultDeclaredValueInEur float64
	// DefaultCurrency holds the default value on creation for the "currency" field.
	DefaultCurrency string
	// DefaultHsCode holds the default value on creation for the "hs_code" field.
	DefaultHsCode string
	// DefaultMaterial holds the default value on creation for the "material" field.
	DefaultMaterial string
	// DefaultPurpose holds the default value on creation for the "purpose" field.
	DefaultPurpose string
	// DefaultWithBarcode holds the default value on creation for the "with_barcode" field.
	DefaultWithBarcode int8
	// DefaultBarcodeService holds the default value on creation for the "barcode_service" field.
	DefaultBarcodeService int8
	// DefaultGrams holds the default value on creation for the "grams" field.
	DefaultGrams int
	// DefaultInboundGrams holds the default value on creation for the "inbound_grams" field.
	DefaultInboundGrams int
	// DefaultLength holds the default value on creation for the "length" field.
	DefaultLength int
	// DefaultWidth holds the default value on creation for the "width" field.
	DefaultWidth int
	// DefaultHeight holds the default value on creation for the "height" field.
	DefaultHeight int
	// DefaultMaxAgvQty holds the default value on creation for the "max_agv_qty" field.
	DefaultMaxAgvQty int
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
)
