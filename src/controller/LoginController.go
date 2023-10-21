package controller

import (
	"api/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostLogin(c *gin.Context) {
	var newLogin domain.Login
	if err := c.BindJSON(&newLogin); err != nil {
		panic(err.Error())
	}
	fmt.Println(newLogin)
}
