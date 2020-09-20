package main

import "net/http"
import "io/ioutil"
import "fmt"

func PrintResponse(resp *http.Response) {
	body_bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(body_bytes))
}

