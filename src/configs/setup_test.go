package configs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestConnectDb(t *testing.T) {
	uri := "mongodb://localhost:27017" // use a table approach for multiple inputs of uri
	client, ctx, cancel, err := ConnectDb(uri)
	defer cancel()
	defer client.Disconnect(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	err = Ping(client, ctx)
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {

	client, ctx, cancel := createMockClientAndContext()
	defer cancel()

	Close(client, ctx, cancel)
}

func TestPing(t *testing.T) {
	client, ctx, cancel := createMockClientAndContext()
	defer cancel()

	err := Ping(client, ctx)
	assert.NoError(t, err)
}

func createMockClientAndContext() (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	client, _ := mongo.Connect(ctx, options.Client())

	return client, ctx, cancel
}
