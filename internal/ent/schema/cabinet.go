package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/GalvinGao/otogenavi-backend/internal/x/entid"
)

// Cabinet holds the schema definition for the Cabinet entity.
type Cabinet struct {
	ent.Schema
}

// Fields of the Cabinet.
func (Cabinet) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("cabinet_id").
			Unique().
			Immutable().
			DefaultFunc(entid.Cabinet),
		field.Int("count").
			Optional(),
		field.String("location_id").
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("game_id").
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
	}
}

// Indexes of the Cabinet.
func (Cabinet) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("game", "location").Unique(),
	}
}

// Edges of the Cabinet.
func (Cabinet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("location", Location.Type).
			Field("location_id").
			Unique().
			Required(),
		edge.To("game", Game.Type).
			Field("game_id").
			Unique().
			Required(),
	}
}
