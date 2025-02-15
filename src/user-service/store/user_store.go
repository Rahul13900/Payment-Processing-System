package store

import(
	"user-service/models"
)

type UserStore interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}


