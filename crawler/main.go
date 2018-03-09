package main

import (
	"fmt"
	"log"
	"strconv"

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
	doc, err := goquery.NewDocument("http://www.rt-mart.com.tw/direct/index.php?action=product_sort&prod_sort_uid=24347" + size)
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

	// s.Find("h5 a").Each(func(j int, h *goquery.Selection) {
	// 	p := &Product{
	// 		Name: h.Text(),
	// 	}
	// 	products.Products = append(products.Products, p)
	// })

	// s.Find(".for_pricebox").Each(func(j int, h *goquery.Selection) {
	// 	products.Products[j].Price = h.Text()
	// })
	addProduct(s, products)
	fmt.Printf("%s\n", products)

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
