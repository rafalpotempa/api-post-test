package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
