package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client *mongo.Client
	ctx    context.Context
}

func NewMongoConnection(config Config) (conn Connection, err error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Username, config.Password, config.Host, config.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		return
	}

	conn = MongoConnection{client: client}
	return
}

func (conn MongoConnection) Interface() (i interface{}) {
	return conn.client
}

func (conn MongoConnection) Close() (err error) {
	err = conn.client.Disconnect(conn.ctx)
	return
}
