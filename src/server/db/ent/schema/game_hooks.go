package schema

import (
	"context"
	"errors"
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

var (
	ErrTicTacToeBoardMustHaveThreeRows       = errors.New("tic tac toe board must have exactly 3 rows")
	ErrTicTacToeBoardRowMustHaveThreeColumns = errors.New("tic tac toe board row must have exactly 3 columns")
)

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

func ValidateBoardShapeForType(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		board, ok := m.Board()
		if !ok {
			return next.Mutate(ctx, m)
		}

		gameType, err := gameTypeFromMutation(ctx, m)
		if err != nil {
			return nil, err
		}

		switch gameType {
		case game.TypeTIC_TAC_TOE:
			if err := validateTicTacToeBoardShape(board); err != nil {
				return nil, err
			}
		}

		return next.Mutate(ctx, m)
	})
}

func ValidateReferencedPlayersBelongToGame(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		if !m.Op().Is(ent.OpUpdateOne) {
			return next.Mutate(ctx, m)
		}

		gameID, ok := m.ID()
		if !ok {
			return nil, fmt.Errorf("missing id for %s (only UpdateOne/DeleteOne supported)", m.Op())
		}

		tx, err := m.Tx()
		if err != nil {
			return nil, fmt.Errorf("game mutation must run in transaction: %w", err)
		}

		if playerID, ok := m.WinnerPlayerID(); ok {
			if err := validateReferencedPlayerBelongsToGame(ctx, tx, playerID, gameID, "winner_player"); err != nil {
				return nil, err
			}
		}

		if playerID, ok := m.CurrentTurnPlayerID(); ok {
			if err := validateReferencedPlayerBelongsToGame(ctx, tx, playerID, gameID, "current_turn_player"); err != nil {
				return nil, err
			}
		}

		return next.Mutate(ctx, m)
	})
}

func RejectBulkGameMutation(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		return nil, fmt.Errorf("bulk game mutation (%s) is not allowed; use UpdateOne/DeleteOne", m.Op())
	})
}

func getGamePlayerCounts(gameType game.Type) (minimum, maximum int, ok bool) {
	switch gameType {
	case game.TypeTIC_TAC_TOE:
		return 2, 2, true
	}
	return -1, -1, false
}

func validatePlayerCount(numPlayers int, gameType game.Type) error {
	minimumPlayers, maximumPlayers, ok := getGamePlayerCounts(gameType)
	if !ok {
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

	switch oldStatus {
	case game.StatusPENDING:
		return validatePlayerCountForMutation(ctx, m)
	}

	return nil
}

func gameTypeFromMutation(ctx context.Context, m *codegen.GameMutation) (game.Type, error) {
	if m.Op().Is(ent.OpCreate) {
		gameType, ok := m.GetType()
		if !ok {
			return "", fmt.Errorf("missing type for create")
		}
		return gameType, nil
	}
	if m.Op().Is(ent.OpUpdateOne) {
		return m.OldType(ctx)
	}

	return "", fmt.Errorf("unsupported op for game type lookup: %s", m.Op())
}

func validateTicTacToeBoardShape(board [][]string) error {
	if len(board) != 3 {
		return ErrTicTacToeBoardMustHaveThreeRows
	}
	for rowIndex, row := range board {
		if len(row) != 3 {
			return fmt.Errorf("%w: row %d", ErrTicTacToeBoardRowMustHaveThreeColumns, rowIndex)
		}
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

func validateReferencedPlayerBelongsToGame(ctx context.Context, tx *codegen.Tx, playerID int, gameID int, edgeName string) error {
	exists, err := tx.GamePlayer.Query().
		Where(
			gameplayer.ID(playerID),
			gameplayer.GameID(gameID),
		).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("%s must reference a player in the same game", edgeName)
	}
	return nil
}
