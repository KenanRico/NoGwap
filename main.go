package main

import (
	"fmt"
)

func main() {
	search_phrase := "Nike shoes" //placeholder
	crawler_for_seller := AssignCrawlerFunctions()
	on_sale_items := CrawlSellers(crawler_for_seller, search_phrase)
	fmt.Printf("%v\n", on_sale_items)
}