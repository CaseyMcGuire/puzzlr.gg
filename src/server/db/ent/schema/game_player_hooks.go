package schema

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
	"puzzlr.gg/src/server/db/ent/codegen/predicate"
)

func RejectPlayerMutationUnlessPending(next ent.Mutator) ent.Mutator {
	return hook.GamePlayerFunc(func(ctx context.Context, m *ent.GamePlayerMutation) (ent.Value, error) {
		gameId, err := gameIDFromGamePlayerMutation(ctx, m)

		if err != nil {
			return nil, err
		}

		tx, err := m.Tx()
		if err != nil {
			return nil, fmt.Errorf("game player mutation must run in transaction: %w", err)
		}

		gameState, err := tx.Game.Query().
			Where(
				game.ID(gameId),
				predicate.Game(func(s *sql.Selector) {
					s.ForUpdate()
				}),
			).
			Only(ctx)
		if err != nil {
			return nil, err
		}

		if gameState.Status != game.StatusPENDING {
			return nil, fmt.Errorf("cannot mutate game player in non-pending state")
		}

		return next.Mutate(ctx, m)
	})
}

func RejectBulkGamePlayerMutation(next ent.Mutator) ent.Mutator {
	return hook.GamePlayerFunc(func(ctx context.Context, m *ent.GamePlayerMutation) (ent.Value, error) {
		return nil, fmt.Errorf("bulk GamePlayer update/delete is not allowed; use UpdateOne/DeleteOne")
	})
}

func gameIDFromGamePlayerMutation(ctx context.Context, m *ent.GamePlayerMutation) (int, error) {
	// Create: game_id is set on the mutation.
	if m.Op().Is(ent.OpCreate) {
		gameID, ok := m.GameID()
		if !ok {
			return 0, fmt.Errorf("missing game_id on create")
		}
		return gameID, nil
	}

	// UpdateOne/DeleteOne: load existing row by ID and read its game_id.
	id, ok := m.ID()
	if !ok {
		return 0, fmt.Errorf("missing id for %s (only UpdateOne/DeleteOne supported)", m.Op())
	}

	gp, err := m.Client().GamePlayer.Get(ctx, id)
	if err != nil {
		return 0, err
	}
	return gp.GameID, nil
}
