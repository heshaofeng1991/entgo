// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/heshaofeng1991/entgo/ent/gen/order"
	"github.com/heshaofeng1991/entgo/ent/gen/orderitem"
	"github.com/heshaofeng1991/entgo/ent/gen/tenant"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int64 `json:"order_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int64 `json:"product_id,omitempty"`
	// PlatformProductID holds the value of the "platform_product_id" field.
	PlatformProductID int64 `json:"platform_product_id,omitempty"`
	// Barcode holds the value of the "barcode" field.
	Barcode string `json:"barcode,omitempty"`
	// FulfillmentService holds the value of the "fulfillment_service" field.
	FulfillmentService string `json:"fulfillment_service,omitempty"`
	// ExtOrderItemID holds the value of the "ext_order_item_id" field.
	ExtOrderItemID string `json:"ext_order_item_id,omitempty"`
	// ExtProductID holds the value of the "ext_product_id" field.
	ExtProductID string `json:"ext_product_id,omitempty"`
	// IsCustomItem holds the value of the "is_custom_item" field.
	IsCustomItem bool `json:"is_custom_item,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DeclaredCnName holds the value of the "declared_cn_name" field.
	DeclaredCnName string `json:"declared_cn_name,omitempty"`
	// HsCode holds the value of the "hs_code" field.
	HsCode string `json:"hs_code,omitempty"`
	// Material holds the value of the "material" field.
	Material string `json:"material,omitempty"`
	// Purpose holds the value of the "purpose" field.
	Purpose string `json:"purpose,omitempty"`
	// Images holds the value of the "images" field.
	Images string `json:"images,omitempty"`
	// Attributes holds the value of the "attributes" field.
	Attributes string `json:"attributes,omitempty"`
	// Grams holds the value of the "grams" field.
	Grams int `json:"grams,omitempty"`
	// Length holds the value of the "length" field.
	Length int `json:"length,omitempty"`
	// Width holds the value of the "width" field.
	Width int `json:"width,omitempty"`
	// Height holds the value of the "height" field.
	Height int `json:"height,omitempty"`
	// Qty holds the value of the "qty" field.
	Qty int `json:"qty,omitempty"`
	// UnitPrice holds the value of the "unit_price" field.
	UnitPrice float64 `json:"unit_price,omitempty"`
	// DeclaredValueInUsd holds the value of the "declared_value_in_usd" field.
	DeclaredValueInUsd float64 `json:"declared_value_in_usd,omitempty"`
	// DeclaredValueInEur holds the value of the "declared_value_in_eur" field.
	DeclaredValueInEur float64 `json:"declared_value_in_eur,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// FulfillQty holds the value of the "fulfill_qty" field.
	FulfillQty int `json:"fulfill_qty,omitempty"`
	// RequiresShipping holds the value of the "requires_shipping" field.
	RequiresShipping bool `json:"requires_shipping,omitempty"`
	// GiftCard holds the value of the "gift_card" field.
	GiftCard bool `json:"gift_card,omitempty"`
	// Taxable holds the value of the "taxable" field.
	Taxable bool `json:"taxable,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// Sku holds the value of the "sku" field.
	Sku string `json:"sku,omitempty"`
	// ListingSku holds the value of the "listing_sku" field.
	ListingSku string `json:"listing_sku,omitempty"`
	// DeclaredEnName holds the value of the "declared_en_name" field.
	DeclaredEnName string `json:"declared_en_name,omitempty"`
	// ProductName holds the value of the "product_name" field.
	ProductName string `json:"product_name,omitempty"`
	// CustomerCode holds the value of the "customer_code" field.
	CustomerCode string `json:"customer_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges             OrderItemEdges `json:"edges"`
	order_item_tenant *int64
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) TenantOrErr() (*Tenant, error) {
	if e.loadedTypes[0] {
		if e.Tenant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tenant.Label}
		}
		return e.Tenant, nil
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[1] {
		if e.Order == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldIsCustomItem, orderitem.FieldRequiresShipping, orderitem.FieldGiftCard, orderitem.FieldTaxable:
			values[i] = new(sql.NullBool)
		case orderitem.FieldUnitPrice, orderitem.FieldDeclaredValueInUsd, orderitem.FieldDeclaredValueInEur:
			values[i] = new(sql.NullFloat64)
		case orderitem.FieldID, orderitem.FieldOrderID, orderitem.FieldProductID, orderitem.FieldPlatformProductID, orderitem.FieldGrams, orderitem.FieldLength, orderitem.FieldWidth, orderitem.FieldHeight, orderitem.FieldQty, orderitem.FieldFulfillQty, orderitem.FieldStatus:
			values[i] = new(sql.NullInt64)
		case orderitem.FieldBarcode, orderitem.FieldFulfillmentService, orderitem.FieldExtOrderItemID, orderitem.FieldExtProductID, orderitem.FieldName, orderitem.FieldDeclaredCnName, orderitem.FieldHsCode, orderitem.FieldMaterial, orderitem.FieldPurpose, orderitem.FieldImages, orderitem.FieldAttributes, orderitem.FieldCurrency, orderitem.FieldSku, orderitem.FieldListingSku, orderitem.FieldDeclaredEnName, orderitem.FieldProductName, orderitem.FieldCustomerCode:
			values[i] = new(sql.NullString)
		case orderitem.FieldCreatedAt, orderitem.FieldUpdatedAt, orderitem.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case orderitem.ForeignKeys[0]: // order_item_tenant
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int64(value.Int64)
		case orderitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oi.CreatedAt = value.Time
			}
		case orderitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oi.UpdatedAt = value.Time
			}
		case orderitem.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				oi.DeletedAt = value.Time
			}
		case orderitem.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oi.OrderID = value.Int64
			}
		case orderitem.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				oi.ProductID = value.Int64
			}
		case orderitem.FieldPlatformProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field platform_product_id", values[i])
			} else if value.Valid {
				oi.PlatformProductID = value.Int64
			}
		case orderitem.FieldBarcode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field barcode", values[i])
			} else if value.Valid {
				oi.Barcode = value.String
			}
		case orderitem.FieldFulfillmentService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fulfillment_service", values[i])
			} else if value.Valid {
				oi.FulfillmentService = value.String
			}
		case orderitem.FieldExtOrderItemID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ext_order_item_id", values[i])
			} else if value.Valid {
				oi.ExtOrderItemID = value.String
			}
		case orderitem.FieldExtProductID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ext_product_id", values[i])
			} else if value.Valid {
				oi.ExtProductID = value.String
			}
		case orderitem.FieldIsCustomItem:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_custom_item", values[i])
			} else if value.Valid {
				oi.IsCustomItem = value.Bool
			}
		case orderitem.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				oi.Name = value.String
			}
		case orderitem.FieldDeclaredCnName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field declared_cn_name", values[i])
			} else if value.Valid {
				oi.DeclaredCnName = value.String
			}
		case orderitem.FieldHsCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hs_code", values[i])
			} else if value.Valid {
				oi.HsCode = value.String
			}
		case orderitem.FieldMaterial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field material", values[i])
			} else if value.Valid {
				oi.Material = value.String
			}
		case orderitem.FieldPurpose:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field purpose", values[i])
			} else if value.Valid {
				oi.Purpose = value.String
			}
		case orderitem.FieldImages:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value.Valid {
				oi.Images = value.String
			}
		case orderitem.FieldAttributes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attributes", values[i])
			} else if value.Valid {
				oi.Attributes = value.String
			}
		case orderitem.FieldGrams:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field grams", values[i])
			} else if value.Valid {
				oi.Grams = int(value.Int64)
			}
		case orderitem.FieldLength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				oi.Length = int(value.Int64)
			}
		case orderitem.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				oi.Width = int(value.Int64)
			}
		case orderitem.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				oi.Height = int(value.Int64)
			}
		case orderitem.FieldQty:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field qty", values[i])
			} else if value.Valid {
				oi.Qty = int(value.Int64)
			}
		case orderitem.FieldUnitPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_price", values[i])
			} else if value.Valid {
				oi.UnitPrice = value.Float64
			}
		case orderitem.FieldDeclaredValueInUsd:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field declared_value_in_usd", values[i])
			} else if value.Valid {
				oi.DeclaredValueInUsd = value.Float64
			}
		case orderitem.FieldDeclaredValueInEur:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field declared_value_in_eur", values[i])
			} else if value.Valid {
				oi.DeclaredValueInEur = value.Float64
			}
		case orderitem.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				oi.Currency = value.String
			}
		case orderitem.FieldFulfillQty:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fulfill_qty", values[i])
			} else if value.Valid {
				oi.FulfillQty = int(value.Int64)
			}
		case orderitem.FieldRequiresShipping:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field requires_shipping", values[i])
			} else if value.Valid {
				oi.RequiresShipping = value.Bool
			}
		case orderitem.FieldGiftCard:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field gift_card", values[i])
			} else if value.Valid {
				oi.GiftCard = value.Bool
			}
		case orderitem.FieldTaxable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field taxable", values[i])
			} else if value.Valid {
				oi.Taxable = value.Bool
			}
		case orderitem.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				oi.Status = int8(value.Int64)
			}
		case orderitem.FieldSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sku", values[i])
			} else if value.Valid {
				oi.Sku = value.String
			}
		case orderitem.FieldListingSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field listing_sku", values[i])
			} else if value.Valid {
				oi.ListingSku = value.String
			}
		case orderitem.FieldDeclaredEnName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field declared_en_name", values[i])
			} else if value.Valid {
				oi.DeclaredEnName = value.String
			}
		case orderitem.FieldProductName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_name", values[i])
			} else if value.Valid {
				oi.ProductName = value.String
			}
		case orderitem.FieldCustomerCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_code", values[i])
			} else if value.Valid {
				oi.CustomerCode = value.String
			}
		case orderitem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_item_tenant", value)
			} else if value.Valid {
				oi.order_item_tenant = new(int64)
				*oi.order_item_tenant = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryTenant queries the "tenant" edge of the OrderItem entity.
func (oi *OrderItem) QueryTenant() *TenantQuery {
	return (&OrderItemClient{config: oi.config}).QueryTenant(oi)
}

// QueryOrder queries the "order" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrder() *OrderQuery {
	return (&OrderItemClient{config: oi.config}).QueryOrder(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return (&OrderItemClient{config: oi.config}).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("gen: OrderItem is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("created_at=")
	builder.WriteString(oi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(oi.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ProductID))
	builder.WriteString(", ")
	builder.WriteString("platform_product_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.PlatformProductID))
	builder.WriteString(", ")
	builder.WriteString("barcode=")
	builder.WriteString(oi.Barcode)
	builder.WriteString(", ")
	builder.WriteString("fulfillment_service=")
	builder.WriteString(oi.FulfillmentService)
	builder.WriteString(", ")
	builder.WriteString("ext_order_item_id=")
	builder.WriteString(oi.ExtOrderItemID)
	builder.WriteString(", ")
	builder.WriteString("ext_product_id=")
	builder.WriteString(oi.ExtProductID)
	builder.WriteString(", ")
	builder.WriteString("is_custom_item=")
	builder.WriteString(fmt.Sprintf("%v", oi.IsCustomItem))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(oi.Name)
	builder.WriteString(", ")
	builder.WriteString("declared_cn_name=")
	builder.WriteString(oi.DeclaredCnName)
	builder.WriteString(", ")
	builder.WriteString("hs_code=")
	builder.WriteString(oi.HsCode)
	builder.WriteString(", ")
	builder.WriteString("material=")
	builder.WriteString(oi.Material)
	builder.WriteString(", ")
	builder.WriteString("purpose=")
	builder.WriteString(oi.Purpose)
	builder.WriteString(", ")
	builder.WriteString("images=")
	builder.WriteString(oi.Images)
	builder.WriteString(", ")
	builder.WriteString("attributes=")
	builder.WriteString(oi.Attributes)
	builder.WriteString(", ")
	builder.WriteString("grams=")
	builder.WriteString(fmt.Sprintf("%v", oi.Grams))
	builder.WriteString(", ")
	builder.WriteString("length=")
	builder.WriteString(fmt.Sprintf("%v", oi.Length))
	builder.WriteString(", ")
	builder.WriteString("width=")
	builder.WriteString(fmt.Sprintf("%v", oi.Width))
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(fmt.Sprintf("%v", oi.Height))
	builder.WriteString(", ")
	builder.WriteString("qty=")
	builder.WriteString(fmt.Sprintf("%v", oi.Qty))
	builder.WriteString(", ")
	builder.WriteString("unit_price=")
	builder.WriteString(fmt.Sprintf("%v", oi.UnitPrice))
	builder.WriteString(", ")
	builder.WriteString("declared_value_in_usd=")
	builder.WriteString(fmt.Sprintf("%v", oi.DeclaredValueInUsd))
	builder.WriteString(", ")
	builder.WriteString("declared_value_in_eur=")
	builder.WriteString(fmt.Sprintf("%v", oi.DeclaredValueInEur))
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(oi.Currency)
	builder.WriteString(", ")
	builder.WriteString("fulfill_qty=")
	builder.WriteString(fmt.Sprintf("%v", oi.FulfillQty))
	builder.WriteString(", ")
	builder.WriteString("requires_shipping=")
	builder.WriteString(fmt.Sprintf("%v", oi.RequiresShipping))
	builder.WriteString(", ")
	builder.WriteString("gift_card=")
	builder.WriteString(fmt.Sprintf("%v", oi.GiftCard))
	builder.WriteString(", ")
	builder.WriteString("taxable=")
	builder.WriteString(fmt.Sprintf("%v", oi.Taxable))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", oi.Status))
	builder.WriteString(", ")
	builder.WriteString("sku=")
	builder.WriteString(oi.Sku)
	builder.WriteString(", ")
	builder.WriteString("listing_sku=")
	builder.WriteString(oi.ListingSku)
	builder.WriteString(", ")
	builder.WriteString("declared_en_name=")
	builder.WriteString(oi.DeclaredEnName)
	builder.WriteString(", ")
	builder.WriteString("product_name=")
	builder.WriteString(oi.ProductName)
	builder.WriteString(", ")
	builder.WriteString("customer_code=")
	builder.WriteString(oi.CustomerCode)
	builder.WriteByte(')')
	return builder.String()
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem

func (oi OrderItems) config(cfg config) {
	for _i := range oi {
		oi[_i].config = cfg
	}
}
