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

	/*
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
				sel.Find("span").Each(
					func(j int, s *goquery.Selection){
						attrib1, exists1 := s.Attr("class")
						attrib2, exists2 := s.Attr("data-a-color")
						//node := s.Find("span").Nodes[0]
						if exists1 {
							if attrib1 == "a-size-base-plus a-color-base" {
								name += s.Text() + " : "
							} else if attrib1 == "a-size-base-plus a-color-base a-text-normal"{
								name += s.Text()
							}
							if exists2 {
								if attrib2 == "secondary" {
									// if this selection is the element for original price, use this selection and prev to create an OnSaleItem
									price_sale, _ := strconv.ParseFloat(
										prev.Find("span").First().Text()[6:],
										32,
									)
									price_orig, _ := strconv.ParseFloat(
										s.Find("span").First().Text()[6:],
										32,
									)
									on_sale_items = append(
										on_sale_items,
										OnSaleItem{name, float32(price_orig), float32(price_sale), "Amazon"},
									)
								} else {
									prev = s
								}
							}
						}
					},
				)
			}
		},
	)
*/

	return on_sale_items, nil

}