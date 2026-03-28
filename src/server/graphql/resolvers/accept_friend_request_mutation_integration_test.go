//go:build integration

package resolvers_test

import (
	"context"
	"testing"

	"puzzlr.gg/src/server/db/ent/codegen/friendrequest"
	"puzzlr.gg/src/server/db/ent/codegen/friendship"
	"puzzlr.gg/src/server/graphql/models"
	"puzzlr.gg/src/server/reqctx"
)

func TestAcceptFriendRequestSuccess(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	sender := mustCreateUser(t, ctx)
	recipient := mustCreateUser(t, ctx)

	// Create a pending friend request from sender to recipient.
	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(sender.ID).
		SetRecipientID(recipient.ID).
		Save(reqctx.WithUserID(ctx, sender.ID))
	if err != nil {
		t.Fatalf("creating friend request failed: %v", err)
	}

	result, err := resolver.Mutation().AcceptFriendRequest(
		reqctx.WithUserID(ctx, recipient.ID),
		models.AcceptFriendRequestInput{
			SenderID: sender.ID,
		},
	)
	if err != nil {
		t.Fatalf("acceptFriendRequest returned an error: %v", err)
	}

	success, ok := result.(*models.AcceptFriendRequestSuccess)
	if !ok {
		t.Fatalf("expected success result, got %T", result)
	}
	if success.Friend.ID != sender.ID {
		t.Fatalf("expected friend ID %d, got %d", sender.ID, success.Friend.ID)
	}

	// Verify friendship exists bidirectionally.
	friendshipCount, err := integrationClient.Friendship.Query().
		Where(
			friendship.Or(
				friendship.And(
					friendship.UserIDEQ(sender.ID),
					friendship.FriendIDEQ(recipient.ID),
				),
				friendship.And(
					friendship.UserIDEQ(recipient.ID),
					friendship.FriendIDEQ(sender.ID),
				),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatalf("counting friendships failed: %v", err)
	}
	if friendshipCount != 2 {
		t.Fatalf("expected 2 mirrored friendship rows, got %d", friendshipCount)
	}

	// Verify the friend request was deleted.
	pendingCount, err := integrationClient.FriendRequest.Query().
		Where(
			friendrequest.RequesterIDEQ(sender.ID),
			friendrequest.RecipientIDEQ(recipient.ID),
		).
		Count(ctx)
	if err != nil {
		t.Fatalf("counting friend requests failed: %v", err)
	}
	if pendingCount != 0 {
		t.Fatalf("expected friend request to be deleted, but %d remain", pendingCount)
	}
}

func TestAcceptFriendRequestFailsWithoutPendingRequest(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	// Try to accept without a pending request.
	_, err := resolver.Mutation().AcceptFriendRequest(
		reqctx.WithUserID(ctx, bob.ID),
		models.AcceptFriendRequestInput{
			SenderID: alice.ID,
		},
	)
	if err == nil {
		t.Fatal("expected acceptFriendRequest to fail without a pending request, got nil")
	}
}

func TestAcceptFriendRequestFailsWhenNotRecipient(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	sender := mustCreateUser(t, ctx)
	recipient := mustCreateUser(t, ctx)

	// Create a pending friend request from sender to recipient.
	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(sender.ID).
		SetRecipientID(recipient.ID).
		Save(reqctx.WithUserID(ctx, sender.ID))
	if err != nil {
		t.Fatalf("creating friend request failed: %v", err)
	}

	// The sender tries to accept their own request (they are not the recipient).
	_, err = resolver.Mutation().AcceptFriendRequest(
		reqctx.WithUserID(ctx, sender.ID),
		models.AcceptFriendRequestInput{
			SenderID: sender.ID,
		},
	)
	if err == nil {
		t.Fatal("expected acceptFriendRequest to fail when caller is not the recipient, got nil")
	}
}

func TestAcceptFriendRequestRequiresAuthentication(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	sender := mustCreateUser(t, ctx)

	_, err := resolver.Mutation().AcceptFriendRequest(
		ctx,
		models.AcceptFriendRequestInput{
			SenderID: sender.ID,
		},
	)
	if err == nil {
		t.Fatal("expected acceptFriendRequest to fail without authentication, got nil")
	}
}
