//go:build integration

package resolvers_test

import (
	"context"
	"errors"
	"testing"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/friendrequest"
	"puzzlr.gg/src/server/db/ent/schema"
	"puzzlr.gg/src/server/reqctx"
)

func TestUserQueryFriendRequestsReturnsDirectedRequests(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)
	charlie := mustCreateUser(t, ctx)

	outgoing, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID))
	if err != nil {
		t.Fatalf("creating outgoing friend request failed: %v", err)
	}

	incoming, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(charlie.ID).
		SetRecipientID(alice.ID).
		Save(reqctx.WithUserID(ctx, charlie.ID))
	if err != nil {
		t.Fatalf("creating incoming friend request failed: %v", err)
	}

	sentRequests, err := alice.QuerySentFriendRequests().All(ctx)
	if err != nil {
		t.Fatalf("querying sent friend requests failed: %v", err)
	}
	assertOnlyFriendRequest(t, sentRequests, alice.ID, bob.ID)

	receivedRequests, err := alice.QueryReceivedFriendRequests().All(ctx)
	if err != nil {
		t.Fatalf("querying received friend requests failed: %v", err)
	}
	assertOnlyFriendRequest(t, receivedRequests, charlie.ID, alice.ID)

	requester, err := outgoing.QueryRequester().Only(ctx)
	if err != nil {
		t.Fatalf("querying outgoing requester failed: %v", err)
	}
	if requester.ID != alice.ID {
		t.Fatalf("expected requester %d, got %d", alice.ID, requester.ID)
	}

	recipient, err := outgoing.QueryRecipient().Only(ctx)
	if err != nil {
		t.Fatalf("querying outgoing recipient failed: %v", err)
	}
	if recipient.ID != bob.ID {
		t.Fatalf("expected recipient %d, got %d", bob.ID, recipient.ID)
	}

	incomingRecipient, err := incoming.QueryRecipient().Only(ctx)
	if err != nil {
		t.Fatalf("querying incoming recipient failed: %v", err)
	}
	if incomingRecipient.ID != alice.ID {
		t.Fatalf("expected incoming recipient %d, got %d", alice.ID, incomingRecipient.ID)
	}

	friendRequestCount, err := integrationClient.FriendRequest.Query().
		Where(
			friendrequest.Or(
				friendrequest.And(
					friendrequest.RequesterIDEQ(alice.ID),
					friendrequest.RecipientIDEQ(bob.ID),
				),
				friendrequest.And(
					friendrequest.RequesterIDEQ(charlie.ID),
					friendrequest.RecipientIDEQ(alice.ID),
				),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatalf("counting friend request rows failed: %v", err)
	}
	if friendRequestCount != 2 {
		t.Fatalf("expected 2 friend request rows for the queried pairs, got %d", friendRequestCount)
	}
}

func TestFriendRequestRejectsSelfRequest(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)

	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(alice.ID).
		Save(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected self-request to fail, got nil")
	}
	if !ent.IsConstraintError(err) {
		t.Fatalf("expected constraint error, got: %v", err)
	}
}

func TestFriendRequestRejectsDuplicateDirectedRequest(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID))
	if err != nil {
		t.Fatalf("creating initial friend request failed: %v", err)
	}

	_, err = integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected duplicate friend request to fail, got nil")
	}
	if !ent.IsConstraintError(err) {
		t.Fatalf("expected constraint error, got: %v", err)
	}
}

func TestFriendRequestRejectsExistingFriendship(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	if _, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID)); err != nil {
		t.Fatalf("creating initial friend request failed: %v", err)
	}

	if err := integrationClient.User.
		UpdateOneID(bob.ID).
		AddFriendIDs(alice.ID).
		Exec(reqctx.WithUserID(ctx, bob.ID)); err != nil {
		t.Fatalf("accepting friendship failed: %v", err)
	}

	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected friend request to existing friend to fail, got nil")
	}
	if !errors.Is(err, schema.ErrFriendRequestAlreadyFriends) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFriendRequestRejectsReversePendingRequest(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)

	if _, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(bob.ID).
		SetRecipientID(alice.ID).
		Save(reqctx.WithUserID(ctx, bob.ID)); err != nil {
		t.Fatalf("creating initial friend request failed: %v", err)
	}

	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, alice.ID))
	if err == nil {
		t.Fatal("expected reverse pending friend request to fail, got nil")
	}
	if !errors.Is(err, schema.ErrFriendRequestReversePending) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFriendRequestCreateRequiresRequesterActor(t *testing.T) {
	ctx := context.Background()

	alice := mustCreateUser(t, ctx)
	bob := mustCreateUser(t, ctx)
	charlie := mustCreateUser(t, ctx)

	_, err := integrationClient.FriendRequest.
		Create().
		SetRequesterID(alice.ID).
		SetRecipientID(bob.ID).
		Save(reqctx.WithUserID(ctx, charlie.ID))
	if err == nil {
		t.Fatal("expected friend request create with wrong actor to fail, got nil")
	}
	if !errors.Is(err, schema.ErrOnlyRequesterCanCreateFriendRequest) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertOnlyFriendRequest(t *testing.T, got []*ent.FriendRequest, wantRequesterID, wantRecipientID int) {
	t.Helper()

	if len(got) != 1 {
		t.Fatalf("expected 1 friend request, got %d", len(got))
	}
	if got[0].RequesterID != wantRequesterID {
		t.Fatalf("expected requester %d, got %d", wantRequesterID, got[0].RequesterID)
	}
	if got[0].RecipientID != wantRecipientID {
		t.Fatalf("expected recipient %d, got %d", wantRecipientID, got[0].RecipientID)
	}
}
