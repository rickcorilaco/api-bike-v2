package bike

import (
	"fmt"
)

type Bike struct {
	ID    int64  `json:"id" pg:",pk"`
	Model string `json:"model"`
}

func (b Bike) String() string {
	return fmt.Sprintf("Bike: <ID: %d Model: %s>", b.ID, b.Model)
}

type Filter struct {
	ID    int64  `query:"id"`
	Model string `query:"model"`
}
