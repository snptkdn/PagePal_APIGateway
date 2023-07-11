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

func BookHandler(c *gin.Context) {
  dto := dto.BookRequestDto{}
  if err := c.ShouldBindJSON(&dto); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s: book's data json is bad.", err)})
    c.Abort()
    return
  }

  book := model.Books {
    ISBN: dto.ISBN,
    Title: dto.Title,
    Author: dto.Author,
    ImageURL: dto.ImageURL,
    Description: dto.Description,
    PageCount: dto.PageCount,
  }
  
  err := service.InsertBook(book)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
    c.Abort()
    return
  }

  c.JSON(http.StatusOK, book)
}

