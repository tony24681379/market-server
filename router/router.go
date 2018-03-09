package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/tony24681379/market-server/rtMart"
)

func headerMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Token, token")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT,DELETE")
	c.Next()
}

// NewRouter create a router
func NewRouter(RTMart *rtMart.RTMart) http.Handler {
	r := gin.New()
	r.Use(ginglog.Logger(3 * time.Second))
	r.Use(gin.Recovery())

	r.Use(headerMiddleware)

	r.GET("/health", checkHealth)

	rtMartRoute := r.Group("/rt-mart")
	{
		rtMartRoute.GET("/category", getRtMartCategory(RTMart))
		rtMartRoute.GET("/product", getRtMartProduct(RTMart))
	}

	return r
}
