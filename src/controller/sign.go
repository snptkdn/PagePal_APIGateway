package controller

import (
	"fmt"
	"net/http"
	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"page-pal/dto"
	"page-pal/model"
	"page-pal/service"
)

func SignUpHandler(c *gin.Context) {
  dto := dto.SignUpRequestDto{}
  if err := c.ShouldBindJSON(&dto); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s: sign up json is bad.", err)})
    c.Abort()
    return
  }

  isExist, err := service.ExistUser(dto.UserName)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: user select query is failed.", err)})
    c.Abort()
    return
  }

  if isExist {
    c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s: user is already exists.", err)})
    c.Abort()
    return
  }

  user := model.User{
    Name: dto.UserName,
    Pass: dto.Password,
  }
  
  if err := service.InsertUser(user); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: insert user is failed.", err)})
    c.Abort()
    return
  }

  c.JSON(http.StatusOK, user)
}

func SignInHandler(c *gin.Context) {
  dto := dto.SignUpRequestDto{}
  if err := c.ShouldBindJSON(&dto); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s: sign up json is bad.", err)})
    c.Abort()
    return
  }

  isExist, err := service.ExistUser(dto.UserName)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: user select query is failed.", err)})
    c.Abort()
    return
  }

  if !isExist {
    c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%v: user is not exists.", err)})
    c.Abort()
    return
  }

  isValid, err := service.IsValidUser(dto.UserName, dto.Password); 
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: insert user is failed.", err)})
    c.Abort()
    return
  }

  if isValid {
    c.String(http.StatusOK, fmt.Sprint(service.FindUserID(dto.UserName)))
  } else {
    c.Status(http.StatusBadRequest)
  }
}
