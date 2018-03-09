package rtMart

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

type RTMart struct {
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

func NewRtMart() *RTMart {
	rtmark := &RTMart{
		Name: "RtMart",
		Url:  "http://www.rt-mart.com.tw/direct/",
	}
	rtmark.TopCategories = rtmark.getCategory()
	return rtmark
}

func (r *RTMart) GetCategory() []*TopCategory {
	return r.TopCategories
}

func (r *RTMart) getCategory() []*TopCategory {
	top := []*TopCategory{}
	doc, err := goquery.NewDocument(r.Url)
	if err != nil {
		log.Fatal(err)
	}
	top = append(top, r.getSubCategory(doc, top, ".rtmartHEADER02 .main_nav .nav01")...)
	top = append(top, r.getSubCategory(doc, top, ".rtmartHEADER02 .main_nav .nav02")...)
	return top
}

func (r *RTMart) getSubCategory(doc *goquery.Document, top []*TopCategory, selector string) []*TopCategory {
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
