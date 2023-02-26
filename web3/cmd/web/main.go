package main

import (
	"log"
	"net/http"
	"strconv"
	"web3/pkg/handlers"
)

var port = 8080
var address = "localhost:" + strconv.Itoa(port)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	err := http.ListenAndServe(address, nil)
	log.Fatal(err)
}
