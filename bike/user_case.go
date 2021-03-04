package bike

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

type UserCase struct {
	repository Repository
}

func NewUserCase(repository Repository) (userCase *UserCase, err error) {
	userCase = &UserCase{repository: repository}
	return
}

func (u UserCase) GetByFilter(filter Filter) (bikes []Bike, err error) {
	return u.repository.GetByFilter(filter)
}

func (u UserCase) GetByID(bikeID int64) (bike Bike, err error) {
	return u.repository.GetByID(bikeID)
}

func (u UserCase) Register(bike Bike) (bikeID int64, err error) {
	return u.repository.Register(bike)
}

func (u UserCase) Update(bike Bike) (err error) {
	return u.repository.Update(bike)
}

func (u UserCase) Delete(bikeID int64) (err error) {
	return u.repository.Delete(bikeID)
}

func Start(database interface{}, handlers ...interface{}) (err error) {
	var respository Repository

	switch db := database.(type) {
	case *pg.DB:
		respository, err = NewPostgresORMRepository(db)
		if err != nil {
			return
		}

		err = respository.Start()
		if err != nil {
			return
		}

	default:
		err = errors.New("db is not implemented")
		return
	}

	userCase, err := NewUserCase(respository)
	if err != nil {
		return
	}

	if len(handlers) > 0 {
		switch h := handlers[0].(type) {
		case *echo.Echo:
			var handler EchoHandler

			handler, err = NewEchoHandler(h, userCase)
			if err != nil {
				return
			}

			err = handler.Start()
			if err != nil {
				return
			}

		default:
			err = errors.New("handler is not implemented")
			return
		}
	}

	return
}
