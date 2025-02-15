package store

import (
	"database/sql"
	"fmt"
	"user-service/models"
)

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{DB: db}
}

func (ps *PostgresStore) CreateUser(user *models.User) error {
	// if the user already exists
	existingUser, err := ps.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	// if not exists then create new user
	sqlStatement := "INSERT INTO users(name,email,password) VALUES ($1,$2,$3)"
	_, err = ps.DB.Exec(sqlStatement,user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (ps *PostgresStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	sqlStatement := "SELECT * FROM users where email = $1"
	err := ps.DB.QueryRow(sqlStatement,email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
