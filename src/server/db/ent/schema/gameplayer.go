package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// GamePlayer holds the schema definition for the GamePlayer entity.
type GamePlayer struct {
	ent.Schema
}

// Fields of the GamePlayer.
func (GamePlayer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("game_id"),
	}
}

// Edges of the GamePlayer.
func (GamePlayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("game", Game.Type).
			Unique().
			Required().
			Field("game_id"),
	}
}

func (GamePlayer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "game_id").
			Unique(),
	}
}

func (GamePlayer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(),
	}
}
