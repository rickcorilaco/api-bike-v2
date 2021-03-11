package connection

import (
	"errors"

	"github.com/rickcorilaco/go/env"
)

const (
	PostegresORMKind = "postgres-orm"
	FirestoreKind    = "firestore"
	MockKind         = "mock"
)

type Connection interface {
	Interface() (i interface{})
	Close() (err error)
}

type Config struct {
	Kind     string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	FilePath string
}

func New(config Config) (conn Connection, err error) {
	switch config.Kind {
	case PostegresORMKind:
		conn, err = NewPostgresORMConnection(config)
	case FirestoreKind:
		conn, err = NewFirestoreConnection(config)
	case MockKind:
		conn, err = NewMockConnection(config)
	default:
		err = errors.New("invalid kind of connection")
	}

	return
}

func MustNew(config Config) (conn Connection) {
	conn, err := New(config)
	if err != nil {
		panic(err)
	}

	return
}

func NewFromEnv() (conn Connection, err error) {
	config := Config{
		Kind:     env.MustString("connection.kind"),
		Host:     env.TryString("connection.host"),
		Port:     env.TryString("connection.port"),
		Username: env.TryString("connection.user"),
		Password: env.TryString("connection.password"),
		Name:     env.TryString("connection.name"),
		FilePath: env.TryString("connection.file_path"),
	}

	return New(config)
}

func MustNewFromEnv() (conn Connection) {
	conn, err := NewFromEnv()
	if err != nil {
		panic(err)
	}

	return
}
