package controller

import (
	"api/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var newUser domain.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	fmt.Println(newUser)
}
