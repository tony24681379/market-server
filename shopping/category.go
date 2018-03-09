package shopping

import (
	"log"
	"time"

	"github.com/golang/glog"

	"github.com/PuerkitoBio/goquery"
)

type Shopping struct {
	Name          string
	Url           string
	TopCategories []*TopCategory
}

type Category struct {
	Name string
	Url  string
}

type TopCategory struct {
	Name       string
	Url        string
	Categories []*Category
}

func NewShopping() *Shopping {
	shopping := &Shopping{
		Name: "Shopping",
		Url:  "http://shopping.friday.tw/shopping/1/s/12/",
	}
	ticker := time.NewTicker(5 * time.Minute)
	shopping.TopCategories = shopping.getCategory()
	go func() {
		for {
			select {
			case <-ticker.C:
				shopping.TopCategories = shopping.getCategory()
				glog.Info("update rt-mart category")
			}
		}
	}()

	return shopping
}

func (r *Shopping) GetCategory() []*TopCategory {
	return r.TopCategories
}

func (r *Shopping) getCategory() []*TopCategory {
	top := []*TopCategory{}
	doc, err := goquery.NewDocument("http://shopping.friday.tw/shopping/1/s/12/")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".navi-middle").Find(".amartnav").Not("ol").Each(func(j int, h *goquery.Selection) {
		firstHref, _ := h.Find("a").First().Attr("href")
		t := &TopCategory{
			Name: h.Find("a").First().Text(),
			Url:  firstHref,
		}
		glog.Info(h.Find("a").First().Text())
		h.Find(".amartblock a").Each(func(k int, l *goquery.Selection) {
			href, _ := l.Attr("href")
			c := &Category{
				Name: l.Text(),
				Url:  href,
			}
			t.Categories = append(t.Categories, c)
		})
		top = append(top, t)
	})
	return top
}

func (r *Shopping) getSubCategory(doc *goquery.Document, top []*TopCategory, selector string) []*TopCategory {
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		context := s.Find("a")
		firstHref, _ := context.First().Attr("href")
		topCategory := &TopCategory{
			Name: context.First().Text(),
			Url:  firstHref,
		}

		s.Not("h4").Find("a").Slice(1, s.Not("h4").Find("a").Size()).Each(func(j int, v *goquery.Selection) {
			href, _ := v.Attr("href")
			c := &Category{
				Name: v.Text(),
				Url:  href,
			}
			topCategory.Categories = append(topCategory.Categories, c)

		})
		top = append(top, topCategory)
	})
	return top
}
