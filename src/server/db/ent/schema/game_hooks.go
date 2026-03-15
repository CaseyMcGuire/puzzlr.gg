package schema

import (
	"context"
	"fmt"
	"slices"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/codegen/gameplayer"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
	"puzzlr.gg/src/server/db/ent/codegen/predicate"
)

func ValidatePlayerCountOnCreate(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		addedUserIDs := m.UserIDs()
		removedUserIDs := m.RemovedUserIDs()
		if len(addedUserIDs) > 0 || len(removedUserIDs) > 0 || m.UserCleared() {
			return nil, fmt.Errorf("direct game-user edge mutation is forbidden; create GamePlayer rows with marker instead")
		}
		return next.Mutate(ctx, m)
	})
}

func ValidatePlayerCountOnUpdate(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		addedUserIDs := m.UserIDs()
		removedUserIDs := m.RemovedUserIDs()
		if len(addedUserIDs) > 0 || len(removedUserIDs) > 0 || m.UserCleared() {
			return nil, fmt.Errorf("cannot change players in a game")
		}
		return next.Mutate(ctx, m)
	})
}

func ValidateStatusOnUpdate(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		newStatus, ok := m.Status()
		if !ok {
			return next.Mutate(ctx, m)
		}

		oldStatus, err := m.OldStatus(ctx)
		if err != nil {
			return nil, err
		}

		if oldStatus == newStatus {
			return next.Mutate(ctx, m)
		}

		if err := validateStatusTransition(ctx, m, oldStatus, newStatus); err != nil {
			return nil, err
		}

		return next.Mutate(ctx, m)
	})
}

func RejectBulkGameMutation(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		return nil, fmt.Errorf("bulk game mutation (%s) is not allowed; use UpdateOne/DeleteOne", m.Op())
	})
}

func getGamePlayerCounts(gameType game.Type) (minimum, maximum int) {
	switch gameType {
	case game.TypeTIC_TAC_TOE:
		return 2, 2
	}
	return -1, -1
}

func validatePlayerCount(numPlayers int, gameType game.Type) error {
	minimumPlayers, maximumPlayers := getGamePlayerCounts(gameType)
	if minimumPlayers == -1 || maximumPlayers == -1 {
		return fmt.Errorf("game of type %s is not supported", gameType)
	}

	if numPlayers < minimumPlayers || numPlayers > maximumPlayers {
		return fmt.Errorf("game of type %s must have %d-%d players", gameType, minimumPlayers, maximumPlayers)
	}
	return nil
}

func validateStatusTransition(ctx context.Context, m *codegen.GameMutation, oldStatus, newStatus game.Status) error {
	transitions := map[game.Status][]game.Status{
		game.StatusPENDING:     {game.StatusIN_PROGRESS},
		game.StatusIN_PROGRESS: {game.StatusWON, game.StatusDRAW},
		game.StatusWON:         {},
		game.StatusDRAW:        {},
	}

	allowedNextStates, ok := transitions[oldStatus]
	if !ok || !slices.Contains(allowedNextStates, newStatus) {
		return fmt.Errorf("cannot transition from %s to %s", oldStatus, newStatus)
	}

	if oldStatus == game.StatusPENDING && newStatus == game.StatusIN_PROGRESS {
		return validatePlayerCountForMutation(ctx, m)
	}

	return nil
}

func validatePlayerCountForMutation(ctx context.Context, m *codegen.GameMutation) error {
	tx, err := m.Tx()
	if err != nil {
		return fmt.Errorf("game mutation must run in transaction: %w", err)
	}

	id, ok := m.ID()
	if !ok {
		return fmt.Errorf("missing id for %s (only UpdateOne/DeleteOne supported)", m.Op())
	}

	gameState, err := tx.Game.Query().
		Where(
			game.ID(id),
			predicate.Game(func(s *sql.Selector) {
				s.ForUpdate()
			}),
		).
		Only(ctx)
	if err != nil {
		return err
	}

	numPlayers, err := tx.GamePlayer.Query().Where(gameplayer.GameID(id)).Count(ctx)
	if err != nil {
		return err
	}

	return validatePlayerCount(numPlayers, gameState.Type)
}
