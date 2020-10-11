package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func SearchPage(w http.ResponseWriter, req *http.Request) {
	if req.Method=="GET" {
		var templates *template.Template = template.Must(template.ParseGlob("templates/*.html"))
		templates.ExecuteTemplate(w, "index.html", nil)
	} else if req.Method=="POST" {
		req.ParseForm()
		search_string := req.PostForm.Get("search_string")
		http.Redirect(w, req, "/results?search_string="+search_string, http.StatusSeeOther)
	}
	//w.Header().Set("Content-Type", "text/html; charset=utf8")
	//fmt.Fprintf(w, "<html><body><h1>Search page placeholde</h1></body></html>")
}

func ResultPage(w http.ResponseWriter, req *http.Request) {
	search_string, ok := req.URL.Query()["search_string"]
	fmt.Fprintf(w, "%+v\n", search_string)
	if !ok || len(search_string[0])<1 {
		fmt.Fprintf(w, "Bad search string\n")
		return
	}
	var search_phrase string
	for i:=0; i<len(search_string); i++ {
		search_phrase += search_string[i]+" "
	}
	fmt.Fprintf(w, "%+v\n", search_phrase)
	crawler_for_seller := AssignCrawlerFunctions()
	on_sale_items := CrawlSellers(crawler_for_seller, search_phrase)
	fmt.Fprintf(w, OnSaleItemList(on_sale_items).ToString())

	//var templates *template.Template = template.Must(template.ParseGlob("templates/*.html"))
	//templates.ExecuteTemplate(w, "index.html", nil)
}