package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homePage)
	fmt.Printf("port: %s", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	for {

	}
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage")
}
