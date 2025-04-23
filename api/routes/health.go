package routes

import "github.com/gin-gonic/gin"

func HealthRoute(group gin.RouterGroup) gin.RouterGroup {
	group.GET("health", func(ctx *gin.Context) {
		ctx.JSONP(202, gin.H{
			"status": "OK",
		})
	})

	return group
}
