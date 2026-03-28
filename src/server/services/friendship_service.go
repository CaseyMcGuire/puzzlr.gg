package services

import (
	"context"
	"fmt"

	ent "puzzlr.gg/src/server/db/ent/codegen"
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
