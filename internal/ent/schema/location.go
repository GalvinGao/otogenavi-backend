package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/GalvinGao/otogenavi-backend/internal/x/entid"
	"github.com/GalvinGao/otogenavi-backend/internal/x/postgis"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("location_id").
			Unique().
			Immutable().
			DefaultFunc(entid.Location),
		field.String("name").
			NotEmpty(),
		field.String("deduplication_key").
			NotEmpty().
			Unique().
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		field.String("raw_address").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.Other("coordinate", &postgis.PointS{}).
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			).
			SchemaType(map[string]string{
				dialect.Postgres: "geometry(Point,4326)",
			}),
	}
}

// Indexes of the Location.
func (Location) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coordinate"),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cabinets", Cabinet.Type).
			Ref("location"),
	}
}

// Annotations of the Location.
func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate()),
		entgql.MultiOrder(),
	}
}
