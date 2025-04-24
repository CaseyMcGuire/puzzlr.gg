package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	ent "puzzlr.gg/src/server/db/ent/codegen"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ent *ent.Client
}
