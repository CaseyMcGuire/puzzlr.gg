package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/friendrequest"
	"puzzlr.gg/src/server/db/ent/codegen/friendship"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

func ValidateFriendRequestCreate(next ent.Mutator) ent.Mutator {
	return hook.FriendRequestFunc(func(ctx context.Context, m *codegen.FriendRequestMutation) (ent.Value, error) {
		requesterID, ok := m.RequesterID()
		if !ok {
			return nil, fmt.Errorf("missing requester_id on create")
		}

		recipientID, ok := m.RecipientID()
		if !ok {
			return nil, fmt.Errorf("missing recipient_id on create")
		}

		client := m.Client()

		alreadyFriends, err := client.Friendship.Query().
			Where(
				friendship.Or(
					friendship.And(
						friendship.UserIDEQ(requesterID),
						friendship.FriendIDEQ(recipientID),
					),
					friendship.And(
						friendship.UserIDEQ(recipientID),
						friendship.FriendIDEQ(requesterID),
					),
				),
			).
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if alreadyFriends {
			return nil, fmt.Errorf("cannot request friendship with someone who is already your friend")
		}

		reversePending, err := client.FriendRequest.Query().
			Where(
				friendrequest.RequesterIDEQ(recipientID),
				friendrequest.RecipientIDEQ(requesterID),
			).
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if reversePending {
			return nil, fmt.Errorf("cannot send a friend request to someone who already has a pending request for you")
		}

		return next.Mutate(ctx, m)
	})
}
