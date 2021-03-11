package controller

import (
	"errors"

	"github.com/rickcorilaco/go/env"
)

const (
	EchoKind = "echo"
)

type Controller interface {
	Framework() (framework interface{})
	Start() (err error)
}

type Config struct {
	Kind string
	Port string
	Log  bool
}

func New(config Config) (ctrl Controller, err error) {
	switch config.Kind {
	case EchoKind:
		ctrl, err = NewEchoController(config)
	default:
		err = errors.New("invalid kind of controller")
	}

	return
}

func MustNew(config Config) (ctrl Controller) {
	ctrl, err := New(config)
	if err != nil {
		panic(err)
	}

	return
}

func NewFromEnv() (ctrl Controller, err error) {
	config := Config{
		Kind: env.MustString("controller.kind"),
		Port: env.TryString("controller.port"),
		Log:  env.TryBool("controller.log"),
	}

	return New(config)
}

func MustNewFromEnv() (ctrl Controller) {
	ctrl, err := NewFromEnv()
	if err != nil {
		panic(err)
	}

	return
}
