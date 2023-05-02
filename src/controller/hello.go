package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

  "page-pal/dto"
)

func HelloHandler(c *gin.Context) {
  dto := dto.HelloResponseDto {
    Text: "hello",
  }
  
	c.JSON(200, dto)
}
