package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
