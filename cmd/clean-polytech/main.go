package main

import (
	"clean-polytech/internal/infra/config"
	"clean-polytech/internal/infra/db/postgres"
	"fmt"
	"log"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.ConnectionPostgres(cfg)
	if err != nil {
		log.Fatal("Error connect postgres", err.Error())
	}

	fmt.Println(db)
	//fmt.Println(cfg)
}
