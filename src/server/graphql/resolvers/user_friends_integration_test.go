//go:build integration

package resolvers_test

import (
	"context"
	"testing"

	ent "puzzlr.gg/src/server/db/ent/codegen"
)

func TestUserQueryFriendsReturnsBidirectionalFriendship(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	if err := integrationClient.User.
		UpdateOneID(alice.ID).
		AddFriendIDs(bob.ID).
		Exec(ctx); err != nil {
		t.Fatalf("adding friendship failed: %v", err)
	}

	aliceFriends, err := alice.QueryFriends().All(ctx)
	if err != nil {
		t.Fatalf("querying alice friends failed: %v", err)
	}
	assertOnlyFriend(t, aliceFriends, bob.ID)

	bobFriends, err := bob.QueryFriends().All(ctx)
	if err != nil {
		t.Fatalf("querying bob friends failed: %v", err)
	}
	assertOnlyFriend(t, bobFriends, alice.ID)

	friendshipCount, err := integrationClient.Friendship.Query().Count(ctx)
	if err != nil {
		t.Fatalf("counting friendship rows failed: %v", err)
	}
	if friendshipCount != 2 {
		t.Fatalf("expected 2 mirrored friendship rows, got %d", friendshipCount)
	}
}

func TestUserQueryFriendsRejectsSelfFriendship(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)

	err := integrationClient.User.
		UpdateOneID(alice.ID).
		AddFriendIDs(alice.ID).
		Exec(ctx)
	if err == nil {
		t.Fatal("expected self-friendship to fail, got nil")
	}
	if !ent.IsConstraintError(err) {
		t.Fatalf("expected constraint error, got: %v", err)
	}
}

func assertOnlyFriend(t *testing.T, got []*ent.User, wantID int) {
	t.Helper()

	if len(got) != 1 {
		t.Fatalf("expected 1 friend, got %d", len(got))
	}
	if got[0].ID != wantID {
		t.Fatalf("expected friend %d, got %d", wantID, got[0].ID)
	}
}
