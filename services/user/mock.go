package user

import "github.com/luviz/gin-test/models"

type MockUserService struct {
}

func InitMockUserService() MockUserService {
	return MockUserService{}
}

func (MockUserService) Get(id string) models.User {
	return models.User{
		Id:   "some id",
		Name: "Bardia Jedi",
		Mail: "me@bardiajedi.com",
	}
}

func (MockUserService) List() []models.User {
	return []models.User{
		{
			Id:   "some id",
			Name: "Bardia Jedi",
			Mail: "me@bardiajedi.com",
		},
		{
			Id:   "some other id",
			Name: "Meh Peh",
			Mail: "meh.peh@bardiajedi.com",
		}}

}
