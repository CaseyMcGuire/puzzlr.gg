//go:build integration

package resolvers_test

import (
	"context"
	"testing"
)

func TestUserQueryReturnsUserByID(t *testing.T) {
	ctx := context.Background()
	user := mustCreateUser(t, ctx)

	got, err := newTestResolver().Query().User(ctx, user.ID)
	if err != nil {
		t.Fatalf("querying user failed: %v", err)
	}
	if got == nil {
		t.Fatal("expected user, got nil")
	}
	if got.ID != user.ID {
		t.Fatalf("expected user ID %d, got %d", user.ID, got.ID)
	}
	if got.Email != user.Email {
		t.Fatalf("expected email %q, got %q", user.Email, got.Email)
	}
}

func TestUserQueryReturnsNilForMissingUser(t *testing.T) {
	ctx := context.Background()

	got, err := newTestResolver().Query().User(ctx, 999999)
	if err != nil {
		t.Fatalf("expected missing user lookup to return nil without error, got: %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil user for missing ID, got %#v", got)
	}
}
