package main

import (
	"fmt"

	"github.com/luviz/gin-test/api"
	"github.com/luviz/gin-test/api/routes"
	"github.com/luviz/gin-test/services/user"
)

func main() {

	ginServer := api.InitGin()

	routes.HealthRoute(*ginServer.Group("/api"))
	routes.UserRoutes(*ginServer.Group("/api"), user.InitMockUserService())

	ginServer.Run(":5432")
	fmt.Println("hello world")
}
