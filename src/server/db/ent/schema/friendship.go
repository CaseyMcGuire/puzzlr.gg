package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

// Friendship holds the schema definition for the bidirectional friendship relationship.
type Friendship struct {
	ent.Schema
}

// Fields of the Friendship.
func (Friendship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Immutable(),
		field.Int("friend_id").
			Immutable(),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Immutable().
			Required().
			Field("user_id"),
		edge.To("friend", User.Type).
			Unique().
			Immutable().
			Required().
			Field("friend_id"),
	}
}

func (Friendship) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "friend_id").
			Unique(),
	}
}

func (Friendship) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			RejectDirectFriendshipMutation,
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne,
		),
	}
}

func (Friendship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Check: "user_id <> friend_id",
		},
		entgql.Skip(),
	}
}
