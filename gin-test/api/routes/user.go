package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luviz/gin-test/models"
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

	userGroup.POST("", createUser(s))
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

func createUser(s s.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := s.Logger.WithName("create")
		log.Info("creating object")
		err := s.User.Create(models.User{
			Id:   "u0012",
			Name: "bardia",
			Mail: "someName@bardiajedi.com",
		})
		if err != nil {
			log.Error(err, "could not create object")
			ctx.AbortWithError(http.StatusConflict, err)
		}
		ctx.JSON(http.StatusAccepted, gin.H{"status": "ok"})
	}
}
