package main

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//user endpoint
	router.POST("/user", controller.PostUser)

	//login endpoint
	router.POST("/login", controller.PostLogin)
	router.Run(":8080")
}
