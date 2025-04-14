package main

import (
	"fmt"
	"gameboard/src/server/util"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
)

func main() {

	if util.FileExists(".env") {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Failed to load .env file: %v", err)
		}
	}

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Ensure we have all required variables
	if dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatal("Missing required environment variables in .env file")
	}

	// Run PostgreSQL commands
	commands := []string{
		fmt.Sprintf("CREATE USER %s WITH ENCRYPTED PASSWORD '%s';", dbUser, dbPass),
		fmt.Sprintf("CREATE DATABASE %s;", dbName),
		fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", dbName, dbUser),
	}

	for _, cmd := range commands {
		psqlCmd := exec.Command("psql", "-U", "postgres", "-c", cmd)
		psqlCmd.Stdout = os.Stdout
		psqlCmd.Stderr = os.Stderr

		if err := psqlCmd.Run(); err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}
	}

	fmt.Println("Database setup completed successfully")
}
