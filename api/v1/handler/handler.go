package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	fw "com.sbk/api-rate-limiter/pkg/fixedwindow"
)

func GetFixedWindowUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fw.GetFixedWindowUsers())
}

func GetResource(c *gin.Context) {
	userId := c.Param("userId")
	result := fw.GetResource(userId)
	if result {
		c.IndentedJSON(http.StatusOK, "Request is allowed for user "+userId)
	} else {
		msg := fmt.Sprintf("Too much reqeuests for %s", userId)
		c.IndentedJSON(http.StatusForbidden, msg)
	}
}
