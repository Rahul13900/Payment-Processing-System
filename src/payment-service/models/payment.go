package models

// Payment represents the payment record stored in the database
type Payment struct {
	ID       string `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Status   string `json:"status"`
}

// PaymentRequest represents the incoming payment request from API
type PaymentRequest struct {
	Amount          int64  `json:"amount"`
	Currency        string `json:"currency"`
	PaymentMethodID string `json:"payment_method_id"`
}

// PaymentResponse represents the response sent back to the client
type PaymentResponse struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}

// PaymentEvent represents the payment event published to Kafka
type PaymentEvent struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}

// StripeWebhookEvent represents the structure of webhook events received from Stripe
type StripeWebhookEvent struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
}
