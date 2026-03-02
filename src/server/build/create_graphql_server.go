package build

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/ast"
	ent "puzzlr.gg/src/server/db/ent/codegen"
	graphql "puzzlr.gg/src/server/graphql/generated"
	"puzzlr.gg/src/server/graphql/resolvers"
	"puzzlr.gg/src/server/services"
)

func CreateGraphqlServer(ent *ent.Client) *handler.Server {
	srv := handler.New(graphql.NewExecutableSchema(
		graphql.Config{
			Resolvers: &resolvers.Resolver{
				Ent:         ent,
				GameService: services.NewGameService(ent),
			}},
	))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}
