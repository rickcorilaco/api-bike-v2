package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type EchoController struct {
	config Config
	echo   *echo.Echo
}

func NewEchoController(config Config) (ctrl Controller, err error) {
	e := echo.New()

	if config.Log {
		e.Use(middleware.Logger())
	}

	ctrl = EchoController{
		config: config,
		echo:   e,
	}
	return
}

func (ctrl EchoController) Framework() (framework interface{}) {
	return ctrl.echo
}

func (ctrl EchoController) Start() (err error) {
	return ctrl.echo.Start(":" + ctrl.config.Port)
}
