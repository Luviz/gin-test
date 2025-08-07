package user

import (
	"maps"
	"slices"

	"github.com/luviz/gin-test/models"
)

var MockUsers = map[string]models.User{
	"u0001": {
		Id:   "u0001",
		Name: "Bardia Jedi",
		Mail: "me@bardiajedi.com",
	},
	"u0002": {
		Id:   "u0002",
		Name: "Meh Peh",
		Mail: "meh.peh@bardiajedi.com",
	},
	"u0003": {
		Id:   "u0003",
		Name: "John Duo",
		Mail: "john.duo@bardiajedi.com",
	},
}

type MockUserService struct {
}

func InitMockUserService() MockUserService {
	return MockUserService{}
}

func (MockUserService) Get(id string) (models.User, bool) {
	user, has := MockUsers[id]
	return user, has
}

func (MockUserService) List() []models.User {
	return slices.Collect(maps.Values(MockUsers))
}
