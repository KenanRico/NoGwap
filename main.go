package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Please specify product name after cmd\n")
		os.Exit(-1)
	}
	var search_phrase string
	for i:=1; i<len(os.Args); i++ {
		search_phrase += os.Args[i]+" "
	}
	crawler_for_seller := AssignCrawlerFunctions()
	on_sale_items := CrawlSellers(crawler_for_seller, search_phrase)

	fmt.Println(OnSaleItemList(on_sale_items).ToString())
}