package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint(GET): /")
	fmt.Fprintf(w, "Welcome to my book store!")
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint(GET): /books")
	json.NewEncoder(w).Encode(Books)
}

func singleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint(GET): /book/" + key)

	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

// Send from POST Canary: {"id":"4", "title":"The Hobbit", "author":"J.R.R. Tolkien", "year":"1937"}
func newBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint(POST): /book/")
	body, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(body, &book)
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Endpoint(DELETE): /book/" + id)

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", allBooks)
	router.HandleFunc("/book", newBook).Methods("POST")
	router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", singleBook)
	fmt.Println("Listening to port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	Books = []Book{
		{Id: "1", Title: "Leaf by Niggle", Author: "J.R.R. Tolkien", Year: "1945"},
		{Id: "2", Title: "The Silmarillon", Author: "J.R.R. Tolkien", Year: "1977"},
	}
	handleRequests()
}
