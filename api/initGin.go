package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luviz/gin-test/api/routes"
	s "github.com/luviz/gin-test/services"
)

func InitGin(s s.Services) *gin.Engine {
	log := s.Logger.WithName("gin")
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	gin.DebugPrintFunc = func(format string, values ...any) {
		msg := fmt.Sprintf(format, values...)
		log.Info("GIN-DEBUG > " + msg)
	}
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		msg := fmt.Sprintf("GIN-DEBUG > adding endpoint: %v %v (%v)",
			httpMethod, absolutePath, handlerName)
		log.Info(msg, "nuHandlers", nuHandlers)
	}

	api := r.Group("/api")
	api.Use(func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		status := ctx.Writer.Status()

		log.Info("/api called",
			"method", ctx.Request.Method,
			"path", ctx.Request.URL.Path,
			"status", status,
			"ip", ctx.ClientIP(),
			"latency", latency,
		)
	})
	routes.HealthRoute(*api)
	routes.UserRoutes(*api, s)

	return r
}
