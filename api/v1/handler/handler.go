package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	fw "com.sbk/api-rate-limiter/pkg/fixedwindow"
)

func GetFixedWindowUsersAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fw.GetFixedWindowUsers())
}
