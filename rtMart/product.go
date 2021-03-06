package rtMart

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

type Products struct {
	Category string
	Total    int
	Products []*Product
}

type Product struct {
	Name  string
	Price string
}

func (r *RTMart) GetProduct(url, category string) *Products {
	size := "&prod_size=&usort=&&p_data_num=72"
	doc, err := goquery.NewDocument(url + category + size)
	if err != nil {
		glog.Fatal(err)
	}

	s := doc.Find(".FOR_MAIN .main_content_center .classify_centerBox")
	total, err := strconv.Atoi(s.Find(".classify_clist .t02").Text())
	if err != nil {
		return nil
	}
	products := &Products{
		Category: s.Find(".classify_clist .t01").Text(),
		Total:    total,
	}

	page := s.Find(".classify_numBox").First().Find("li").Length() - 2
	addProduct(s, products)
	if page > 1 {
		for i := 2; i <= page; i++ {
			doc, err = goquery.NewDocument(url + category + size + "&page=" + strconv.Itoa(i))
			if err != nil {
				glog.Fatal(err)
			}
			s = doc.Find(".FOR_MAIN .main_content_center .classify_centerBox")
			addProduct(s, products)
		}
	}

	glog.Info(products)
	return products
}

func addProduct(s *goquery.Selection, products *Products) {
	productsNew := []*Product{}
	s.Find("h5 a").Each(func(j int, h *goquery.Selection) {
		p := &Product{
			Name: strings.TrimSpace(h.Text()),
		}
		productsNew = append(productsNew, p)
	})

	s.Find(".for_pricebox").Each(func(j int, h *goquery.Selection) {
		productsNew[j].Price = h.Text()
	})
	products.Products = append(products.Products, productsNew...)
}
