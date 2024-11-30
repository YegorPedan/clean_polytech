package postgres

import (
	"clean-polytech/internal/domain/model"
	"clean-polytech/internal/infra/config"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func ConnectionPostgres(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgre.Host, cfg.Postgre.Port, cfg.Postgre.User, cfg.Postgre.Password, cfg.Postgre.Database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS "user" (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		family_name VARCHAR(100) NOT NULL)`); err != nil {
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
func (r *UserRepository) SaveUser(ctx context.Context, user *model.User) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // Rollback if there's a panic
			panic(p)
		} else if err != nil {
			_ = tx.Rollback() // Rollback if an error occurred
		}
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO user (id, name, family_name, phone_id) VALUES ($1, $2, $3, $4)",
		user.ID, user.Name, user.FamilyName, user.PhoneID)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *UserRepository) GetAllUsers() ([]*model.User, error) {
	rows, err := r.db.Query("SELECT id, name, family_name, phone_id FROM user")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.FamilyName, &user.PhoneID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
