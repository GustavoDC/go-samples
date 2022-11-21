package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"Author"`
	Year   int    `json:"text"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: /")
	fmt.Fprintf(w, "Welcome to my book store!")
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: /books")
	json.NewEncoder(w).Encode(Books)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/", allBooks)
	fmt.Println("Listening to port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	Books = []Book{
		{Title: "Leaf by Niggle", Author: "J.R.R. Tolkien", Year: 1945},
		{Title: "The Silmarillon", Author: "J.R.R. Tolkien", Year: 1977},
	}
	handleRequests()
}
