package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	name, price string
}

func scraped() ([]Product, error) {

	var products []Product

	res, err := http.Get("https://www.scrapingcourse.com/ecommerce/")
	if err != nil {
		return products, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return products, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return products, err
	}

	doc.Find("li.product").Each(func(i int, p *goquery.Selection) {

		product := Product{} 

		product.name = p.Find("h2").Text()
		product.price = p.Find("span.price").Text()

		products = append(products, product)

	})

	return products, nil
}

func main() {

	result, err := scraped()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range result {

		fmt.Println(val)
	}
}
