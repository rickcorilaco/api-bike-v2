package bike

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/rickcorilaco/api-bike-v2/connection"
	"github.com/rickcorilaco/api-bike-v2/controller"
)

type Service struct {
	repository Repository
}

func NewService(conn connection.Connection) (srvc *Service, err error) {
	repository, err := NewRepository(conn)
	if err != nil {
		return
	}

	srvc = &Service{repository: repository}
	return
}

func MustNewService(conn connection.Connection) (srvc *Service) {
	srvc, err := NewService(conn)
	if err != nil {
		panic(err)
	}

	return
}

func (srvc Service) Controller(ctrl controller.Controller) (err error) {
	switch framework := ctrl.Framework().(type) {
	case *echo.Echo:
		err = srvc.Echo(framework)
	default:
		err = errors.New("controller type is not implemented")
	}
	return
}

func (srvc Service) MustController(ctrl controller.Controller) {
	err := srvc.Controller(ctrl)
	if err != nil {
		panic(err)
	}

	return
}

func (srvc Service) GetByFilter(filter Filter) (bikes []Bike, err error) {
	return srvc.repository.GetByFilter(filter)
}

func (srvc Service) GetByID(bikeID string) (bike Bike, err error) {
	return srvc.repository.GetByID(bikeID)
}

func (srvc Service) Register(bike Bike) (bikeID string, err error) {
	return srvc.repository.Register(bike)
}

func (srvc Service) Update(bike Bike) (err error) {
	return srvc.repository.Update(bike)
}

func (srvc Service) Delete(bikeID string) (err error) {
	return srvc.repository.Delete(bikeID)
}
