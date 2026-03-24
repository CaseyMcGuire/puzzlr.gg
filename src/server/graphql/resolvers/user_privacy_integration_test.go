//go:build integration

package resolvers_test

import (
	"context"
	"strings"
	"testing"

	"puzzlr.gg/src/server/reqctx"
)

func TestUserUpdateOneAllowsSelfActor(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)

	if err := integrationClient.User.
		UpdateOneID(alice.ID).
		SetEmail(uniqueEmail()).
		Exec(reqctx.WithUserID(ctx, alice.ID)); err != nil {
		t.Fatalf("expected self update to succeed: %v", err)
	}
}

func TestUserUpdateOneRejectsOtherActor(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	err := integrationClient.User.
		UpdateOneID(bob.ID).
		SetEmail(uniqueEmail()).
		Exec(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected updating another user to fail, got nil")
	}
	if !strings.Contains(err.Error(), "only the user can mutate their own record") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUserDeleteOneRejectsOtherActor(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	err := integrationClient.User.
		DeleteOneID(bob.ID).
		Exec(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected deleting another user to fail, got nil")
	}
	if !strings.Contains(err.Error(), "only the user can mutate their own record") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUserBulkUpdateIsRejected(t *testing.T) {
	ctx := context.Background()

	mustCreateUser(t, ctx)

	err := integrationClient.User.
		Update().
		SetEmail(uniqueEmail()).
		Exec(reqctx.WithUserID(ctx, 1))
	if err == nil {
		t.Fatal("expected bulk update to fail, got nil")
	}
	if !strings.Contains(err.Error(), "bulk user mutation is forbidden") {
		t.Fatalf("unexpected error: %v", err)
	}
}
