package db

import (
	"database/sql"
	"fmt"

	"github.com/kuma-coffee/go-clean-archi/config"
)

func CreateConn() *sql.DB {
	conf := config.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.HOST, conf.PORT, conf.USER_NAME, conf.PASSWORD, conf.DB_NAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
