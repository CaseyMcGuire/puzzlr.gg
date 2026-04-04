package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

// GamePlayer holds the schema definition for the GamePlayer entity.
type GamePlayer struct {
	ent.Schema
}

// Fields of the GamePlayer.
func (GamePlayer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Optional().
			Nillable().
			Immutable().
			Annotations(
				entgql.Skip(),
			),
		field.Int("game_id").
			Immutable().
			Annotations(
				entgql.Skip(),
			),
		field.Enum("kind").
			Values("HUMAN", "AI").
			Immutable().
			Annotations(
				entgql.Type("GamePlayerKind"),
			),
		field.String("marker").
			Immutable().
			NotEmpty(),
	}
}

// Edges of the GamePlayer.
func (GamePlayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Immutable().
			Field("user_id"),
		edge.To("game", Game.Type).
			Unique().
			Immutable().
			Required().
			Field("game_id").
			Annotations(
				entgql.Skip(),
			),
	}
}

func (GamePlayer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "game_id").
			Unique(),
		index.Fields("game_id"),
	}
}

func (GamePlayer) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			ValidateGamePlayerIdentity,
			ent.OpCreate,
		),
		hook.On(
			RejectPlayerMutationUnlessPending,
			ent.OpCreate|ent.OpUpdateOne|ent.OpDeleteOne,
		),
		hook.On(
			RejectBulkGamePlayerMutation,
			ent.OpUpdate|ent.OpDelete,
		),
	}
}
