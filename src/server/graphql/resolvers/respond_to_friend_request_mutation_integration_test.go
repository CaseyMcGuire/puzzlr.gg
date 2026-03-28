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

func mustCreateFriendRequest(t *testing.T, ctx context.Context, senderID, recipientID int) {
	t.Helper()
	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(senderID).
		SetRecipientID(recipientID).
		Save(reqctx.WithUserID(ctx, senderID))
	if err != nil {
		t.Fatalf("creating friend request failed: %v", err)
	}
}

func TestRespondToFriendRequestAcceptSuccess(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	sender := mustCreateUser(t, ctx)
	recipient := mustCreateUser(t, ctx)
	mustCreateFriendRequest(t, ctx, sender.ID, recipient.ID)

	result, err := resolver.Mutation().RespondToFriendRequest(
		reqctx.WithUserID(ctx, recipient.ID),
		models.RespondToFriendRequestInput{
			SenderID: sender.ID,
			Accept:   true,
		},
	)
	if err != nil {
		t.Fatalf("respondToFriendRequest returned an error: %v", err)
	}

	success, ok := result.(*models.RespondToFriendRequestSuccess)
	if !ok {
		t.Fatalf("expected success result, got %T", result)
	}
	if success.Sender.ID != sender.ID {
		t.Fatalf("expected sender ID %d, got %d", sender.ID, success.Sender.ID)
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

func TestRespondToFriendRequestRejectSuccess(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	sender := mustCreateUser(t, ctx)
	recipient := mustCreateUser(t, ctx)
	mustCreateFriendRequest(t, ctx, sender.ID, recipient.ID)

	result, err := resolver.Mutation().RespondToFriendRequest(
		reqctx.WithUserID(ctx, recipient.ID),
		models.RespondToFriendRequestInput{
			SenderID: sender.ID,
			Accept:   false,
		},
	)
	if err != nil {
		t.Fatalf("respondToFriendRequest returned an error: %v", err)
	}

	if _, ok := result.(*models.RespondToFriendRequestSuccess); !ok {
		t.Fatalf("expected success result, got %T", result)
	}

	// Verify no friendship was created.
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
	if friendshipCount != 0 {
		t.Fatalf("expected 0 friendship rows after rejection, got %d", friendshipCount)
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

func TestRespondToFriendRequestFailsWithoutPendingRequest(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	_, err := resolver.Mutation().RespondToFriendRequest(
		reqctx.WithUserID(ctx, bob.ID),
		models.RespondToFriendRequestInput{
			SenderID: alice.ID,
			Accept:   true,
		},
	)
	if err == nil {
		t.Fatal("expected respondToFriendRequest to fail without a pending request, got nil")
	}
}

func TestRespondToFriendRequestFailsWhenNotRecipient(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	sender := mustCreateUser(t, ctx)
	recipient := mustCreateUser(t, ctx)
	mustCreateFriendRequest(t, ctx, sender.ID, recipient.ID)

	// The sender tries to accept their own request.
	_, err := resolver.Mutation().RespondToFriendRequest(
		reqctx.WithUserID(ctx, sender.ID),
		models.RespondToFriendRequestInput{
			SenderID: sender.ID,
			Accept:   true,
		},
	)
	if err == nil {
		t.Fatal("expected respondToFriendRequest to fail when caller is not the recipient, got nil")
	}
}

func TestRespondToFriendRequestRequiresAuthentication(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	sender := mustCreateUser(t, ctx)

	_, err := resolver.Mutation().RespondToFriendRequest(
		ctx,
		models.RespondToFriendRequestInput{
			SenderID: sender.ID,
			Accept:   true,
		},
	)
	if err == nil {
		t.Fatal("expected respondToFriendRequest to fail without authentication, got nil")
	}
}
