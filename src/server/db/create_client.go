package db

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	ent "gameboard/src/server/db/ent/codegen"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func CreateDatabaseClientAndRunMigrations() (*ent.Client, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf("missing database parameters")
	}
	client, err := ent.Open(
		dialect.Postgres,
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost,
			dbPort,
			dbUser,
			dbName,
			dbPass,
		),
	)

	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, nil
}
