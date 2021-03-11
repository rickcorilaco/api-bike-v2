package connection

import (
	"fmt"

	"github.com/go-pg/pg"
)

type PostegresORMConnection struct {
	db *pg.DB
}

func NewPostgresORMConnection(config Config) (conn Connection, err error) {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		User:     config.Username,
		Password: config.Password,
		Database: config.Name,
	})

	conn = PostegresORMConnection{db: db}
	return
}

func (conn PostegresORMConnection) Interface() (i interface{}) {
	return conn.db
}

func (conn PostegresORMConnection) Close() (err error) {
	return conn.db.Close()
}
