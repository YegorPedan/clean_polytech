package postgres

import (
	"clean-polytech/internal/domain/model"
	"clean-polytech/internal/infra/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type UserRepository struct {
	db *pgx.Conn
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
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	_, err = tx.Exec(ctx, "INSERT INTO user (id, name, family_name, phone_id) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.FamilyName, user.PhoneID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *UserRepository) GetAllUsers() ([]*model.User, error) {
	ctx := context.Background()
	rows, err := r.db.Query(ctx, "SELECT id, name, family_name, phone_id FROM user")
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
