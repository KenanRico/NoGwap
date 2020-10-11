package main

import "fmt"

type CrawlerFunc func(string)([]OnSaleItem, error)

func AssignCrawlerFunctions() []CrawlerFunc {
	var funcs []CrawlerFunc
	funcs = append(funcs, AmazonCrawler)
	funcs = append(funcs, BestbuyCrawler)
	return funcs
}

func CrawlSellers(crawler_funcs []CrawlerFunc, search_phrase string) []OnSaleItem{
	var on_sale_items []OnSaleItem
	// TODO: Change this loop to be multithreaded later on
	for _, crawler := range crawler_funcs {
		items, err := crawler(search_phrase)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		on_sale_items = append(on_sale_items, items...)
	}
	return on_sale_items
}