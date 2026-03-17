//go:build integration

package resolvers_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"entgo.io/ent/dialect"
	ent "puzzlr.gg/src/server/db/ent/codegen"
	_ "puzzlr.gg/src/server/db/ent/codegen/runtime"
	"puzzlr.gg/src/server/graphql/resolvers"
	"puzzlr.gg/src/server/services"
)

const (
	testPostgresUser     = "postgres"
	testPostgresPassword = "postgres"
	testPostgresDatabase = "puzzlr_test"
)

var integrationClient *ent.Client

func TestMain(m *testing.M) {
	ctx := context.Background()

	if os.Getenv("RUN_INTEGRATION_TESTS") != "1" {
		log.Print("RUN_INTEGRATION_TESTS is not set to 1, skipping integration tests")
		os.Exit(0)
	}

	container, dsn, err := startPostgresContainer(ctx)
	if err != nil {
		log.Printf("failed to start postgres test container: %v", err)
		os.Exit(1)
	}

	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		_ = container.Terminate(ctx)
		log.Printf("failed to open ent postgres client: %v", err)
		os.Exit(1)
	}

	if err := client.Schema.Create(ctx); err != nil {
		_ = client.Close()
		_ = container.Terminate(ctx)
		log.Printf("failed to run schema migration: %v", err)
		os.Exit(1)
	}

	integrationClient = client
	code := m.Run()

	if err := client.Close(); err != nil {
		log.Printf("failed to close ent client: %v", err)
	}
	if err := container.Terminate(ctx); err != nil {
		log.Printf("failed to terminate postgres test container: %v", err)
	}

	os.Exit(code)
}

func startPostgresContainer(ctx context.Context) (container testcontainers.Container, dsn string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("start container panic: %v", r)
		}
	}()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:16-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     testPostgresUser,
			"POSTGRES_PASSWORD": testPostgresPassword,
			"POSTGRES_DB":       testPostgresDatabase,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").
			WithOccurrence(2).
			WithStartupTimeout(2 * time.Minute),
	}

	container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", fmt.Errorf("start container: %w", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		_ = container.Terminate(ctx)
		return nil, "", fmt.Errorf("read container host: %w", err)
	}

	port, err := container.MappedPort(ctx, "5432/tcp")
	if err != nil {
		_ = container.Terminate(ctx)
		return nil, "", fmt.Errorf("read container port: %w", err)
	}

	dsn = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port.Port(),
		testPostgresUser,
		testPostgresDatabase,
		testPostgresPassword,
	)

	return container, dsn, err
}

func newTestResolver() *resolvers.Resolver {
	gameService, err := services.NewGameService(integrationClient)
	if err != nil {
		panic(err)
	}

	return &resolvers.Resolver{
		Ent:         integrationClient,
		GameService: gameService,
	}
}
