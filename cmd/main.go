package main

import (
	"com.sbk/api-rate-limiter/api/v1/router"
)

func main() {
	router := router.InitalizeRouter()

	router.Run("localhost:8081")
}
