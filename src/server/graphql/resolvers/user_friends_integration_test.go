//go:build integration

package resolvers_test

import (
	"context"
	"strings"
	"testing"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/friendship"
	"puzzlr.gg/src/server/reqctx"
)

func TestUserQueryFriendsReturnsBidirectionalFriendship(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	if _, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID)); err != nil {
		t.Fatalf("creating friend request failed: %v", err)
	}

	if err := integrationClient.User.
		UpdateOneID(bob.ID).
		AddFriendIDs(alice.ID).
		Exec(reqctx.WithUserID(ctx, bob.ID)); err != nil {
		t.Fatalf("accepting friendship failed: %v", err)
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

	friendshipCount, err := integrationClient.Friendship.Query().
		Where(
			friendship.Or(
				friendship.And(
					friendship.UserIDEQ(alice.ID),
					friendship.FriendIDEQ(bob.ID),
				),
				friendship.And(
					friendship.UserIDEQ(bob.ID),
					friendship.FriendIDEQ(alice.ID),
				),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatalf("counting friendship rows failed: %v", err)
	}
	if friendshipCount != 2 {
		t.Fatalf("expected 2 mirrored friendship rows for the accepted pair, got %d", friendshipCount)
	}
}

func TestUserQueryFriendsRejectsSelfFriendship(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)

	err := integrationClient.User.
		UpdateOneID(alice.ID).
		AddFriendIDs(alice.ID).
		Exec(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected self-friendship to fail, got nil")
	}
	if !ent.IsConstraintError(err) {
		t.Fatalf("expected constraint error, got: %v", err)
	}
}

func TestFriendshipAcceptanceRequiresPendingIncomingRequest(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	err := integrationClient.User.
		UpdateOneID(bob.ID).
		AddFriendIDs(alice.ID).
		Exec(reqctx.WithUserID(ctx, bob.ID))
	if err == nil {
		t.Fatal("expected accepting without a request to fail, got nil")
	}
	if err.Error() != "cannot create friendship without a pending incoming friend request" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFriendshipAcceptanceRequiresRecipientActor(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	if _, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID)); err != nil {
		t.Fatalf("creating friend request failed: %v", err)
	}

	err := integrationClient.User.
		UpdateOneID(bob.ID).
		AddFriendIDs(alice.ID).
		Exec(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected non-recipient acceptance to fail, got nil")
	}
	if err.Error() != "only the user can mutate their own record" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDirectFriendshipMutationIsRejected(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	_, err := integrationClient.Friendship.
		Create().
		SetUserID(bob.ID).
		SetFriendID(alice.ID).
		Save(ctx)
	if err == nil {
		t.Fatal("expected direct friendship mutation to fail, got nil")
	}
	if !strings.Contains(err.Error(), "direct friendship mutation is forbidden") {
		t.Fatalf("unexpected error: %v", err)
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
