package postgres

import (
	"clean-polytech/internal/infra/config"
	"database/sql"
	"fmt"
)

func ConnectionPostgres(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgre.Host, cfg.Postgre.Port, cfg.Postgre.User, cfg.Postgre.Password, cfg.Postgre.DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS user (
    id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	family_name VARCHAR(100) NOT NULL`); err != nil {
		return nil, err
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS smartphone (
	id SERIAL PRIMARY KEY,
	model VARCHAR(100) NOT NULL,
	charge INT NOT NULL,
	connection_time TIMESTAMP NOT NULL,
	disconnect_time TIMESTAMP NOT NULL)`); err != nil {
		return nil, err
	}

	return db, nil
}
