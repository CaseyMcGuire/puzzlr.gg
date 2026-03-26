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

var (
	ErrBulkUserMutationForbidden  = errors.New("bulk user mutation is forbidden")
	ErrOnlyUserCanMutateOwnRecord = errors.New("only the user can mutate their own record")
)

func (User) Policy() ent.Policy {
	return entprivacy.Policy{
		Mutation: entprivacy.MutationPolicy{
			entprivacy.OnMutationOperation(
				entprivacy.MutationRuleFunc(authorizeUserMutation),
				ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne,
			),
		},
	}
}

func authorizeUserMutation(ctx context.Context, m ent.Mutation) error {
	mutation, ok := m.(*codegen.UserMutation)
	if !ok {
		return entprivacy.Denyf("unexpected mutation type %T", m)
	}

	if mutation.Op().Is(ent.OpUpdate | ent.OpDelete) {
		return ErrBulkUserMutationForbidden
	}

	userID, ok := mutation.ID()
	if !ok {
		return fmt.Errorf("missing user id for user mutation")
	}

	actorID, err := reqctx.UserIDFromContext(ctx)
	if err != nil {
		return err
	}
	if actorID != userID {
		return ErrOnlyUserCanMutateOwnRecord
	}

	return nil
}
