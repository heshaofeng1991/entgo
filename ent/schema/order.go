package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		// base info
		field.String("order_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),
		field.Int8("mode").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(1),
		field.Int8("type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(1),
		field.Bool("offline").Default(false),
		field.Int64("store_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.Int64("warehouse_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.String("customer_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),
		field.Bool("prepay_tariff").Default(false),
		field.String("ioss_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),
		field.String("ioss_country_code").SchemaType(map[string]string{
			dialect.MySQL: "char(2)", // Override MySQL.
		}).Default(""),
		field.Int8("ioss_number_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(0),
		field.String("uk_vat_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),
		field.String("inbound_order_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),

		// status
		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).Default(1),
		field.Bool("sqs_status").Default(false).Optional(),
		field.Bool("hold_status").Default(false).Optional(),
		field.Int8("payment_status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(3)", // Override MySQL.
		}).Default(0).Optional(),
		field.Int8("cancel_request").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Default(0).Optional(),

		// platform info
		field.String("platform").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),
		field.String("platform_order_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("platform_order_no").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("platform_order_status_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Default("").Optional(),
		field.String("platform_status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Optional(),
		field.String("paid_status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default("").Optional(),
		field.Bool("fulfillment_pushed").Default(false).Optional(),
		field.String("fulfillment_status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.Time("platform_created_at").Default(nil).Optional().Nillable(),
		field.String("location_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.String("tags").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional().Nillable(),
		field.String("note").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.String("cancel_reason").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.Time("cancelled_at").Default(nil).Optional().Nillable(),
		field.Time("closed_at").Default(nil).Optional().Nillable(),

		// shipping_address
		field.String("contact_email").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),
		field.String("contact_phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),
		field.String("shipping_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.String("shipping_first_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.String("shipping_last_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.String("shipping_company").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.String("shipping_country").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.String("shipping_country_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(2)", // Override MySQL.
		}).Default(""),
		field.String("shipping_province").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.String("shipping_province_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).Default(""),
		field.String("shipping_city").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		field.String("shipping_address1").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Default(""),
		field.String("shipping_address2").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Default(""),
		field.String("shipping_zip_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),
		field.String("shipping_phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default(""),
		field.String("shipping_certificate_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(16)", // Override MySQL.
		}).Default("").Optional(),
		field.String("shipping_certificate_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("shipping_credentials_period").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),

		// other
		field.String("remark").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),

		field.Int8("prefer_shipped_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).Default(0).Optional().Comment("prefer 渠道类型"),
		field.Int64("prefer_channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional().Comment("prefer 渠道ID"),

		// shipping info
		field.String("channel_options").SchemaType(map[string]string{
			dialect.MySQL: "json", // Override MySQL.
		}).Optional(),
		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.Int64("channel_cost_id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.Int8("ship_type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(3) unsigned", // Override MySQL.
		}).Default(0).Optional().Comment("出货类型 1、传统库位出 2、AGV库位"),

		field.String("tracking_company").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.String("tracking_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("tracking_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)", // Override MySQL.
		}).Default("").Optional(),
		field.String("waybill_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("courier_order_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("nss_courier_order_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("nss_tracking_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("nss_tracking_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.String("shipping_label_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.String("shipping_label_path").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.String("label_data_path").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.String("package_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.Time("ship_date").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Default(nil).Nillable().Optional(),
		field.Time("request_shipping_at").Default(nil).Nillable().Optional(),
		field.Time("unshelf_requested_time").Default(nil).Nillable().Optional(),
		field.String("container_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.String("basket_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Default("").Optional(),
		field.Int("print_times").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(3) unsigned", // Override MySQL.
		}).Default(0).Optional(),

		// sorting & weight
		field.String("sorting_tracking_number").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default("").Optional(),
		field.Int32("sorting_port").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.Int("sorting_length").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.Int("sorting_width").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.Int("sorting_height").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional(),
		field.Int("sorting_weight").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional().Comment("签出计费重量 checkout_actual_weight和volume_weight的最大值"),
		field.Int("estimated_weight").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional().Comment("预估重量"),
		field.Int("checkout_actual_weight").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned",
		}).Default(0).Optional().Comment("签出实际重量"), // Override MySQL.
		field.Int("volume_weight").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned",
		}).Default(0).Optional().Comment("材积重量"), // Override MySQL.
		field.Int("courier_weight").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Default(0).Optional().Comment("渠道计费重量"),

		// price & fee
		field.Float("declared_value_in_usd").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000).Optional(),
		field.Float("declared_value_in_eur").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000).Optional(),
		field.Float("total_items_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0.000).Optional(),
		field.String("currency").SchemaType(map[string]string{
			dialect.MySQL: "char(3)", // Override MySQL.
		}).Optional(),

		field.Float("delivery_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("misc_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("fuel_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("registration_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("processing_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("package_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("handling_fee").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("vat").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,3)", // Override MySQL.
		}).Default(0),
		field.Float("estimated_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,2)", // Override MySQL.
		}).Default(0),
		field.String("estimated_currency").SchemaType(map[string]string{
			dialect.MySQL: "char(3)", // Override MySQL.
		}).Default("USD"),
		field.Float("receivables_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(15,2)", // Override MySQL.
		}).Default(0),
		field.String("receivables_currency").SchemaType(map[string]string{
			dialect.MySQL: "char(3)", // Override MySQL.
		}).Default("USD"),

		field.Bool("not_accept_platform_update").Default(false),
		field.Bool("not_accept_platform_update_item").Default(false),
		field.String("jd_pushed_status").SchemaType(map[string]string{
			dialect.MySQL: "char(50)", // Override MySQL.
		}).Default("000000"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_items", OrderItem.Type),
		edge.To("order_hold_reasons", OrderHoldReason.Type),
		edge.To("order_taxations", OrderTaxation.Type),
		edge.From("stores", Store.Type).Ref("orders").Field("store_id").Unique(),
		edge.From("warehouses", Warehouse.Type).Ref("orders").Field("warehouse_id").Unique(),
		edge.From("channels", Channel.Type).Ref("orders").Field("channel_id").Unique(),
		edge.To("track_mappings", TrackMapping.Type),
	}
}
