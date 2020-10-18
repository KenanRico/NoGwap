package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"errors"
	"strconv"
)

func AmazonCrawler(search_phrase string) ([]OnSaleItem, error){

	var on_sale_items []OnSaleItem

	// Get request search results (currently only get 1st page)
	search_phrase = strings.ReplaceAll(search_phrase, " ", "+")
	url := "https://www.amazon.ca/s?k=" + search_phrase
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return on_sale_items, errors.New("Amazon: "+err.Error())
	}
	//request.Header.Set("", "")
	resp, err := new(http.Client).Do(request)
	if err != nil {
		return on_sale_items, errors.New("Amazon: "+err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return on_sale_items, errors.New("Amazon: "+resp.Status)
	}

	// Create goquery.Document for search
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return on_sale_items, errors.New("Amazon: "+resp.Status)
	}

	doc.Find("div").Each(
		func(i int, sel *goquery.Selection){
			_, exists := sel.Attr("data-index")
			if exists {
				var name string
				var prev *goquery.Selection = nil
				var on_sale_item OnSaleItem
				if GetNamePrice(sel, &name, prev, &on_sale_item) {
					if GetLink(sel, &on_sale_item) {
						if GetImg(sel, &on_sale_item){
							on_sale_items = append(on_sale_items, on_sale_item)
						}
					}
				}
			}
		},
	)

	return on_sale_items, nil

}


func GetNamePrice(sel *goquery.Selection, name *string, prev *goquery.Selection, on_sale_item *OnSaleItem) bool {
	success := false
	sel.Find("span").Each(
		func(j int, s *goquery.Selection){
			attrib1, exists1 := s.Attr("class")
			attrib2, exists2 := s.Attr("data-a-color")
			//node := s.Find("span").Nodes[0]
			if exists1 {
				if attrib1 == "a-size-base-plus a-color-base" {
					*name += s.Text() + " : "
				} else if attrib1 == "a-size-base-plus a-color-base a-text-normal"{
					*name += s.Text()
				}
				if exists2 {
					if attrib2 == "secondary" && !success {
						success = true
						// if this selection is the element for original price, use this selection and prev to create an OnSaleItem
						price_sale, _ := strconv.ParseFloat(
							prev.Find("span").First().Text()[6:],
							32,
						)
						price_orig, _ := strconv.ParseFloat(
							s.Find("span").First().Text()[6:],
							32,
						)
						on_sale_item.Name = *name
						on_sale_item.Price_orig = float32(price_orig)
						on_sale_item.Price_sale = float32(price_sale)
						on_sale_item.Seller = "Amazon"
					} else {
						prev = s
					}
				}
			}
		},
	)
	return success 
}

func GetLink(sel *goquery.Selection, on_sale_item *OnSaleItem) bool {
	success := false
	sel.Find("a").Each(
		func(i int, s *goquery.Selection){
			link, exists := s.Attr("href")
			if exists && !success {
				success = true
				on_sale_item.Link = "amazon.ca"+link
				return
			}
		},
	)
	return success
}

func GetImg(sel *goquery.Selection, on_sale_item *OnSaleItem) bool {
	success := false
	sel.Find("img").Each(
		func(i int, s *goquery.Selection){
			imglink, exists := s.Attr("src")
			if exists && !success {
				success = true
				on_sale_item.Img = imglink
				return
			}
		},
	)
	return success
}