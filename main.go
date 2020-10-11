package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(SearchPage))
	mux.Handle("/results", http.HandlerFunc(ResultPage))
	http.ListenAndServe("localhost:12345", mux)

}