package main

import (
	"fmt"
	"gameboard/src/server/build"
	"gameboard/src/server/controllers"
	"gameboard/src/server/db"
	ent "gameboard/src/server/db/ent/codegen"
	"gameboard/src/server/util"
	"gameboard/src/server/views"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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
		*dbClient,
	)

	// serve static assets from `/assets`
	fs := http.FileServer(http.Dir("src/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		foo := views.ReactPage("Foo", "index")
		err := foo.Render(w)
		if err != nil {
			return
		}
	})

	srv := build.CreateGraphqlServer()

	http.Handle("/graphql_playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	http.HandleFunc("/login", userController.HandleLogin)
	http.HandleFunc("/register", userController.HandleRegister)
	http.HandleFunc("/logout", userController.HandleLogout)

	fmt.Printf("Starting server...")
	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("Failed starting server")
		return
	}
}
