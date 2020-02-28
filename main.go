package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hotpotatoc/rest-api-mux/models"

	"github.com/gorilla/mux"
)

// Init books var as slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params from request

	for _, book := range books {
		if book.ID == params["id"] {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Book{})
	if err != nil {
		return
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func editBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{
		ID:     "1",
		Isbn:   "2462462351",
		Title:  "A Tale of Two Cities",
		Author: &Author{Firstname: "John", Lastname: "Doe"},
	})
	books = append(books, Book{
		ID:     "2",
		Isbn:   "9456846859",
		Title:  "The Busy Road",
		Author: &Author{Firstname: "Richard", Lastname: "Smith"},
	})

	// Register endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", editBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8000", r))
}
