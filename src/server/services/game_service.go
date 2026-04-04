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
	ErrInvalidOpponentSpec    = errors.New("invalid opponent specification")
	ErrNoAvailableAIMove      = errors.New("no available ai move")
)

type TicTacToeOpponentSpec struct {
	Kind   gameplayer.Kind
	UserID int
}

func NewGameService(client *ent.Client) (*GameService, error) {
	if client == nil {
		return nil, fmt.Errorf("services.NewGameService requires a non-nil dbClient")
	}

	return &GameService{dbClient: client}, nil
}

func (g *GameService) CreateTicTacToeGame(ctx context.Context, userID int, opponent TicTacToeOpponentSpec) (*ent.Game, error) {
	if opponent.Kind == gameplayer.KindHUMAN && opponent.UserID == userID {
		return nil, ErrInvalidOpponentSpec
	}

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
		Save(ctx)

	if err != nil {
		return nil, err
	}

	creatorPlayer, err := tx.GamePlayer.Create().
		SetUserID(userID).
		SetGameID(newGame.ID).
		SetKind(gameplayer.KindHUMAN).
		SetMarker(TictactoeX).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	opponentBuilder := tx.GamePlayer.Create().
		SetGameID(newGame.ID).
		SetMarker(TictactoeO)

	opponentBuilder = opponentBuilder.SetKind(opponent.Kind)
	if opponent.Kind == gameplayer.KindHUMAN {
		opponentBuilder = opponentBuilder.SetUserID(opponent.UserID)
	}

	if _, err := opponentBuilder.Save(ctx); err != nil {
		return nil, err
	}

	newGame, err = tx.Game.UpdateOneID(newGame.ID).
		SetCurrentTurnPlayerID(creatorPlayer.ID).
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

	currentPlayer, err := gameState.QueryCurrentTurnPlayer().Only(ctx)
	if err != nil {
		return nil, err
	}

	if currentPlayer.Kind != gameplayer.KindHUMAN || currentPlayer.UserID == nil || *currentPlayer.UserID != userID {
		return nil, ErrNotYourTurn
	}

	if currentPlayer.Marker != TictactoeX && currentPlayer.Marker != TictactoeO {
		return nil, fmt.Errorf("invalid marker: %s", currentPlayer.Marker)
	}

	updatedGame, err := g.applyTicTacToeMove(ctx, tx, gameState, currentPlayer, row, col)
	if err != nil {
		return nil, err
	}

	updatedGame, err = g.advanceTicTacToeAIPlayers(ctx, tx, updatedGame)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return updatedGame.Unwrap(), nil
}

func (g *GameService) applyTicTacToeMove(
	ctx context.Context,
	tx *ent.Tx,
	gameState *ent.Game,
	currentPlayer *ent.GamePlayer,
	row int,
	col int,
) (*ent.Game, error) {
	if gameState.Board[row][col] != "" {
		return nil, ErrCellAlreadyTaken
	}

	gameState.Board[row][col] = currentPlayer.Marker

	if checkWinner(gameState.Board) != "" {
		return tx.Game.UpdateOneID(gameState.ID).
			SetBoard(gameState.Board).
			SetWinnerPlayerID(currentPlayer.ID).
			SetStatus(game.StatusWON).
			ClearCurrentTurnPlayer().
			Save(ctx)
	}

	if isBoardFull(gameState.Board) {
		return tx.Game.UpdateOneID(gameState.ID).
			SetBoard(gameState.Board).
			SetStatus(game.StatusDRAW).
			ClearWinnerPlayer().
			ClearCurrentTurnPlayer().
			Save(ctx)
	}

	otherPlayer, err := tx.GamePlayer.Query().
		Where(
			gameplayer.GameID(gameState.ID),
			gameplayer.IDNEQ(currentPlayer.ID),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return tx.Game.UpdateOneID(gameState.ID).
		SetBoard(gameState.Board).
		SetCurrentTurnPlayerID(otherPlayer.ID).
		Save(ctx)
}

func (g *GameService) advanceTicTacToeAIPlayers(ctx context.Context, tx *ent.Tx, gameState *ent.Game) (*ent.Game, error) {
	for gameState.Status == game.StatusIN_PROGRESS {
		currentPlayer, err := gameState.QueryCurrentTurnPlayer().Only(ctx)
		if err != nil {
			return nil, err
		}

		if currentPlayer.Kind != gameplayer.KindAI {
			return gameState, nil
		}

		row, col, err := chooseTicTacToeAIMove(gameState.Board, currentPlayer.Marker)
		if err != nil {
			return nil, err
		}

		gameState, err = g.applyTicTacToeMove(ctx, tx, gameState, currentPlayer, row, col)
		if err != nil {
			return nil, err
		}
	}

	return gameState, nil
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

func chooseTicTacToeAIMove(board [][]string, marker string) (int, int, error) {
	opponentMarker := TictactoeX
	if marker == TictactoeX {
		opponentMarker = TictactoeO
	}

	for _, candidateMarker := range []string{marker, opponentMarker} {
		if row, col, ok := findWinningMove(board, candidateMarker); ok {
			return row, col, nil
		}
	}

	preferredMoves := [][2]int{
		{1, 1},
		{0, 0},
		{0, 2},
		{2, 0},
		{2, 2},
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},
	}

	for _, move := range preferredMoves {
		if board[move[0]][move[1]] == "" {
			return move[0], move[1], nil
		}
	}

	return 0, 0, ErrNoAvailableAIMove
}

func findWinningMove(board [][]string, marker string) (int, int, bool) {
	for rowIndex, row := range board {
		for colIndex := range row {
			if board[rowIndex][colIndex] != "" {
				continue
			}

			nextBoard := cloneBoard(board)
			nextBoard[rowIndex][colIndex] = marker
			if checkWinner(nextBoard) == marker {
				return rowIndex, colIndex, true
			}
		}
	}

	return 0, 0, false
}

func cloneBoard(board [][]string) [][]string {
	cloned := make([][]string, len(board))
	for rowIndex, row := range board {
		cloned[rowIndex] = append([]string(nil), row...)
	}
	return cloned
}
