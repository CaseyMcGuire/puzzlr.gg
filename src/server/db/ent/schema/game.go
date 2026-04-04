package schema

import (
	"encoding/json"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
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
		field.JSON("metadata", json.RawMessage(nil)).
			Optional().
			Annotations(
				entgql.Skip(),
			),
		field.Int("winner_player_id").
			Optional().
			Nillable().
			Annotations(
				entgql.Skip(),
			),
		field.Int("current_turn_player_id").
			Optional().
			Nillable().
			Annotations(
				entgql.Skip(),
			),
		field.Enum("status").
			Values("PENDING", "IN_PROGRESS", "WON", "DRAW").
			Default("PENDING").
			Annotations(
				entgql.Type("GameStatus"),
			),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("players", GamePlayer.Type).
			Ref("game"),
		edge.To("winner_player", GamePlayer.Type).
			Field("winner_player_id").
			Unique(),
		edge.To("current_turn_player", GamePlayer.Type).
			Field("current_turn_player_id").
			Unique(),
	}
}

func (Game) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			ValidateBoardShapeForType,
			ent.OpCreate|ent.OpUpdateOne,
		),
		hook.On(
			RejectBulkGameMutation,
			ent.OpUpdate|ent.OpDelete,
		),
		hook.On(
			ValidateStatusOnUpdate,
			ent.OpUpdateOne,
		),
		hook.On(
			ValidateReferencedPlayersBelongToGame,
			ent.OpUpdateOne,
		),
	}
}

func (Game) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
