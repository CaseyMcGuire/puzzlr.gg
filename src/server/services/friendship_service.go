package services

import (
	"context"
	"fmt"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/friendrequest"
	"puzzlr.gg/src/server/db/ent/codegen/friendship"
	"puzzlr.gg/src/server/graphql/models"
)

type FriendshipService struct {
	dbClient *ent.Client
}

func NewFriendshipService(dbClient *ent.Client) (*FriendshipService, error) {
	if dbClient == nil {
		return nil, fmt.Errorf("services.NewFriendshipService requires a non-nil dbClient")
	}
	return &FriendshipService{dbClient: dbClient}, nil
}

func (f *FriendshipService) CreateFriendRequest(ctx context.Context, requestorID int, recipientID int) (*ent.FriendRequest, error) {
	return f.dbClient.FriendRequest.Create().SetRequesterID(requestorID).SetRecipientID(recipientID).Save(ctx)
}

func (f *FriendshipService) GetViewerFriendshipStatus(ctx context.Context, viewerID int, targetUserID int) (models.ViewerFriendshipStatus, error) {
	areFriends, err := f.dbClient.Friendship.Query().
		Where(
			friendship.Or(
				friendship.And(
					friendship.UserIDEQ(viewerID),
					friendship.FriendIDEQ(targetUserID),
				),
				friendship.And(
					friendship.UserIDEQ(targetUserID),
					friendship.FriendIDEQ(viewerID),
				),
			),
		).
		Exist(ctx)
	if err != nil {
		return "", err
	}
	if areFriends {
		return models.ViewerFriendshipStatusFriends, nil
	}

	requestSent, err := f.dbClient.FriendRequest.Query().
		Where(
			friendrequest.RequesterIDEQ(viewerID),
			friendrequest.RecipientIDEQ(targetUserID),
		).
		Exist(ctx)
	if err != nil {
		return "", err
	}
	if requestSent {
		return models.ViewerFriendshipStatusRequestSent, nil
	}

	requestReceived, err := f.dbClient.FriendRequest.Query().
		Where(
			friendrequest.RequesterIDEQ(targetUserID),
			friendrequest.RecipientIDEQ(viewerID),
		).
		Exist(ctx)
	if err != nil {
		return "", err
	}
	if requestReceived {
		return models.ViewerFriendshipStatusRequestReceived, nil
	}

	return models.ViewerFriendshipStatusNotFriends, nil
}
