package interfaces

import (
	"goravel/app/models"
)

type UserRepositoryInterface interface {
	All() ([]models.User, error)
	Find(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(id int, user models.User) (models.User, error)
	Delete(id int) error
}
