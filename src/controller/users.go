package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"page-pal/service"
)

func UsersHandler(c *gin.Context) {
	userID := c.Query("user_id")

	if userID != "" {
		user, err := service.FetchUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, user)

	} else {
		users, err := service.AllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
