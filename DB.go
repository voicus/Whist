package main

import (
	"database/sql"
	"github.com/lib/pq"
)

var DB *sql.DB

func openDatabase() error {
	var err error
	pgURL, err := pq.ParseURL("postgres://bfpvzhah:cld_-IeKKFi7avbOPqoiKRu4V4GEMSx7@snuffleupagus.db.elephantsql.com/bfpvzhah")
	if err != nil {
		return err 
	}
	DB, err = sql.Open("postgres", pgURL)
	if err != nil {
		return err
	}
	return nil
}

func closeDatabase() error {
	return DB.Close()
}