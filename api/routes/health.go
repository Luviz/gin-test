package routes

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func HealthRoute(group gin.RouterGroup) gin.RouterGroup {
	myUrl, _ := url.Parse("http://myhost.local/ep/")
	myUrl.Path += "/test/api"
	group.GET("health", func(ctx *gin.Context) {
		ctx.JSONP(202, gin.H{
			"status":  "OK",
			"baseUrl": myUrl.String(),
		})
	})

	return group
}
