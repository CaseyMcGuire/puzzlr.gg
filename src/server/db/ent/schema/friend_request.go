package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

// FriendRequest holds the schema definition for a pending directed friend request.
type FriendRequest struct {
	ent.Schema
}

// Fields of the FriendRequest.
func (FriendRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("requester_id").
			Immutable(),
		field.Int("recipient_id").
			Immutable(),
	}
}

// Edges of the FriendRequest.
func (FriendRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requester", User.Type).
			Unique().
			Immutable().
			Required().
			Field("requester_id"),
		edge.To("recipient", User.Type).
			Unique().
			Immutable().
			Required().
			Field("recipient_id"),
	}
}

func (FriendRequest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("requester_id", "recipient_id").
			Unique(),
		index.Fields("recipient_id", "requester_id"),
	}
}

func (FriendRequest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (FriendRequest) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			ValidateFriendRequestCreate,
			ent.OpCreate,
		),
	}
}

func (FriendRequest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Check: "requester_id <> recipient_id",
		},
		entgql.Skip(),
	}
}
