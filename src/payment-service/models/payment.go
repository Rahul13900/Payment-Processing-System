package models

type PaymentRequest struct {
	UserID string `json:"user_id"`
	Amount int64  `json:"amount"`
	Method string `json:"method"`
}

type PaymentResponse struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}
