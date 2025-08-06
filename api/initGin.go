package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luviz/gin-test/api/routes"
	s "github.com/luviz/gin-test/services"
)

func InitGin(s s.Services) *gin.Engine {
	log := s.Logger.WithName("gin")
	r := gin.Default()
	gin.DebugPrintFunc = func(format string, values ...interface{}) {
		log.Info("GIN-DEBUG", "values", values)
	}
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		msg := fmt.Sprintf("GIN-DEBUG > adding endpoint: %v %v (%v)",
			httpMethod, absolutePath, handlerName)
		log.Info(msg, "nuHandlers", nuHandlers)
	}
	api := r.Group("/api")
	routes.HealthRoute(*api)
	routes.UserRoutes(*api, s.User)

	return r
}
