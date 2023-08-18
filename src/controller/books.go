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

	book := &model.Book{
		ISBN:        dto.ISBN,
		Title:       dto.Title,
		Author:      dto.Author,
		ImageURL:    dto.ImageURL,
		Description: dto.Description,
		PageCount:   dto.PageCount,
	}

	err := service.InsertBook(*book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
		c.Abort()
		return
	}
  
	book, err = service.FindBook(book.ISBN)
	if err != nil {
		if book == nil {
			c.Status(http.StatusNotFound)
			c.Abort()
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetBookHandler(c *gin.Context) {
	isbn := c.Query("isbn")

	book, err := service.FindBook(isbn)
	if err != nil {
		if book == nil {
			c.Status(http.StatusNotFound)
			c.Abort()
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, book)
}

func ReadHistoryHandler(c *gin.Context) {
	dto := dto.ReadHistoryRequestDto{}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("%s: read request data json is bad.", err)})
		c.Abort()
		return
	}

	history := model.ReadHistory{
		BookID: dto.BookID,
		UserID: dto.UserID,
		IsRead: dto.IsRead,
		Rate:   dto.Rate,
		Date:   dto.Date,
	}

	err := service.InsertReadHistory(history)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: history insert is failed.", err)})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, history)
}

func GetReadHistoryHandler(c *gin.Context) {
	user_id := c.Query("user_id")

	histories, err := service.FindHistory(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s: book insert is failed.", err)})
		c.Abort()
		return
	}

	for _, history := range histories {
		fmt.Println(history.Book)
	}

	c.JSON(http.StatusOK, histories)
}
