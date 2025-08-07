package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/luviz/gin-test/services"
)

func UserRoutes(group gin.RouterGroup, s s.Services) gin.RouterGroup {
	userGroup := group.Group("users")
	userGroup.GET("", func(ctx *gin.Context) {
		ctx.JSONP(202, s.User.List())
	})
	userGroup.GET("/:id", func(ctx *gin.Context) {
		ctx.Set("id", ctx.Param("id"))
	}, getUserById(s))

	return group
}

func getUserById(s s.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := s.Logger.WithName("getUserById")
		id := ctx.GetString("id")
		log.Info("getting user", "id", id)
		if user, has := s.User.Get(id); has {
			ctx.JSON(http.StatusOK, user)
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		}
	}
}
