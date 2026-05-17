package auth

import (
	"context"
	"database/sql"
	"runtime"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestRegisterLoginFlow(t *testing.T) {
	// Skip test on Windows due to Docker container limitations
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Docker-based test on Windows")
	}

	ctx := context.Background()
	container, db := setupPostgresContainer(ctx, t)
	defer container.Terminate(ctx)
	defer db.Close()
	_, err := db.ExecContext(ctx, `CREATE SCHEMA IF NOT EXISTS gateway_service; SET search_path = gateway_service; CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, email TEXT UNIQUE NOT NULL, name TEXT NOT NULL, password TEXT NOT NULL, role TEXT NOT NULL)`)
	require.NoError(t, err)

	svc := NewService(db, "test-secret")
	user, token, err := svc.Register(ctx, "test@example.com", "password", "Test User")
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.Equal(t, "test@example.com", user.Email)

	logged, token2, err := svc.Login(ctx, "test@example.com", "password")
	require.NoError(t, err)
	require.NotEmpty(t, token2)
	require.Equal(t, user.ID, logged.ID)
}

func setupPostgresContainer(ctx context.Context, t *testing.T) (testcontainers.Container, *sql.DB) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "food_delivery",
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "password",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{ContainerRequest: req, Started: true})
	require.NoError(t, err)
	endpoint, err := container.Endpoint(ctx, "5432/tcp")
	require.NoError(t, err)

	db, err := sql.Open("pgx", "postgres://postgres:password@"+endpoint+"/food_delivery?sslmode=disable")
	require.NoError(t, err)
	require.NoError(t, db.PingContext(ctx))
	return container, db
}
