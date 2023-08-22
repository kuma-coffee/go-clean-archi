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
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
