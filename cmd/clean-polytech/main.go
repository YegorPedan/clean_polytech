package main

import (
	"clean-polytech/internal/infra/config"
	"clean-polytech/internal/infra/db/postgres"
	"database/sql"
	"log"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.ConnectionPostgres(cfg)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	if err != nil {
		log.Fatal("Error connect postgres", err.Error())
	}

}
