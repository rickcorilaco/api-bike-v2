package main

import (
	"fmt"

	"github.com/rickcorilaco/api-bike-v2/bike"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rickcorilaco/go/env"
)

func main() {
	err := env.FromFile("env.json")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", env.MustString("database.host"), env.MustString("database.port")),
		User:     env.MustString("database.user"),
		Password: env.MustString("database.password"),
		Database: env.MustString("database.name"),
	})

	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	err = bike.Start(db, e)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":9000"))
}
