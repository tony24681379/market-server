package main

import (
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

func main() {
	doc, err := goquery.NewDocument("http://shopping.friday.tw/shopping/Browse.do?op=vc&cid=303131&sid=12&icn=%E6%9E%9C%E4%B9%BE&ici=%E7%BE%8E%E9%A3%9F%E5%90%8D%E7%94%A2%E4%B8%8B%E6%8B%89%E4%BA%8C%E5%8D%80%E7%9B%AE%E9%8C%842-5")
	if err != nil {
		glog.Fatal(err)
	}

	s := doc.Find(".main .product_block")
	page := []string{}
	s.Find(".page_number a").Each(func(j int, h *goquery.Selection) {
		totalString, _ := h.Attr("href")
		page = append(page, totalString)
	})
	if len(page) > 1 {
		page = page[:len(page)-1]
	}
	s.Find(".product_list h5").Each(func(j int, h *goquery.Selection) {
		a := h.Text()
		glog.Info(a)
	})
	s.Find(".product_list span.price").Each(func(j int, h *goquery.Selection) {
		a := h.Text()
		glog.Info(a)
	})
	// total, err := strconv.Atoi(s.Find("simplenum").Text())
	// if err != nil {
	// 	return
	// }
	// products := &Products{
	// 	Category: s.Find(".classify_clist .t01").Text(),
	// 	Total:    total,
	// }

	// page := s.Find(".classify_numBox").First().Find("li").Length() - 2
	// addProduct(s, products)
	// if page > 1 {
	// 	for i := 2; i <= page; i++ {
	// 		doc, err = goquery.NewDocument(url + category + size + "&page=" + strconv.Itoa(i))
	// 		if err != nil {
	// 			glog.Fatal(err)
	// 		}
	// 		s = doc.Find(".FOR_MAIN .main_content_center .classify_centerBox")
	// 		// addProduct(s, products)
	// 	}
	// }

	// glog.Info(products)
	// return products
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
