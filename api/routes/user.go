package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luviz/gin-test/services/user"
)

func UserRoutes(group gin.RouterGroup, userService user.UserService) gin.RouterGroup {

	group.GET("users/", func(ctx *gin.Context) {
		ctx.JSONP(202, userService.List())
	})
	group.GET("users/:id", func(ctx *gin.Context) {
		ctx.JSONP(202, userService.Get(""))
	})

	return group
}
