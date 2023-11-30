package database

import (
	"database/sql"
	"log"

	util "boonsboos.nl/boonsu/Util"
	_ "github.com/lib/pq"
)

var BoonsuDB Database = NewDatabase()

type Database struct {
	db *sql.DB
}

func NewDatabase() Database {
	db_inst, err := sql.Open("postgres", util.Options.DatabaseURL)
	if err != nil {
		log.Fatal("Could not connect to database! " + err.Error())
	}
	return Database{db_inst}
}
