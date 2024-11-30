package postgres

import "database/sql"

type PhoneRepository struct {
	db *sql.DB
}

func NewPhoneRepository(db *sql.DB) *PhoneRepository {
	return &PhoneRepository{db: db}
}
