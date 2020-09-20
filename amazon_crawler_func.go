package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"errors"
	"fmt"
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

	_ = doc
	doc.Find("span").Each(
		func(i int, s *goquery.Selection) {
			attrib1, exists1 := s.Attr("class")
			attrib2, exists2 := s.Attr("data-a-color")
			price := s.Find("span").Text()
			if exists1 && exists2 {
				fmt.Printf("%s, %s\n", attrib1, attrib2)
				fmt.Printf("%s\n", price)
			}
		},
	)

	return on_sale_items, nil

}