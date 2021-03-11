package bike

import (
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/go-pg/pg"
	"github.com/rickcorilaco/api-bike-v2/connection"
)

type Repository interface {
	Start() (err error)
	GetByFilter(filte Filter) (bikes []Bike, err error)
	GetByID(bikeID string) (bike Bike, err error)
	Register(bike Bike) (bikeID string, err error)
	Update(bike Bike) (err error)
	Delete(bikeID string) (err error)
}

func NewRepository(conn connection.Connection) (repository Repository, err error) {
	switch c := conn.Interface().(type) {
	case *pg.DB:
		repository, err = NewPostgresORMRepository(c)
	case *firestore.Client:
		repository, err = NewFirestoreRepository(c)
	case *connection.MockClient:
		repository, err = NewMockRepository(c)
	default:
		err = errors.New("connection type is not implemented")
	}

	return
}
