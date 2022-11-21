package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	// Using gorilla/mux router so we can later retrieve path and query parameters
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", allBooks)
	fmt.Println("Listening to port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	Books = []Book{
		{Title: "Leaf by Niggle", Author: "J.R.R. Tolkien", Year: 1945},
		{Title: "The Silmarillon", Author: "J.R.R. Tolkien", Year: 1977},
	}
	handleRequests()
}
