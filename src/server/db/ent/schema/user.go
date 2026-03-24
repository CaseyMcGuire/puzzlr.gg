package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("hashed_password").
			Sensitive().
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("games", Game.Type).
			Through("game_player", GamePlayer.Type),
		edge.To("friends", User.Type).
			Through("friendships", Friendship.Type),
		edge.From("sent_friend_requests", FriendRequest.Type).
			Ref("requester").
			Annotations(entgql.Skip()),
		edge.From("received_friend_requests", FriendRequest.Type).
			Ref("recipient").
			Annotations(entgql.Skip()),
		edge.To("won_games", Game.Type).
			Annotations(entgql.Skip()),
		edge.To("current_turn_games", Game.Type).
			Annotations(entgql.Skip()),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			ValidateFriendshipAcceptance,
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
