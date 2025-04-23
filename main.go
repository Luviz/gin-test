package main

import (
	"fmt"

	"github.com/luviz/gin-test/api"
	"github.com/luviz/gin-test/api/routes"
)

func main() {

	ginServer := api.InitGin()

	routes.HealthRoute(*ginServer.Group("/api"))
	routes.UserRoutes(*ginServer.Group("/api"))

	ginServer.Run(":5432")
	fmt.Println("hello world")
}
