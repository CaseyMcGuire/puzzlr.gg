//go:build integration

package resolvers_test

import (
	"context"
	"strings"
	"testing"

	"puzzlr.gg/src/server/db/ent/codegen/game"
)

func TestGameCreateRejectsInvalidTicTacToeBoardShape(t *testing.T) {
	ctx := context.Background()

	_, err := integrationClient.Game.
		Create().
		SetType(game.TypeTIC_TAC_TOE).
		SetBoard([][]string{
			{"", "", ""},
			{"", ""},
			{"", "", ""},
		}).
		Save(ctx)
	if err == nil {
		t.Fatal("expected create to fail for invalid board shape")
	}
	if !strings.Contains(err.Error(), "tic tac toe board row") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGameCreateAllowsValidTicTacToeBoardShape(t *testing.T) {
	ctx := context.Background()

	created, err := integrationClient.Game.
		Create().
		SetType(game.TypeTIC_TAC_TOE).
		SetBoard([][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}).
		Save(ctx)
	if err != nil {
		t.Fatalf("expected valid board to be accepted, got error: %v", err)
	}
	if created == nil {
		t.Fatal("expected game to be created")
	}
}

func TestGameUpdateRejectsInvalidTicTacToeBoardShape(t *testing.T) {
	ctx := context.Background()

	created, err := integrationClient.Game.
		Create().
		SetType(game.TypeTIC_TAC_TOE).
		SetBoard([][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		}).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create baseline game: %v", err)
	}

	_, err = integrationClient.Game.
		UpdateOneID(created.ID).
		SetBoard([][]string{
			{"", "", ""},
			{"", "", ""},
		}).
		Save(ctx)
	if err == nil {
		t.Fatal("expected update to fail for invalid board shape")
	}
	if !strings.Contains(err.Error(), "tic tac toe board must have exactly 3 rows") {
		t.Fatalf("unexpected error: %v", err)
	}
}
