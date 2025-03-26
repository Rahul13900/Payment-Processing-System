package app

import (
	"fmt"
	"log"
	"payment-service/config"
	"payment-service/handlers"
	"payment-service/middleware"
	"payment-service/services"
	"payment-service/store"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server structure.
type Server struct {
	Router         *gin.Engine
	PaymentHandler *handlers.PaymentHandler
	WebhookHandler *handlers.WebhookHandler
	Port           string
}

// NewServer initializes a new HTTP server with routes and dependencies.
func NewServer(cfg *config.Config) (*Server, error) {
	// Connect to PostgreSQL
	db := middleware.CreateConnection(cfg.PostgresURL)
	if db == nil {
		return nil, logError("Failed to connect to database")
	}
	// Initialize NewPostgreSQL store
	store := store.NewPostgresStore(db)

	// Initialize Kafka Producer
	kafkaProducer := services.NewKafkaProducer(cfg.KafkaBroker) // No error return needed

	// Initialize Services
	paymentService := services.NewPaymentService(store, kafkaProducer, cfg.StripeSecret)

	// Initialize Handlers
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	webhookHandler := handlers.NewWebhookHandler(paymentService, cfg.StripeSecret)

	// Setup Router
	r := setupRouter(paymentHandler, webhookHandler)

	return &Server{
		Router:         r,
		PaymentHandler: paymentHandler,
		WebhookHandler: webhookHandler,
		Port:           cfg.ServerPort,
	}, nil
}

// setupRouter sets up the Gin router with middleware and routes.
func setupRouter(paymentHandler *handlers.PaymentHandler, webhookHandler *handlers.WebhookHandler) *gin.Engine {
	r := gin.Default()

	// Define routes
	r.POST("/payments", paymentHandler.InitiatePayment)           // Initiate a payment
	r.POST("/webhook/stripe", webhookHandler.HandleStripeWebhook) // Handle Stripe events

	return r
}

// Start launches the HTTP server.
func (s *Server) Start() error {
	log.Println("Payment Service running on port", s.Port)
	return s.Router.Run(":" + s.Port)
}

// logError is a helper function to log errors and return them.
func logError(format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	log.Println(err)
	return err
}
