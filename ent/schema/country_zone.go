package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CountryZone holds the schema definition for the Country Zones entity.
type CountryZone struct {
	ent.Schema
}

// Mixin of the CountryZone.
func (CountryZone) Mixin() []ent.Mixin {
	return nil
}

// Fields of the CountryZone.
func (CountryZone) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("channel_id").SchemaType(map[string]string{
			dialect.MySQL: "int(10) unsigned", // Override MySQL.
		}).Default(0),

		field.String("country_code").SchemaType(map[string]string{
			dialect.MySQL: "char(2)", // Override MySQL.
		}).NotEmpty(),

		field.String("zip_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),

		field.String("start_zip_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),

		field.String("end_zip_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Default(""),

		field.String("zone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).NotEmpty(),
	}
}

// Edges of the CountryZone.
func (CountryZone) Edges() []ent.Edge {
	return nil
}
