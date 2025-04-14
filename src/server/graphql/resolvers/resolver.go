package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import "gameboard/src/server/graphql/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*models.Todo
}
