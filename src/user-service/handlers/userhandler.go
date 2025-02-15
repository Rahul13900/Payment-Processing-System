package handlers

import (
	"log"
	"net/http"
	"user-service/middleware"
	"user-service/models"
	"user-service/store"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Store store.UserStore
}

func NewUserHandler(store store.UserStore) *UserHandler {
	return &UserHandler{Store: store}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	user.Password = hashPassword
	err = h.Store.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (h *UserHandler) SignIn(c *gin.Context) {
	log.Println("Signing in...")
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	// retrieve user
	existingUser, err := h.Store.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}
	if !utils.ComparePassword(existingUser.Password, []byte(user.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := middleware.GenerateJWT(existingUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
