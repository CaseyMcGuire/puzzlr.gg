package services

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/codegen/gameplayer"
	"puzzlr.gg/src/server/db/ent/codegen/predicate"
)

type GameService struct {
	dbClient *ent.Client
}

const (
	TictactoeX = "X"
	TictactoeO = "O"
)

var (
	ErrInvalidMoveCoordinates = errors.New("invalid row or column")
	ErrGameNotInProgress      = errors.New("game is not in progress")
	ErrNotYourTurn            = errors.New("it is not your turn")
	ErrCellAlreadyTaken       = errors.New("cell already taken")
)

func NewGameService(client *ent.Client) (*GameService, error) {
	if client == nil {
		return nil, fmt.Errorf("services.NewGameService requires a non-nil dbClient")
	}

	return &GameService{dbClient: client}, nil
}

func (g *GameService) CreateTicTacToeGame(ctx context.Context, userID int, opponentID int) (*ent.Game, error) {
	emptyBoard := [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	tx, err := g.dbClient.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	newGame, err := tx.Game.
		Create().
		SetType(game.TypeTIC_TAC_TOE).
		SetBoard(emptyBoard).
		SetCurrentTurnID(userID).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	_, err = tx.GamePlayer.CreateBulk(
		tx.GamePlayer.Create().
			SetUserID(userID).
			SetGameID(newGame.ID).
			SetMarker(TictactoeX),
		tx.GamePlayer.Create().
			SetUserID(opponentID).
			SetGameID(newGame.ID).
			SetMarker(TictactoeO),
	).Save(ctx)

	if err != nil {
		return nil, err
	}

	newGame, err = tx.Game.UpdateOneID(newGame.ID).
		SetStatus(game.StatusIN_PROGRESS).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return newGame.Unwrap(), nil
}

func (g *GameService) MakeTicTacToeMove(
	ctx context.Context,
	gameId int,
	userID int,
	row int,
	col int,
) (*ent.Game, error) {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return nil, ErrInvalidMoveCoordinates
	}

	tx, err := g.dbClient.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	gameState, err := tx.Game.Query().
		Where(
			game.ID(gameId),
			predicate.Game(func(s *sql.Selector) {
				s.ForUpdate(sql.WithLockAction(sql.NoWait))
			}),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	if gameState.Status != game.StatusIN_PROGRESS {
		return nil, ErrGameNotInProgress
	}

	currentPlayer, err := gameState.CurrentTurn(ctx)
	if err != nil {
		return nil, err
	}

	if currentPlayer.ID != userID {
		return nil, ErrNotYourTurn
	}

	gp, err := tx.GamePlayer.Query().
		Where(
			gameplayer.UserID(userID),
			gameplayer.GameID(gameId),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	if gp.Marker != TictactoeX && gp.Marker != TictactoeO {
		return nil, fmt.Errorf("invalid marker: %s", gp.Marker)
	}

	if gameState.Board[row][col] != "" {
		return nil, ErrCellAlreadyTaken
	}

	gameState.Board[row][col] = gp.Marker

	if checkWinner(gameState.Board) != "" {
		updatedGame, err := tx.Game.UpdateOneID(gameId).
			SetBoard(gameState.Board).
			SetWinnerID(currentPlayer.ID).
			SetStatus(game.StatusWON).
			ClearCurrentTurn().
			Save(ctx)
		if err != nil {
			return nil, err
		}
		if err := tx.Commit(); err != nil {
			return nil, err
		}
		return updatedGame.Unwrap(), nil
	} else if isBoardFull(gameState.Board) {
		updatedGame, err := tx.Game.UpdateOneID(gameId).
			SetBoard(gameState.Board).
			SetStatus(game.StatusDRAW).
			ClearWinner().
			ClearCurrentTurn().
			Save(ctx)
		if err != nil {
			return nil, err
		}
		if err := tx.Commit(); err != nil {
			return nil, err
		}
		return updatedGame.Unwrap(), nil
	}

	otherPlayer, err := tx.GamePlayer.Query().
		Where(
			gameplayer.GameID(gameId),
			gameplayer.UserIDNEQ(userID),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	updatedGame, err := tx.Game.UpdateOneID(gameId).
		SetBoard(gameState.Board).
		SetCurrentTurnID(otherPlayer.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return updatedGame.Unwrap(), nil
}

func checkWinner(board [][]string) string {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	// Check diagonals
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return ""
}

func isBoardFull(board [][]string) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}
