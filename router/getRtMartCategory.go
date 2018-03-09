package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tony24681379/market-server/rtMart"
)

func getRtMartCategory(r *rtMart.RTMart) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, r.GetCategory())
	}
}
