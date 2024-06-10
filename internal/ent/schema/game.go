package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/GalvinGao/otogenavi-backend/internal/x/entid"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("game_id").
			Unique().
			Immutable().
			DefaultFunc(entid.Game),
		field.String("name").
			Unique(),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return nil
}
