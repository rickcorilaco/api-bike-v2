package main

import (
	"github.com/rickcorilaco/api-bike-v2/bike"
	"github.com/rickcorilaco/api-bike-v2/connection"
	"github.com/rickcorilaco/api-bike-v2/controller"

	"github.com/rickcorilaco/go/env"
)

func main() {
	env.MustFromFile("./env.json")

	conn := connection.MustNewFromEnv()

	ctrl := controller.MustNewFromEnv()

	bike.MustNewService(conn).MustController(ctrl)

	ctrl.Start()
}
