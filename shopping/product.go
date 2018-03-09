package shopping

import (
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

func (r *Shopping) GetProduct(category, url string) *Products {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		glog.Fatal(err)
	}
	products := &Products{
		Category: category,
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

	addProduct(s, products)
	if len(page) > 0 {
		for i := range page {
			doc, err = goquery.NewDocument("http://shopping.friday.tw" + page[i])
			if err != nil {
				glog.Error(err)
			}
			s = doc.Find(".main .product_block")
			addProduct(s, products)
		}
	}

	products.Total = len(products.Products)
	return products
}

func addProduct(s *goquery.Selection, products *Products) {
	productsNew := []*Product{}

	s.Find(".product_list h5").Each(func(j int, h *goquery.Selection) {
		p := &Product{
			Name: h.Text(),
		}
		productsNew = append(productsNew, p)
	})

	s.Find(".product_list span.price").Each(func(j int, h *goquery.Selection) {
		productsNew[j].Price = h.Text()
	})
	products.Products = append(products.Products, productsNew...)
}
