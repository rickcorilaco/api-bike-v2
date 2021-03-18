package bike

import (
	"errors"

	"github.com/rickcorilaco/api-bike-v2/connection"
	uuid "github.com/satori/go.uuid"
)

type MockRepository struct {
	bikes []Bike
}

func NewMockRepository(mockConn *connection.MockClient) (mockRepository *MockRepository, err error) {
	bikes := []Bike{}
	mockRepository = &MockRepository{bikes: bikes}
	return
}

func (repo MockRepository) GetByFilter(filter Filter) (bikes []Bike, err error) {
	bikes = []Bike{}

	for _, bike := range repo.bikes {
		if filter.ID != "" && filter.ID != bike.ID {
			continue
		}

		if filter.Model != "" && filter.Model != bike.Model {
			continue
		}

		bikes = append(bikes, bike)
	}

	return
}

func (repo MockRepository) GetByID(bikeID string) (bike Bike, err error) {
	bikes, err := repo.GetByFilter(Filter{ID: bikeID})
	if err != nil {
		return
	}

	if len(bikes) > 1 {
		err = errors.New("database failure")
		return
	}

	if len(bikes) < 1 {
		err = ErrBikeNotFound
		return
	}

	bike = bikes[0]
	return
}

func (repo *MockRepository) Register(bike Bike) (bikeID string, err error) {
	bike.ID = uuid.NewV4().String()
	repo.bikes = append(repo.bikes, bike)

	bikeID = bike.ID
	return
}

func (repo *MockRepository) Update(bike Bike) (err error) {
	for i := range repo.bikes {
		if repo.bikes[i].ID == bike.ID {
			repo.bikes[i] = bike
			return
		}
	}

	err = ErrBikeNotFound
	return
}

func (repo *MockRepository) Delete(bikeID string) (err error) {
	for i := range repo.bikes {
		if repo.bikes[i].ID == bikeID {
			repo.bikes = append(repo.bikes[:i], repo.bikes[i+1:]...)
			return
		}
	}

	err = ErrBikeNotFound
	return
}

func (repo MockRepository) Start() (err error) {
	// todo: create bikes collection here
	return
}
