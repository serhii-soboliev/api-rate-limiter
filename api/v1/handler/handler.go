package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	fw "com.sbk/api-rate-limiter/pkg/fixedwindow"
	sw "com.sbk/api-rate-limiter/pkg/slidingwindow"
)

func GetUsersFW(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fw.GetFixedWindowUsers())
}

func GetResourceFW(c *gin.Context) {
	userId := c.Param("userId")
	result := fw.GetResource(userId)
	if result {
		c.IndentedJSON(http.StatusOK, "Request is allowed for user "+userId)
	} else {
		msg := fmt.Sprintf("Too much reqeuests for %s", userId)
		c.IndentedJSON(http.StatusForbidden, msg)
	}
}

func GetUsersSW(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sw.GetSlidingWindowUsers())
}

func GetResourceSW(c *gin.Context) {
	userId := c.Param("userId")
	result := sw.GetResource(userId)
	if result {
		c.IndentedJSON(http.StatusOK, "Request is allowed for user "+userId)
	} else {
		msg := fmt.Sprintf("Too much reqeuests for %s", userId)
		c.IndentedJSON(http.StatusForbidden, msg)
	}
}
