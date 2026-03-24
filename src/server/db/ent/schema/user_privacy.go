package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	entprivacy "entgo.io/ent/privacy"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/reqctx"
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
		return fmt.Errorf("bulk user mutation is forbidden")
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
		return fmt.Errorf("only the user can mutate their own record")
	}

	return nil
}
