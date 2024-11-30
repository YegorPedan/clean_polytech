package main

import (
	"clean-polytech/internal/infra/config"
	"clean-polytech/internal/infra/db/postgres"
	userHttp "clean-polytech/internal/transport/http"
	"database/sql"
	"log"
	"net/http"
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

	userRepo := postgres.NewUserRepository(db)
	phoneRepo := postgres.NewPhoneRepository(db)
	saveDataHandler := userHttp.NewUserHandler(userRepo, phoneRepo)
	http.Handle("/save_user", saveDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
