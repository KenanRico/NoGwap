package main

import (
	"net/http"
	"fmt"
	"html/template"
	"encoding/json"
)

func SearchPage(w http.ResponseWriter, req *http.Request) {
	if req.Method=="GET" {
		var templates *template.Template = template.Must(template.ParseGlob("templates/index.html"))
		templates.ExecuteTemplate(w, "index.html", nil)
	} else if req.Method=="POST" {
		req.ParseForm()
		search_string := req.PostForm.Get("search_string")
		http.Redirect(w, req, "/results?search_string="+search_string, http.StatusSeeOther)
	}
}

func ResultPage(w http.ResponseWriter, req *http.Request) {
	if req.Method=="GET" {
		var templates *template.Template = template.Must(template.ParseGlob("templates/result.html"))
		// Retrieve search string from URL
		search_string, ok := req.URL.Query()["search_string"]
		if !ok || len(search_string[0])<1 {
			fmt.Fprintf(w, "Bad search string\n")
			return
		}
		// Call crawlers using search string; crawlers generate on_sale_items list
		var search_phrase string
		for i:=0; i<len(search_string); i++ {
			search_phrase += search_string[i]+" "
		}
		crawler_for_seller := AssignCrawlerFunctions()
		on_sale_items := CrawlSellers(crawler_for_seller, search_phrase)
		_ = on_sale_items
		// Send on_sale_items data to frontend
		j, err := json.Marshal(on_sale_items)
		if err != nil {
			panic(err)
		}
		templates.ExecuteTemplate(w, "result.html", string(j))
	} else if req.Method=="POST" {
		req.ParseForm()
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}