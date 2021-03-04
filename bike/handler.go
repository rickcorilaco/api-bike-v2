package bike

const (
	httpEndpoint = "bikes"
	httpID       = "bike_id"
)

type Handler interface {
	Start() (err error)
}
