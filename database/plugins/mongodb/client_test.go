package mongodb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestNewClient(t *testing.T) {
	client := testcontainers.ContainerRequest{
		Image:        "mongo:8.0",
		ExposedPorts: []string{"27017/tcp", "27017:27017"},
	}
	mongoClient, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: client,
		Started:          true,
	})

	t.Run("Test Client Connection", func(t *testing.T) {
		_, err := New("mongodb://localhost:27017", "test-db")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	testcontainers.CleanupContainer(t, mongoClient)
	require.NoError(t, err)
}
