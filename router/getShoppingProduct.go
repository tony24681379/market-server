package router

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/tony24681379/market-server/shopping"

	"github.com/gin-gonic/gin"
)

func getShoppingProduct(s *shopping.Shopping) gin.HandlerFunc {
	return func(g *gin.Context) {
		c := g.Query("category")
		if c == "" {
			g.Status(404)
			return
		}
		for _, t := range s.TopCategories {
			for _, category := range t.Categories {
				if category.Name == c {
					p := s.GetProduct(category.Name, category.Url)
					glog.Info(p)
					g.JSON(http.StatusOK, p)
					return
				}
			}
		}
		g.Status(404)
	}
}
