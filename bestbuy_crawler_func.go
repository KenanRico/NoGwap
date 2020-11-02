package main

import (
	_ "github.com/PuerkitoBio/goquery"
	_ "net/http"
	_ "strings"
	_ "errors"
	_ "strconv"
)

func BestbuyCrawler(search_phrase string) ([]OnSaleItem, error){

	var on_sale_items []OnSaleItem

	// Get request search results (currently only get 1st page)
	search_phrase = strings.ReplaceAll(search_phrase, " ", "+")
	url := "https://www.bestbuy.ca/en-ca/search?search=" + search_phrase
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return on_sale_items, errors.New("Bestbuy: "+err.Error())
	}
	resp, err := new(http.Client).Do(request)
	if err != nil {
		return on_sale_items, errors.New("Bestbuy: "+err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return on_sale_items, errors.New("Bestbuy: "+resp.Status)
	}

	// Create goquery.Document for search
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return on_sale_items, errors.New("Bestbuy: "+resp.Status)
	}

	doc.Find("div").Each(
		func(i int, sel *goquery.Selection){
			attr, exists := sel.Attr("class")
			if exists && attr=="col-xs-8_1VO-Q col-sm-12_1kbJA productItemTextContainer_HocvR"{
				var on_sale_item OnSaleItem
				if GetNamePrice(sel, &on_sale_item) && GetLink(sel, &on_sale_item) && GetImg(sel, &on_sale_item) {
							on_sale_items = append(on_sale_items, on_sale_item)
				}
			}
		},
	)

	return on_sale_items, nil

}


func GetNamePrice(sel *goquery.Selection, on_sale_item *OnSaleItem) bool {
	success := false
	var name string
	var prev *goquery.Selection
	sel.Find("div").Each(
		func(i int, s *goquery.Selection) {
			attrib, exists := s.Attr("class")
			if exists && attrib=="productItemName_3IZ3c" {
				name = s.Text()
			}
		}
	)
	sel.Find("span").Each(
		func(i int, s *goquery.Selection) {
			price_attrib, exists1 := s.Attr("itemprop")
			saving_attrib, exists2 := s.Attr("class")
			if exists1 && price_attrib=="offers" {
				prev = s
			}
			if exists2 && saving_attrib=="productSaving_3YmNX undefined" {
				success = true
				// use prev to retrieve price sale, and use s (current span) to retrieve discount, which is used to compute price orig
				on_sale_item.Name = name
				on_sale_item.Price_orig = float32(price_orig)
				on_sale_item.Price_sale = float32(price_sale)
				on_sale_item.Seller = "Bestbuy"
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
				on_sale_item.Link = "https://www.amazon.ca"+link
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