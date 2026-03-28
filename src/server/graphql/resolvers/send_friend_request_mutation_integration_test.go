//go:build integration

package resolvers_test

import (
	"context"
	"testing"

	"puzzlr.gg/src/server/db/ent/codegen/friendrequest"
	"puzzlr.gg/src/server/graphql/models"
	"puzzlr.gg/src/server/reqctx"
)

func TestSendFriendRequestResolverSuccess(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	requester := mustCreateUser(t, ctx)
	recipient := mustCreateUser(t, ctx)

	result, err := resolver.Mutation().SendFriendRequest(
		reqctx.WithUserID(ctx, requester.ID),
		models.SendFriendRequestInput{
			RecipientID: recipient.ID,
		},
	)
	if err != nil {
		t.Fatalf("sendFriendRequest returned an error: %v", err)
	}
	success, ok := result.(*models.SendFriendRequestSuccess)
	if !ok {
		t.Fatalf("expected success result, got %T", result)
	}
	if success.Recipient.ID != recipient.ID {
		t.Fatalf("expected recipient ID %d, got %d", recipient.ID, success.Recipient.ID)
	}

	requests, err := integrationClient.FriendRequest.
		Query().
		Where(
			friendrequest.RequesterIDEQ(requester.ID),
			friendrequest.RecipientIDEQ(recipient.ID),
		).
		All(ctx)
	if err != nil {
		t.Fatalf("querying friend requests failed: %v", err)
	}
	assertOnlyFriendRequest(t, requests, requester.ID, recipient.ID)
}

func TestSendFriendRequestResolverRequiresUserInContext(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	recipient := mustCreateUser(t, ctx)

	result, err := resolver.Mutation().SendFriendRequest(
		ctx,
		models.SendFriendRequestInput{
			RecipientID: recipient.ID,
		},
	)
	if err != nil {
		t.Fatalf("sendFriendRequest returned an unexpected error: %v", err)
	}

	failure, ok := result.(*models.SendFriendRequestError)
	if !ok {
		t.Fatalf("expected error result, got %T", result)
	}
	if failure.Message != "Log in to send a friend request." {
		t.Fatalf("unexpected error message: %q", failure.Message)
	}
}

func TestSendFriendRequestResolverReturnsFriendlyErrorForDuplicatePending(t *testing.T) {
	ctx := context.Background()
	resolver := newTestResolver()

	requester := mustCreateUser(t, ctx)
	recipient := mustCreateUser(t, ctx)

	_, err := resolver.Mutation().SendFriendRequest(
		reqctx.WithUserID(ctx, requester.ID),
		models.SendFriendRequestInput{
			RecipientID: recipient.ID,
		},
	)
	if err != nil {
		t.Fatalf("initial sendFriendRequest returned an error: %v", err)
	}

	result, err := resolver.Mutation().SendFriendRequest(
		reqctx.WithUserID(ctx, requester.ID),
		models.SendFriendRequestInput{
			RecipientID: recipient.ID,
		},
	)
	if err != nil {
		t.Fatalf("duplicate sendFriendRequest returned an unexpected error: %v", err)
	}

	failure, ok := result.(*models.SendFriendRequestError)
	if !ok {
		t.Fatalf("expected error result, got %T", result)
	}
	if failure.Message != "A friend request is already pending." {
		t.Fatalf("unexpected error message: %q", failure.Message)
	}
}
