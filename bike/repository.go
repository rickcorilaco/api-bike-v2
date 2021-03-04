package bike

type Repository interface {
	Start() (err error)
	GetByFilter(filte Filter) (bikes []Bike, err error)
	GetByID(bikeID int64) (bike Bike, err error)
	Register(bike Bike) (bikeID int64, err error)
	Update(bike Bike) (err error)
	Delete(bikeID int64) (err error)
}
