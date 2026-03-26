package schema

import (
	"context"
	"errors"

	"entgo.io/ent"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

var ErrDirectFriendshipMutationForbidden = errors.New("direct friendship mutation is forbidden; accept a pending friend request via the user friends edge")

func RejectDirectFriendshipMutation(next ent.Mutator) ent.Mutator {
	return hook.FriendshipFunc(func(ctx context.Context, m *codegen.FriendshipMutation) (ent.Value, error) {
		return nil, ErrDirectFriendshipMutationForbidden
	})
}
