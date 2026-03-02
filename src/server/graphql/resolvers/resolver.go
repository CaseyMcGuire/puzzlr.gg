package resolvers

import (
	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Ent         *ent.Client
	GameService *services.GameService
}
