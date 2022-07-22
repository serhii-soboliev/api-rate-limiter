package router

import (
	"github.com/gin-gonic/gin"
	"com.sbk/api-rate-limiter/api/v1/handler"
)


func InitalizeRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users/fixedwindow", handler.GetFixedWindowUsersAPI)
	return router
}


