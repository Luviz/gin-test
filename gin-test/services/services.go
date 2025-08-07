package services

import (
	"github.com/go-logr/logr"
	"github.com/luviz/gin-test/datasources"
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
		User:   tryLoadUser(log),
	}
}

func tryLoadUser(logger logr.Logger) user.UserService {
	log := logger.WithName("user-service-init")
	userService, err := user.InitPGUser(user.PGUserConfig{
		PG: datasources.PostgresDsnConfig{
			Host:     "database",
			User:     "postgres",
			Password: "qwert1234",
			Database: "test-app",
			Port:     "5432",
		},
	}, logger)
	if err != nil {
		log.Error(err, "could not load user service via postgres")
		return user.InitMockUserService()
	}

	return userService
}

func InitServices(test bool) Services {
	log := NewLogger("text")
	if test {
		return loadMockServices(log)
	}
	return loadServices(log)
}
