package schema

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent"
	entprivacy "entgo.io/ent/privacy"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/reqctx"
)

var ErrOnlyRequesterCanCreateFriendRequest = errors.New("only the requester can create a friend request")

func (FriendRequest) Policy() ent.Policy {
	return entprivacy.Policy{
		Mutation: entprivacy.MutationPolicy{
			entprivacy.OnMutationOperation(
				entprivacy.MutationRuleFunc(authorizeFriendRequestCreate),
				ent.OpCreate,
			),
		},
	}
}

func authorizeFriendRequestCreate(ctx context.Context, m ent.Mutation) error {
	mutation, ok := m.(*codegen.FriendRequestMutation)
	if !ok {
		return entprivacy.Denyf("unexpected mutation type %T", m)
	}

	requesterID, ok := mutation.RequesterID()
	if !ok {
		return fmt.Errorf("missing requester_id on create")
	}

	actorID, err := reqctx.UserIDFromContext(ctx)
	if err != nil {
		return err
	}
	if actorID != requesterID {
		return ErrOnlyRequesterCanCreateFriendRequest
	}

	return nil
}
