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

	var prev *goquery.Selection = nil
	doc.Find("span").Each(
		func(i int, s *goquery.Selection) {
			attrib1, exists1 := s.Attr("class")
			attrib2, exists2 := s.Attr("data-a-color")
			//node := s.Find("span").Nodes[0]
			if exists1 && exists2 {
				if attrib2 == "secondary" {
					// if this selection is the element for original price, use this selection and prev to create an OnSaleItem
					pattrib1, _ := prev.Attr("class")
					pattrib2, _ := prev.Attr("data-a-color")
					fmt.Printf("%s, %s, %s\n", pattrib1, pattrib2, prev.Find("span").Text())
					fmt.Printf("%s, %s, %s\n", attrib1, attrib2, s.Find("span").Text())
				} else {
					prev = s
				}
			}
		},
	)

	return on_sale_items, nil

}