package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tony24681379/market-server/shopping"
)

func getShoppingCategory(s *shopping.Shopping) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, s.GetCategory())
	}
}
