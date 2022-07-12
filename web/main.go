package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	APP_VERSION string = "dev"
	PORT        string = "10080"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print(r)
		fmt.Fprintf(w, "Welcome to the Company Web Site!\n")
		fmt.Fprintf(w, "Version: %s\n", APP_VERSION)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
