package services

import (
	"context"

	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/game"
)

type GameService struct {
	dbClient *ent.Client
}

func NewGameService(client *ent.Client) *GameService {
	return &GameService{dbClient: client}
}

func (g *GameService) CreateTicTacToeGame(ctx context.Context, userID int, opponentID int) (*ent.Game, error) {
	emptyBoard := [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	return g.dbClient.Game.
		Create().
		SetType(game.TypeTIC_TAC_TOE).
		SetBoard(emptyBoard).
		AddUserIDs(userID, opponentID).
		SetCurrentTurnID(userID).
		Save(ctx)

}
