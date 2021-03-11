package bike

import (
	"net/http"

	"github.com/labstack/echo"
)

func (srvc Service) Echo(e *echo.Echo) (err error) {
	group := e.Group(colletionName)

	group.GET("", func(c echo.Context) (err error) {
		filter := Filter{}

		err = c.Bind(&filter)
		if err != nil {
			return
		}

		bikes, err := srvc.GetByFilter(filter)
		if err != nil {
			return
		}

		return c.JSON(200, bikes)
	})

	group.GET("/:"+parameterID, func(c echo.Context) (err error) {
		bikeID := c.Param(parameterID)
		if bikeID == "" {
			return c.JSON(http.StatusBadRequest, "invalid parameter")
		}

		bike, err := srvc.GetByID(bikeID)
		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, bike)
	})

	group.POST("", func(c echo.Context) (err error) {
		bike := Bike{}

		err = c.Bind(&bike)
		if err != nil {
			return
		}

		bikeID, err := srvc.Register(bike)
		if err != nil {
			return
		}

		return c.JSON(http.StatusCreated, echo.Map{"id": bikeID})
	})

	group.PUT("/:"+parameterID, func(c echo.Context) (err error) {
		bike := Bike{}

		err = c.Bind(&bike)
		if err != nil {
			return
		}

		bike.ID = c.Param(parameterID)
		if bike.ID == "" {
			return c.JSON(http.StatusBadRequest, "invalid parameter")
		}

		err = srvc.Update(bike)
		if err != nil {
			return
		}

		return c.NoContent(http.StatusNoContent)
	})

	group.DELETE("/:"+parameterID, func(c echo.Context) (err error) {
		bikeID := c.Param(parameterID)
		if bikeID == "" {
			return c.JSON(http.StatusBadRequest, "invalid parameter")
		}

		err = srvc.Delete(bikeID)
		if err != nil {
			return
		}

		return c.NoContent(http.StatusNoContent)
	})

	return
}
