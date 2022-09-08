package data

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

 const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "web_service_gin"
)

var (
	Db  *sqlx.DB
	err error
)

// Init db
func Init() {
	args := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
	Db, err = sqlx.Open("postgres", args)
	if err != nil {
		panic(err)
	}
	// ping is necessary to create connection
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
}
