//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
	// entgql.WithSchemaGenerator(),
	// entgql.WithSchemaPath("../../graphql/schema/ent.graphqls"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate("./schema", &gen.Config{
		Target:  "./codegen",
		Package: "gameboard/src/server/db/ent/codegen",
	}, entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
