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

func NewRepository(connection connection.Connection) (repository Repository, err error) {
	switch conn := connection.Interface().(type) {
	case *pg.DB:
		repository, err = NewPostgresORMRepository(conn)
	case *firestore.Client:
		repository, err = NewFirestoreRepository(conn)
	default:
		err = errors.New("connection type is not implemented")
	}

	return
}
