package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

type Category struct {
	Name string
	Url  string
}

type TopCategory struct {
	Name       string
	Url        string
	Categories []*Category
}

func main() {
	doc, err := goquery.NewDocument("http://shopping.friday.tw/shopping/1/s/12/")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".navi-middle").Find(".amartnav").Not("ol").Each(func(j int, h *goquery.Selection) {
		firstHref, _ := h.First().Attr("href")
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
			// glog.Info(k, l.Text())
		})
	})
	// category := s.Find("a").Not("ol.lightbox").Text()
	// glog.Info(category)

}
