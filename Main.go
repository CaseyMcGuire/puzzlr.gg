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
	"puzzlr.gg/src/server/session"
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

	sessionManager, err := session.NewSessionManager(build.CreateCookieStore())
	if err != nil {
		log.Fatalf("Failed to initialize session manager: %v", err)
	}

	userController, err := controllers.NewUserController(
		dbClient,
		sessionManager,
	)
	if err != nil {
		log.Fatalf("Failed to initialize user controller: %v", err)
	}

	sessionController, err := controllers.NewSessionController(
		dbClient,
		sessionManager,
	)
	if err != nil {
		log.Fatalf("Failed to initialize session controller: %v", err)
	}

	sessionMiddleware, err := middleware.SessionMiddleware(sessionManager)
	if err != nil {
		log.Fatalf("Failed to initialize session middleware: %v", err)
	}

	authMiddleware, err := middleware.AuthMiddleware(sessionManager)
	if err != nil {
		log.Fatalf("Failed to initialize auth middleware: %v", err)
	}

	r := chi.NewRouter()

	// serve static assets from `/assets`
	fs := http.FileServer(http.Dir("src/assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	renderReactPage := func(title string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := views.ReactPage(title, "index").Render(w)
			if err != nil {
				http.Error(w, "failed to render page", http.StatusInternalServerError)
			}
		}
	}

	r.Get("/", renderReactPage("Puzzlr"))
	r.Get("/tictactoe", renderReactPage("Tic Tac Toe"))
	r.Get("/user/{id}", renderReactPage("User Profile"))

	srv, err := build.CreateGraphqlServer(dbClient)
	if err != nil {
		log.Fatalf("Failed to initialize graphql server: %v", err)
	}
	r.Group(func(r chi.Router) {
		r.Use(sessionMiddleware)
		r.Handle("/graphql", srv)
		r.Post("/logout", sessionController.HandleLogout)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Handle("/graphql_playground", playground.Handler("GraphQL playground", "/graphql"))
	})

	r.Get("/login", sessionController.HandleLoginGet)
	r.Post("/session/create", sessionController.HandleLoginPost)
	r.Get("/register", userController.HandleRegisterGet)
	r.Post("/user/create", userController.HandleRegisterPost)

	fmt.Printf("Starting server...")
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed starting server %v", err))
		return
	}
}
