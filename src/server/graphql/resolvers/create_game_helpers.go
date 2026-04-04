package resolvers

import (
	"puzzlr.gg/src/server/db/ent/codegen/gameplayer"
	"puzzlr.gg/src/server/graphql/models"
	"puzzlr.gg/src/server/services"
)

func buildTicTacToeOpponentSpec(input *models.CreateTicTacToeInput) (services.TicTacToeOpponentSpec, error) {
	if input == nil {
		return services.TicTacToeOpponentSpec{}, services.ErrInvalidOpponentSpec
	}

	if input.HumanOpponent != nil {
		return services.TicTacToeOpponentSpec{
			Kind:   gameplayer.KindHUMAN,
			UserID: input.HumanOpponent.OpponentID,
		}, nil
	}

	if input.AiOpponent != nil {
		return services.TicTacToeOpponentSpec{
			Kind: gameplayer.KindAI,
		}, nil
	}

	return services.TicTacToeOpponentSpec{}, services.ErrInvalidOpponentSpec
}
