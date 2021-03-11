package connection

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreConnection struct {
	client *firestore.Client
}

func NewFirestoreConnection(config Config) (conn Connection, err error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.FilePath)

	client, err := firestore.NewClient(ctx, config.Name, opt)
	if err != nil {
		panic(err)
	}

	conn = FirestoreConnection{client: client}
	return
}

func (conn FirestoreConnection) Interface() (i interface{}) {
	return conn.client
}

func (conn FirestoreConnection) Close() (err error) {
	err = conn.client.Close()
	return
}
