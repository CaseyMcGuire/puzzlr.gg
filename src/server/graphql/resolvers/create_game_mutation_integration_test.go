//go:build integration

package resolvers_test

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/graphql/models"
	"puzzlr.gg/src/server/reqctx"
)

var uniqueUserCounter int64

func TestCreateGameResolverSuccess(t *testing.T) {
	ctx := context.Background()

	actor := mustCreateUser(t, ctx)
	opponent := mustCreateUser(t, ctx)

	resolver := newTestResolver()
	game, err := resolver.Mutation().CreateGame(
		reqctx.WithUserID(ctx, actor.ID),
		&models.CreateGameInput{
			TicTacToeInput: &models.CreateTicTacToeInput{
				OpponentID: opponent.ID,
			},
		},
	)
	if err != nil {
		t.Fatalf("createGame returned an error: %v", err)
	}
	if game == nil {
		t.Fatal("createGame returned nil game")
	}

	if got := game.Type.String(); got != "TIC_TAC_TOE" {
		t.Fatalf("unexpected game type: %s", got)
	}
	if len(game.Board) != 3 || len(game.Board[0]) != 3 || len(game.Board[1]) != 3 || len(game.Board[2]) != 3 {
		t.Fatalf("unexpected board dimensions: %#v", game.Board)
	}

	playerCount, err := game.QueryUser().Count(ctx)
	if err != nil {
		t.Fatalf("querying players failed: %v", err)
	}
	if playerCount != 2 {
		t.Fatalf("expected 2 players, got %d", playerCount)
	}
}

func TestCreateGameResolverRequiresUserInContext(t *testing.T) {
	ctx := context.Background()
	opponent := mustCreateUser(t, ctx)

	_, err := newTestResolver().Mutation().CreateGame(
		ctx,
		&models.CreateGameInput{
			TicTacToeInput: &models.CreateTicTacToeInput{
				OpponentID: opponent.ID,
			},
		},
	)
	if err == nil {
		t.Fatal("expected missing user ID error, got nil")
	}
	if !strings.Contains(err.Error(), "no user ID in context") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func mustCreateUser(t *testing.T, ctx context.Context) *ent.User {
	t.Helper()

	if integrationClient == nil {
		t.Fatal("integration client is not initialized")
	}

	email := uniqueEmail()
	user, err := integrationClient.User.
		Create().
		SetEmail(email).
		SetHashedPassword("test-hash").
		Save(ctx)
	if err != nil {
		t.Fatalf("creating user failed: %v", err)
	}
	return user
}

func uniqueEmail() string {
	n := atomic.AddInt64(&uniqueUserCounter, 1)
	return fmt.Sprintf("resolver-integration-%d@example.com", n)
}
