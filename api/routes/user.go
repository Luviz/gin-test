package routes

import "github.com/gin-gonic/gin"

func UserRoutes(group gin.RouterGroup) gin.RouterGroup {
	group.GET("user", func(ctx *gin.Context) {
		ctx.JSONP(202, gin.H{
			"id": "some userid",
		})
	})

	return group
}
