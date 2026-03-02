package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/codegen/hook"
)

func ValidatePlayerCountOnCreate(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		userIds := m.UserIDs()
		if len(userIds) == 0 {
			return nil, fmt.Errorf("cannot create a game without any players")
		}
		gameType, exists := m.GetType()
		if !exists {
			return nil, fmt.Errorf("cannot validate player count without a game type")
		}
		minPlayers, maxPlayers := getGamePlayerCounts(gameType)
		if len(userIds) < minPlayers || len(userIds) > maxPlayers {
			return nil, fmt.Errorf("invalid player count: %d. Must be between %d and %d", len(userIds), minPlayers, maxPlayers)
		}
		return next.Mutate(ctx, m)
	})
}

func ValidatePlayerCountOnUpdate(next ent.Mutator) ent.Mutator {
	return hook.GameFunc(func(ctx context.Context, m *codegen.GameMutation) (ent.Value, error) {
		addedUserIDs := m.UserIDs()
		removedUserIDs := m.RemovedUserIDs()
		if len(addedUserIDs) > 0 || len(removedUserIDs) > 0 || m.UserCleared() {
			return nil, fmt.Errorf("cannot change players in a game")
		}
		return next.Mutate(ctx, m)
	})
}

func getGamePlayerCounts(gameType game.Type) (minimum, maximum int) {
	switch gameType {
	case game.TypeTIC_TAC_TOE:
		return 2, 2
	}
	return -1, -1
}
