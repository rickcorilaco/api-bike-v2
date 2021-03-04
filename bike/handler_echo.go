package bike

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type EchoHandler struct {
	e        *echo.Echo
	userCase *UserCase
}

func NewEchoHandler(e *echo.Echo, userCase *UserCase) (echoHandler EchoHandler, err error) {
	echoHandler = EchoHandler{
		e:        e,
		userCase: userCase,
	}
	return
}

func (h EchoHandler) GetByFilter(c echo.Context) (err error) {
	filter := Filter{}

	err = c.Bind(&filter)
	if err != nil {
		return
	}

	bikes, err := h.userCase.GetByFilter(filter)
	if err != nil {
		return
	}

	return c.JSON(200, bikes)
}

func (h EchoHandler) GetByID(c echo.Context) (err error) {
	bikeID, err := strconv.ParseInt(c.Param("bike_id"), 10, 64)
	if err != nil {
		return
	}

	bike, err := h.userCase.GetByID(bikeID)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, bike)
}

func (h EchoHandler) Register(c echo.Context) (err error) {
	bike := Bike{}

	err = c.Bind(&bike)
	if err != nil {
		return
	}

	bikeID, err := h.userCase.Register(bike)
	if err != nil {
		return
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": bikeID})
}

func (h EchoHandler) Update(c echo.Context) (err error) {
	bike := Bike{}

	err = c.Bind(&bike)
	if err != nil {
		return
	}

	bike.ID, err = strconv.ParseInt(c.Param("bike_id"), 10, 64)
	if err != nil {
		return
	}

	err = h.userCase.Update(bike)
	if err != nil {
		return
	}

	return c.NoContent(http.StatusNoContent)
}

func (h EchoHandler) Delete(c echo.Context) (err error) {
	bikeID, err := strconv.ParseInt(c.Param("bike_id"), 10, 64)
	if err != nil {
		return
	}

	err = h.userCase.Delete(bikeID)
	if err != nil {
		return
	}

	return c.NoContent(http.StatusNoContent)
}

func (h EchoHandler) Start() (err error) {
	bikes := h.e.Group("bikes")

	bikes.GET("", h.GetByFilter)
	bikes.GET("/:bike_id", h.GetByID)
	bikes.POST("", h.Register)
	bikes.PUT("/:bike_id", h.Update)
	bikes.DELETE("/:bike_id", h.Delete)
	return
}
