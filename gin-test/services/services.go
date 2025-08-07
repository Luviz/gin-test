package services

import (
	"github.com/go-logr/logr"
	"github.com/luviz/gin-test/services/user"
)

type Services struct {
	Logger logr.Logger
	User   user.UserService
}

func loadMockServices(log logr.Logger) Services {
	return Services{
		Logger: log.WithName("mock-services"),
		User:   user.InitMockUserService(),
	}
}

func loadServices(log logr.Logger) Services {
	return Services{
		Logger: log.WithName("services"),
	}
}

func InitServices(test bool) Services {
	log := NewLogger("text")
	if test {
		return loadMockServices(log)
	}
	return loadServices(log)
}
