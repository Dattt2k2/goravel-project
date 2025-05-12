package interfaces

import (
	"goravel/app/models"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUser(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id int, user models.User) (models.User, error)
	DeleteUser(id int) error
}
