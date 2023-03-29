package database

import (
	"database/sql"
	"log"

	"github.com/Kratos-28/gqlgen-todos/graph/model"
	_ "github.com/go-sql-driver/mysql"
)

var connectionString string = "root:1020@tcp(127.0.0.1:3306)/test"

type DB struct {
	client *sql.DB
}

func Connect() *DB {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: db,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
var jobListing model.JobListing
return &jobListing
}

