package schema

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/friendrequest"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

var ErrFriendshipAcceptanceRequiresPendingIncomingRequest = errors.New("cannot create friendship without a pending incoming friend request")

func ValidateFriendshipAcceptance(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *codegen.UserMutation) (ent.Value, error) {
		friendIDs := m.FriendsIDs()
		if len(friendIDs) == 0 {
			return next.Mutate(ctx, m)
		}

		if !m.Op().Is(ent.OpUpdateOne) {
			return nil, fmt.Errorf("friends can only be added by accepting a pending friend request on a single user")
		}

		userID, ok := m.ID()
		if !ok {
			return nil, fmt.Errorf("missing user id for friendship acceptance")
		}

		client := m.Client()
		for _, friendID := range friendIDs {
			if friendID == userID {
				continue
			}

			pendingRequest, err := client.FriendRequest.Query().
				Where(
					friendrequest.RequesterIDEQ(friendID),
					friendrequest.RecipientIDEQ(userID),
				).
				Exist(ctx)
			if err != nil {
				return nil, err
			}
			if !pendingRequest {
				return nil, ErrFriendshipAcceptanceRequiresPendingIncomingRequest
			}
		}

		return next.Mutate(ctx, m)
	})
}
