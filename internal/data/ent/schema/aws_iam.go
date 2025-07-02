package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Aws_iam holds the schema definition for the Aws_iam entity.
type Aws_iam struct {
	ent.Schema
}

// Fields of the Aws_iam.

func (Aws_iam) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid"),
		field.String("account_id"),
		field.String("iam_name"),
		field.String("access_key"),
		field.String("secret_key"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Aws_iam.
func (Aws_iam) Edges() []ent.Edge {
	return nil
}
