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
		edge.From("user", User.Type).
			Ref("games").
			Through("game_player", GamePlayer.Type),
		edge.From("winner", User.Type).
			Ref("won_games").
			Unique(),
		edge.From("current_turn", User.Type).
			Ref("current_turn_games").
			Unique(),
	}
}

func (Game) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			ValidatePlayerCountOnCreate,
			ent.OpCreate,
		),
		hook.On(
			ValidateBoardShapeForType,
			ent.OpCreate|ent.OpUpdateOne,
		),
		hook.On(
			RejectBulkGameMutation,
			ent.OpUpdate|ent.OpDelete,
		),
		hook.On(
			ValidatePlayerCountOnUpdate,
			ent.OpUpdateOne,
		),
		hook.On(
			ValidateStatusOnUpdate,
			ent.OpUpdateOne,
		),
	}
}

func (Game) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
