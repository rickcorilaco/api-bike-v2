package bike

import (
	"errors"
	"fmt"
)

const (
	colletionName = "bikes"
	parameterID   = "bike_id"
)

var ErrBikeNotFound = errors.New("bike not found")

type Bike struct {
	ID    string `json:"id" pg:",pk"`
	Model string `json:"model"`
}

func (b Bike) String() string {
	return fmt.Sprintf("Bike: <ID: %s Model: %s>", b.ID, b.Model)
}

type Filter struct {
	ID    string `query:"id"`
	Model string `query:"model"`
}
