package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepage!")
	fmt.Println("Endpoint: /")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	fmt.Println("Listening to port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequests()
}
