package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string `json:"Id"`
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

func singleBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: /book/{id}")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", allBooks)
	router.HandleFunc("/book/{id}", singleBook)
	fmt.Println("Listening to port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	Books = []Book{
		{Id: "1", Title: "Leaf by Niggle", Author: "J.R.R. Tolkien", Year: 1945},
		{Id: "2", Title: "The Silmarillon", Author: "J.R.R. Tolkien", Year: 1977},
	}
	handleRequests()
}
