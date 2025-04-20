package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("TIC_TAC_TOE").
			Immutable().
			Annotations(
				entgql.Type("GameType"),
			),
		field.JSON("board", [][]string{}).Annotations(
			entgql.Type("GameBoard"),
		),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("player_one", User.Type).
			Unique(),
		edge.To("player_two", User.Type).
			Unique(),
		edge.To("winner", User.Type).
			Unique(),
		edge.To("current_turn", User.Type).
			Unique(),
	}
}

func (Game) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
