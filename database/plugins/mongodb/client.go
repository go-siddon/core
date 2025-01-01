package mongodb

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mongo_instance *mongo.Client

var mongo_init sync.Once

type Mongo struct {
	client *mongo.Database
}

func New(url, db string) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if mongo_instance == nil {
		mongo_init.Do(func() {
			client, err := mongo.Connect(options.Client().ApplyURI(url))
			if err != nil {
				panic(err)
			}
			mongo_instance = client
		})
		err := mongo_instance.Ping(ctx, nil)
		if err != nil {
			return &Mongo{}, err
		}
	}
	return &Mongo{
		client: mongo_instance.Database(db),
	}, nil
}
