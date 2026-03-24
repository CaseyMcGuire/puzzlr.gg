package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

func RejectDirectFriendshipMutation(next ent.Mutator) ent.Mutator {
	return hook.FriendshipFunc(func(ctx context.Context, m *codegen.FriendshipMutation) (ent.Value, error) {
		return nil, fmt.Errorf("direct friendship mutation is forbidden; accept a pending friend request via the user friends edge")
	})
}
