package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/todogo/configs"
)

func OpenConnection() (*sql.DB, error) {
	stringConnection := getDatabaseStringConnection()

	conn, err := sql.Open("postgres", stringConnection)
	if err != nil {
		//TODO remove panic 'cause should not use in production enviroment
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}

func getDatabaseStringConnection() string {
	conf := configs.GetDatabaseConfiguration()

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
}
