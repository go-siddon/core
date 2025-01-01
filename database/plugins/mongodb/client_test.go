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
		ExposedPorts: []string{"27017/tcp"},
	}
	mongoClient, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: client,
		Started:          true,
	})
	url, _ := mongoClient.Endpoint(context.Background(), "")

	t.Run("Test Client Connection", func(t *testing.T) {
		_, err := New("mongodb://"+url, "test-db")
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	testcontainers.CleanupContainer(t, mongoClient)
	require.NoError(t, err)
}
