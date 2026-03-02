package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"puzzlr.gg/src/server/build"
	"puzzlr.gg/src/server/controllers"
	"puzzlr.gg/src/server/db"
	ent "puzzlr.gg/src/server/db/ent/codegen"
	"puzzlr.gg/src/server/middleware"
	"puzzlr.gg/src/server/util"
	"puzzlr.gg/src/server/views"
)

func main() {

	if util.FileExists(".env") {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env file with error: %v", err)
		}
	}
	build.RunWebpack()

	dbClient, err := db.CreateDatabaseClientAndRunMigrations()
	if err != nil {
		log.Fatalf("Failed to initialize database client with error: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("error closing database client connection with error: %v", err)
		}
	}(dbClient)

	userController := controllers.NewUserController(
		dbClient,
	)

	r := chi.NewRouter()

	// serve static assets from `/assets`
	fs := http.FileServer(http.Dir("src/assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		foo := views.ReactPage("Foo", "index")
		err := foo.Render(w)
		if err != nil {
			return
		}
	})

	srv := build.CreateGraphqlServer(dbClient)
	r.Group(func(r chi.Router) {
		r.Use(middleware.SessionMiddleware)
		r.Handle("/graphql", srv)
		r.Post("/logout", userController.HandleLogout)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Handle("/graphql_playground", playground.Handler("GraphQL playground", "/graphql"))
	})

	r.Post("/login", userController.HandleLogin)
	r.Post("/register", userController.HandleRegister)

	fmt.Printf("Starting server...")
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed starting server %v", err))
		return
	}
}
