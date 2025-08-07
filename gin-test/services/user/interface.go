package user

import "github.com/luviz/gin-test/models"

type UserService interface {
	List() []models.User
	Get(id string) (models.User, bool)
	Create(user models.User) error
}
