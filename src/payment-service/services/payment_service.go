package services

import (
	"errors"
	"log"
	"payment-service/models"
	"payment-service/store"

	"github.com/stripe/stripe-go/v72"
	stripePaymentIntent "github.com/stripe/stripe-go/v72/paymentintent"
)

// PaymentServiceInterface defines the contract for payment processing
type PaymentServiceInterface interface {
	ProcessPayment(req models.PaymentRequest) (*models.PaymentResponse, error)
	HandleWebhook(event models.StripeWebhookEvent) error
}


// paymentService implements PaymentServiceInterface
type paymentService struct {
	db       *store.PostgresStore
	producer KafkaProducerInterface
}

// Ensure paymentService implements PaymentServiceInterface
var _ PaymentServiceInterface = (*paymentService)(nil)

// NewPaymentService initializes a new PaymentService
func NewPaymentService(db *store.PostgresStore, producer KafkaProducerInterface) PaymentServiceInterface {
	return &paymentService{
		db:       db,
		producer: producer,
	}
}

// ProcessPayment handles payment processing and Kafka event publishing
func (s *paymentService) ProcessPayment(req models.PaymentRequest) (*models.PaymentResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("invalid payment amount")
	}
	stripe.Key = "sk_test_51QuW19RxJOEMEci0mHGC0BWQh69FFVCKlgedX0PeAItEebKH8gjzAEgrWr0A8fQMfOfgi7WA0nRHTfDzeEBHk3Ni00nhviCJaE"
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(req.Amount),
		Currency:           stripe.String(req.Currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		PaymentMethod: stripe.String("pm_card_visa"),
    Confirm:       stripe.Bool(true),
	}

	paymentIntent, err := stripePaymentIntent.New(params)
	if err != nil {
		log.Printf("Failed to create payment intent: %v", err)
		return nil, err
	}

	payment := models.Payment{
		ID:       paymentIntent.ID,
		Amount:   req.Amount,
		Currency: req.Currency,
		Status:   string(paymentIntent.Status),
	}

	if err := s.db.SavePayment(payment); err != nil {
		log.Printf("Failed to save payment record: %v", err)
		return nil, err
	}

	event := models.PaymentEvent{
		PaymentID: payment.ID,
		Status:    payment.Status,
	}

	if err := s.producer.Publish("payments", event); err != nil {
		log.Printf("Failed to publish payment event: %v", err)
		return nil, err
	}

	return &models.PaymentResponse{
		PaymentID: paymentIntent.ID,
		Status:    string(paymentIntent.Status),
	}, nil
}

// HandleWebhook processes Stripe webhook events
func (s *paymentService) HandleWebhook(event models.StripeWebhookEvent) error {
	// Business logic for handling different Stripe webhook events
	if event.Type == "payment_intent.succeeded" {
		payment := models.Payment{
			ID:     event.PaymentID,
			Status: event.Status,
		}

		if err := s.db.UpdatePaymentStatus(payment.ID, payment.Status); err != nil {
			log.Printf("Failed to update payment status: %v", err)
			return err
		}

		// Publish event to Kafka
		kafkaEvent := models.PaymentEvent{
			PaymentID: event.PaymentID,
			Status:    event.Status,
		}
		return s.producer.Publish("payments", kafkaEvent)
	}

	return nil
}

// // **Stripe Integration Notes:**
// // - Using Stripe official Go SDK (github.com/stripe/stripe-go/v72), which is the standard approach.
// // - PaymentIntent is Stripe's recommended workflow for handling payments.
// // - Secure handling of Stripe API keys is critical.

// // **Next Steps:**
// // - Add handler to route requests.
// // - Develop database SavePayment method.
// // - Write unit and integration tests.
