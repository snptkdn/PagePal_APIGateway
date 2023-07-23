package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"page-pal/service"
)

func UsersHandler(c *gin.Context) {
	users, err := service.AllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, users)
}

