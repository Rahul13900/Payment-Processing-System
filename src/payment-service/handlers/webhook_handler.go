package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"payment-service/models"
	"payment-service/services"

	"github.com/gin-gonic/gin"
	// "github.com/stripe/stripe-go/webhook"
)

// WebhookHandler handles incoming webhook events from Stripe
type WebhookHandler struct {
	PaymentService services.PaymentServiceInterface
	StripeSecret   string
}

// NewWebhookHandler initializes a new WebhookHandler
func NewWebhookHandler(paymentService services.PaymentServiceInterface, stripeSecret string) *WebhookHandler {
	return &WebhookHandler{
		PaymentService: paymentService,
		StripeSecret:   stripeSecret,
	}
}

// HandleStripeWebhook processes incoming webhook events from Stripe
func (wh *WebhookHandler) HandleStripeWebhook(c *gin.Context) {
	log.Printf("HandleStripeWebhook Triggered")
	const MaxBodyBytes = int64(65536) // 64KB limit for security
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Request body too large"})
		return
	}

	// // Verify Stripe signature (optional, add security)
	// sigHeader := c.GetHeader("Stripe-Signature")
	// _ , err = webhook.ConstructEvent(payload, sigHeader, wh.StripeSecret)
	// if err != nil {
	// 	log.Printf("Stripe signature verification failed: %v", err)
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
	// 	return
	// }

	var stripeEvent models.StripeWebhookEvent
	err = json.Unmarshal(payload, &stripeEvent)
	if err != nil {
		log.Printf("Failed to parse webhook event: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event data"})
		return
	}

	log.Printf("Received Stripe event: %v", stripeEvent.Type)

	// Handle the webhook event in the service
	err = wh.PaymentService.HandleWebhook(stripeEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process event"})
		return
	}

	// No need to access `KafkaProducer` from handler, service takes care of it
	c.JSON(http.StatusOK, gin.H{"message": "Event processed"})
}
