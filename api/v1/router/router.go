package router

import (
	"com.sbk/api-rate-limiter/api/v1/handler"
	"github.com/gin-gonic/gin"
)

func InitalizeRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users/fixedwindow", handler.GetUsersFW)
	router.GET("/users/slidingwindow", handler.GetUsersSW)
	router.GET("/resource/fixedwindow/:userId", handler.GetResourceFW)
	router.GET("/resource/slidingwindow/:userId", handler.GetResourceSW)
	return router
}
