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
				var is_on_sale uint8
				GetNamePrice(sel, &name, prev, &on_sale_item, &is_on_sale)
				GetLink(sel, &on_sale_item, &is_on_sale)
				GetImg(sel, &on_sale_item, &is_on_sale)
				if is_on_sale == 0x7 {
					on_sale_items = append(on_sale_items, on_sale_item)
				}
			}
		},
	)

	return on_sale_items, nil

}


func GetNamePrice(sel *goquery.Selection, name *string, prev *goquery.Selection, on_sale_item *OnSaleItem, is_on_sale *uint8) {
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
					if attrib2 == "secondary" {
						*is_on_sale |= 0x1
						// if this selection is the element for original price, use this selection and prev to create an OnSaleItem
						price_sale, _ := strconv.ParseFloat(
							prev.Find("span").First().Text()[6:],
							32,
						)
						price_orig, _ := strconv.ParseFloat(
							s.Find("span").First().Text()[6:],
							32,
						)
						on_sale_item.name = *name
						on_sale_item.price_orig = float32(price_orig)
						on_sale_item.price_sale = float32(price_sale)
						on_sale_item.seller = "Amazon"
					} else {
						prev = s
					}
				}
			}
		},
	)
}

func GetLink(sel *goquery.Selection, on_sale_item *OnSaleItem, is_on_sale *uint8) {
	sel.Find("a").Each(
		func(i int, s *goquery.Selection){
			link, exists := s.Attr("href")
			if exists {
				*is_on_sale |= 0x2
				on_sale_item.link = "amazon.ca"+link
			}
		},
	)
}

func GetImg(sel *goquery.Selection, on_sale_item *OnSaleItem, is_on_sale *uint8) {
	sel.Find("img").Each(
		func(i int, s *goquery.Selection){
			imglink, exists := s.Attr("src")
			if exists {
				*is_on_sale |= 0x4
				on_sale_item.img = imglink
			}
		},
	)
}