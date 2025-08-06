package main

import (
	"fmt"

	"github.com/luviz/gin-test/api"
	s "github.com/luviz/gin-test/services"
)

func main() {
	services := s.InitServices(true)
	ginServer := api.InitGin(services)

	ginServer.Run(":5432")
	fmt.Println("hello world")
}
