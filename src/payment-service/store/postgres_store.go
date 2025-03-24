package store

import (
	"database/sql"
	"log"
	"payment-service/models"
)

type PostgresStore struct {
	DB *sql.DB
}

// NewDatabase initializes the PostgreSQL connection
func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{DB: db}
}

// SavePayment inserts a new payment record into the database
func (ps *PostgresStore) SavePayment(payment models.Payment) error {
	query := "INSERT INTO payments (id, amount, currency, status) VALUES ($1, $2, $3, $4)"
	_, err := ps.DB.Exec(query, payment.ID, payment.Amount, payment.Currency, payment.Status)
	if err != nil {
		log.Printf("Error saving payment: %v", err)
		return err
	}
	return nil
}

// GetPaymentByID retrieves a payment record by ID
func (ps *PostgresStore) GetPaymentByID(paymentID string) (*models.Payment, error) {
	query := "SELECT id, amount, currency, status FROM payments WHERE id=$1"
	row := ps.DB.QueryRow(query, paymentID)

	var payment models.Payment
	if err := row.Scan(&payment.ID, &payment.Amount, &payment.Currency, &payment.Status); err != nil {
		log.Printf("Error retrieving payment: %v", err)
		return nil, err
	}
	return &payment, nil
}

// UpdatePaymentStatus updates the status of a payment in the database
func (ps *PostgresStore) UpdatePaymentStatus(paymentID string, status string) error {
	query := "UPDATE payments SET status=$1 WHERE id=$2"
	_, err := ps.DB.Exec(query, status, paymentID)
	if err != nil {
		log.Printf("Error updating payment status: %v", err)
		return err
	}
	return nil
}
