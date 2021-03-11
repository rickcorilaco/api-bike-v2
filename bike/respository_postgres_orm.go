package bike

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type PostgresOrmRepository struct {
	db *pg.DB
}

func NewPostgresORMRepository(db *pg.DB) (postgresOrmRepository PostgresOrmRepository, err error) {
	postgresOrmRepository = PostgresOrmRepository{db: db}
	return
}

func (repo PostgresOrmRepository) GetByFilter(filter Filter) (bikes []Bike, err error) {
	bikes = []Bike{}
	query := repo.db.Model(&bikes)

	if filter.ID != "" {
		query.Where("id = ?", filter.ID)
	}

	if filter.Model != "" {
		query.Where("model = ?", filter.Model)
	}

	err = query.Select()
	return
}

func (repo PostgresOrmRepository) GetByID(bikeID string) (bike Bike, err error) {
	bike = Bike{}
	err = repo.db.Model(&bike).Where("id = ?", bikeID).Select()
	return
}

func (repo PostgresOrmRepository) Register(bike Bike) (bikeID string, err error) {
	_, err = repo.db.Model(&bike).Insert()
	bikeID = bike.ID
	return
}

func (repo PostgresOrmRepository) Update(bike Bike) (err error) {
	_, err = repo.db.Model(&bike).WherePK().Update()
	return
}

func (repo PostgresOrmRepository) Delete(bikeID string) (err error) {
	_, err = repo.db.Model(&Bike{ID: bikeID}).WherePK().Delete()
	return
}

func (repo PostgresOrmRepository) Start() (err error) {
	return repo.db.Model((*Bike)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}
