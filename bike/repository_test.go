package bike

import (
	"errors"
	"testing"

	"github.com/rickcorilaco/api-bike-v2/connection"
	"github.com/rickcorilaco/go/env"
)

func TestNewRepository(t *testing.T) {
	err := env.FromFile("../env.json")
	if err != nil {
		t.Error(err)
	}

	conn, err := connection.NewFromEnv()
	if err != nil {
		t.Error(err)
	}

	repository, err := NewRepository(conn)
	if err != nil {
		t.Error(err)
	}

	bike := Bike{
		Model: "Soul TTR1",
	}

	bike.ID, err = repository.Register(bike)
	if err != nil {
		t.Error(err)
	}

	if bike.ID == "" {
		t.Fatal("returned the empty bike id when registering")
	}

	bikes, err := repository.GetByFilter(Filter{})
	if err != nil {
		t.Error(err)
	}

	bikefound := false

	for i := range bikes {
		if bike.ID == bikes[i].ID {
			bikefound = true
			break
		}
	}

	if !bikefound {
		t.Fatal("returned the empty bike id when registering")
	}

	model := "Sense Criterium comp 2021"
	bike.Model = model

	err = repository.Update(bike)
	if err != nil {
		t.Error(err)
	}

	updatedBike, err := repository.GetByID(bike.ID)
	if err != nil {
		t.Error(err)
	}

	if updatedBike.Model != model {
		t.Fatal("the updated bike do not have the expected id when registering")
	}

	err = repository.Delete(bike.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = repository.GetByID(bike.ID)
	if !errors.Is(err, ErrBikeNotFound) {
		t.Fatal("the bike has not been deleted")
	}
}
