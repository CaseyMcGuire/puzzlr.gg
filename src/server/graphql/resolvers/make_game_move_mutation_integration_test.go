//go:build integration

package resolvers_test

import (
	"context"
	"errors"
	"testing"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/graphql/models"
	"puzzlr.gg/src/server/graphql/resolvers"
	"puzzlr.gg/src/server/reqctx"
	"puzzlr.gg/src/server/services"
)

func TestMakeGameMoveResolverSuccess(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	actor := mustCreateUser(t, ctx)
	opponent := mustCreateUser(t, ctx)
	createdGame := mustCreateTicTacToeGame(t, ctx, resolver, actor.ID, opponent.ID)

	updatedGame, err := resolver.Mutation().MakeGameMove(
		reqctx.WithUserID(ctx, actor.ID),
		makeMoveInput(createdGame.ID, 0, 0),
	)
	if err != nil {
		t.Fatalf("makeGameMove returned an error: %v", err)
	}
	if updatedGame == nil {
		t.Fatal("makeGameMove returned nil game")
	}

	if got := updatedGame.Board[0][0]; got != services.TictactoeX {
		t.Fatalf("expected X at [0][0], got %q", got)
	}
	if updatedGame.Status != game.StatusIN_PROGRESS {
		t.Fatalf("expected status IN_PROGRESS, got %s", updatedGame.Status)
	}

	currentTurnUser, err := updatedGame.QueryCurrentTurn().Only(ctx)
	if err != nil {
		t.Fatalf("querying current turn failed: %v", err)
	}
	if currentTurnUser.ID != opponent.ID {
		t.Fatalf("expected current turn to be opponent (%d), got %d", opponent.ID, currentTurnUser.ID)
	}
}

func TestMakeGameMoveResolverRequiresUserInContext(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	actor := mustCreateUser(t, ctx)
	opponent := mustCreateUser(t, ctx)
	createdGame := mustCreateTicTacToeGame(t, ctx, resolver, actor.ID, opponent.ID)

	_, err := resolver.Mutation().MakeGameMove(
		ctx,
		makeMoveInput(createdGame.ID, 0, 0),
	)
	if err == nil {
		t.Fatal("expected missing user ID error, got nil")
	}
	if !errors.Is(err, reqctx.ErrNoUserIDInContext) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMakeGameMoveResolverRejectsOutOfTurnMove(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	actor := mustCreateUser(t, ctx)
	opponent := mustCreateUser(t, ctx)
	createdGame := mustCreateTicTacToeGame(t, ctx, resolver, actor.ID, opponent.ID)

	_, err := resolver.Mutation().MakeGameMove(
		reqctx.WithUserID(ctx, opponent.ID),
		makeMoveInput(createdGame.ID, 0, 0),
	)
	if err == nil {
		t.Fatal("expected out-of-turn move error, got nil")
	}
	if !errors.Is(err, services.ErrNotYourTurn) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMakeGameMoveResolverRejectsTakenCell(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	actor := mustCreateUser(t, ctx)
	opponent := mustCreateUser(t, ctx)
	createdGame := mustCreateTicTacToeGame(t, ctx, resolver, actor.ID, opponent.ID)

	_, err := resolver.Mutation().MakeGameMove(
		reqctx.WithUserID(ctx, actor.ID),
		makeMoveInput(createdGame.ID, 0, 0),
	)
	if err != nil {
		t.Fatalf("first move returned an error: %v", err)
	}

	_, err = resolver.Mutation().MakeGameMove(
		reqctx.WithUserID(ctx, opponent.ID),
		makeMoveInput(createdGame.ID, 0, 0),
	)
	if err == nil {
		t.Fatal("expected taken-cell error, got nil")
	}
	if !errors.Is(err, services.ErrCellAlreadyTaken) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMakeGameMoveResolverSetsWinnerAndEndsGame(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	actor := mustCreateUser(t, ctx)
	opponent := mustCreateUser(t, ctx)
	createdGame := mustCreateTicTacToeGame(t, ctx, resolver, actor.ID, opponent.ID)

	moves := []struct {
		userID int
		row    int
		col    int
	}{
		{userID: actor.ID, row: 0, col: 0},
		{userID: opponent.ID, row: 1, col: 0},
		{userID: actor.ID, row: 0, col: 1},
		{userID: opponent.ID, row: 1, col: 1},
		{userID: actor.ID, row: 0, col: 2}, // winning move
	}

	var latestGame *ent.Game
	for _, move := range moves {
		updatedGame, err := resolver.Mutation().MakeGameMove(
			reqctx.WithUserID(ctx, move.userID),
			makeMoveInput(createdGame.ID, move.row, move.col),
		)
		if err != nil {
			t.Fatalf("move (%d,%d) by user %d failed: %v", move.row, move.col, move.userID, err)
		}
		latestGame = updatedGame
	}

	if latestGame == nil {
		t.Fatal("expected final game state, got nil")
	}
	if latestGame.Status != game.StatusWON {
		t.Fatalf("expected status WON, got %s", latestGame.Status)
	}

	winner, err := latestGame.QueryWinner().Only(ctx)
	if err != nil {
		t.Fatalf("querying winner failed: %v", err)
	}
	if winner.ID != actor.ID {
		t.Fatalf("expected winner %d, got %d", actor.ID, winner.ID)
	}

	_, err = latestGame.QueryCurrentTurn().Only(ctx)
	if err == nil {
		t.Fatal("expected no current turn after game end, got nil error")
	}
	if !ent.IsNotFound(err) {
		t.Fatalf("expected not found for current turn, got: %v", err)
	}
}

func makeMoveInput(gameID, row, col int) models.MakeGameMoveInput {
	return models.MakeGameMoveInput{
		GameID: gameID,
		Move: &models.GameMoveInput{
			TicTacToeMove: &models.TicTacToeMoveInput{
				Row: row,
				Col: col,
			},
		},
	}
}

func mustCreateTicTacToeGame(t *testing.T, ctx context.Context, resolver *resolvers.Resolver, actorID, opponentID int) *ent.Game {
	t.Helper()

	createdGame, err := resolver.Mutation().CreateGame(
		reqctx.WithUserID(ctx, actorID),
		&models.CreateGameInput{
			TicTacToeInput: &models.CreateTicTacToeInput{
				OpponentID: opponentID,
			},
		},
	)
	if err != nil {
		t.Fatalf("createGame returned an error: %v", err)
	}
	if createdGame == nil {
		t.Fatal("createGame returned nil game")
	}
	return createdGame
}
