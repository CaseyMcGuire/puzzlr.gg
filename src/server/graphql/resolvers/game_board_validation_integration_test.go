//go:build integration

package resolvers_test

import (
	"context"
	"errors"
	"testing"

	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/schema"
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
	if !errors.Is(err, schema.ErrTicTacToeBoardRowMustHaveThreeColumns) {
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
	if !errors.Is(err, schema.ErrTicTacToeBoardMustHaveThreeRows) {
		t.Fatalf("unexpected error: %v", err)
	}
}
