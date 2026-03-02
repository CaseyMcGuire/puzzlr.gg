package metadata

import (
	"encoding/json"
	"fmt"

	"puzzlr.gg/src/server/db/ent/codegen/game"
)

type GameMetadata interface {
	gameMetadata()
}

func DecodeGameMetadata(t game.Type, raw json.RawMessage) (GameMetadata, error) {
	switch t {
	case game.TypeTIC_TAC_TOE:
		var m TicTacToeMetadata
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return &m, nil

	}
	return nil, fmt.Errorf("unknown game type %s", t)
}
