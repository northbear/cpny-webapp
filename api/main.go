package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	APP_VERSION string = "dev"
	PORT        string = "10080"
	URL_PREFIX  string = os.Getenv("URL_PATH_PREFIX")
)

type ApiResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func FailApiMessage(m string) string {
	d, err := json.Marshal(ApiResponse {
		Message: m,
		Version: APP_VERSION,
		Status: "Fail",
	})
	if err != nil {
		log.Fatal("Fatal: Cannot render a fail message")
	}

	return string(d)
}

func main() {
	http.HandleFunc(URL_PREFIX+"/", func(w http.ResponseWriter, r *http.Request) {
		log.Print(r)
		msg := ApiResponse{
			Message: "Welcome to the Company API",
			Version: APP_VERSION,
			Status:  "OK",
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		data, err := json.Marshal(msg)
		if err != nil {
			log.Println("Error: cannot marshal a response message")
			fmt.Fprintln(w, FailApiMessage("A data rendering fails"))
			return
		}
		fmt.Fprintln(w, string(data))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
