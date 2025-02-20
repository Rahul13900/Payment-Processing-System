package handlers

import (
	"net/http"
	"payment-service/models"
	"payment-service/services"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service services.PaymentServiceInterface
}

func NewPaymentHandler(service services.PaymentServiceInterface) *PaymentHandler {
	return &PaymentHandler{service: service}

}

func (h *PaymentHandler) InitiatePayment(c *gin.Context) {
	var req models.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	response, err := h.service.ProcessPayment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
