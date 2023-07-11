package controller

import (
	"net/http"
	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"page-pal/migrate"
)

func MigrateHandler(c *gin.Context) {
  migrate.Migrate()

  c.Status(http.StatusOK)
}

