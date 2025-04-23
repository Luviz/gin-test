package api

import "github.com/gin-gonic/gin"

func InitGin() *gin.Engine {
	r := gin.Default()

	return r
}
