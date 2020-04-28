package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB(databaseUrl string) (err error) {
	Db, err = sqlx.Connect("mysql", databaseUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = Db.Ping()
	return
}
