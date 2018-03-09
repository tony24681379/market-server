package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/golang/glog"

	"github.com/PuerkitoBio/goquery"
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
	size := "&prod_size=&usort=&&p_data_num=72"
	doc, err := goquery.NewDocument("http://www.rt-mart.com.tw/direct/index.php?action=product_sort&prod_sort_uid=3920" + size)
	// index.php?action=product_sort&prod_sort_uid=14361&prod_size=&p_data_num=18&usort=auto_date%2CDESC&page=2
	if err != nil {
		log.Fatal(err)
	}

	s := doc.Find(".FOR_MAIN .main_content_center .classify_centerBox")
	category := s.Find(".classify_clist .t01").Text()

	total, err := strconv.Atoi(s.Find(".classify_clist .t02").Text())
	if err != nil {
		return
	}
	products := &Products{
		Category: category,
		Total:    total,
	}

	page := s.Find(".classify_numBox").First().Find("li").Length() - 2
	// if err != nil {
	// 	return
	// }
	glog.Info(page)

	addProduct(s, products)

	if page > 1 {
		for i := 2; i <= page; i++ {
			doc, err = goquery.NewDocument("http://www.rt-mart.com.tw/direct/index.php?action=product_sort&prod_sort_uid=3920" + size + "&page=" + strconv.Itoa(i))
			if err != nil {
				glog.Fatal(err)
			}
			s = doc.Find(".FOR_MAIN .main_content_center .classify_centerBox")
			addProduct(s, products)
		}
	}

	fmt.Printf("%s\n", products)
	glog.Info(len(products.Products))
}

func addProduct(s *goquery.Selection, products *Products) {
	s.Find("h5 a").Each(func(j int, h *goquery.Selection) {
		p := &Product{
			Name: h.Text(),
		}
		products.Products = append(products.Products, p)
	})

	s.Find(".for_pricebox").Each(func(j int, h *goquery.Selection) {
		products.Products[j].Price = h.Text()
	})
}
