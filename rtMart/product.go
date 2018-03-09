package rtMart

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

func (r *RTMart) GetProduct(url, category string) *Products {
	size := "&prod_size=&usort=&&p_data_num=72"
	doc, err := goquery.NewDocument(url + category + size)
	if err != nil {
		log.Fatal(err)
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

	addProduct(s, products)

	fmt.Printf("%s\n", products)
	return products
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
