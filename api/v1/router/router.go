package router

import (
	"com.sbk/api-rate-limiter/api/v1/handler"
	"github.com/gin-gonic/gin"
)

func InitalizeRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users/fixedwindow", handler.GetFixedWindowUsers)
	router.GET("/resource/fixedwindow/:userId", handler.GetResource)
	return router
}
