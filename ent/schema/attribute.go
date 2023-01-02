package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	
	"time"
)

// Attribute holds the schema definition for the Attribute entity.
type Attribute struct {
	ent.Schema
}

// Fields of the Attribute.
func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "int unsigned", // Override MySQL.
		}).Unique(),
		
		field.Int8("type").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("1=>Product Attribute, 2=>Product Confirm Attribute, 3=>Channel Exclusion Attribute"),
		
		field.String("value").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}),
		
		field.String("description").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Default(""),
		
		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(4)", // Override MySQL.
		}).Default(1).Comment("0=>Inactive, 1=>Active"),
		
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Default(time.Now),
		
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now),
		
		field.Time("deleted_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),
	}
}

// Edges of the Attribute.
func (Attribute) Edges() []ent.Edge {
	return nil
}
