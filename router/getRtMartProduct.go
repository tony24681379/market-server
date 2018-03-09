package router

import (
	"net/http"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"
	"github.com/tony24681379/market-server/rtMart"
)

func getRtMartProduct(r *rtMart.RTMart) gin.HandlerFunc {
	return func(g *gin.Context) {
		c := g.Query("category")
		if c == "" {
			g.Status(404)
			return
		}
		for _, t := range r.TopCategories {
			for _, category := range t.Categories {
				if category.Name == c {
					p := r.GetProduct(r.Url, category.Url)
					glog.Info(p)
					g.JSON(http.StatusOK, p)
					return
				}
			}
		}
		g.Status(404)
	}
}
