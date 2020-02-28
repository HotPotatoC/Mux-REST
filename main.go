package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hotpotatoc/Mux-REST/models"
)

// Init books var as slice Book struct
var books []models.Book

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
	err := json.NewEncoder(w).Encode(&models.Book{})
	if err != nil {
		return
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id
	books = append(books, book)

	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)

			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				return
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func main() {
	r := mux.NewRouter()

	// Mock data
	books = append(books, models.Book{
		ID:     "1",
		Isbn:   "2462462351",
		Title:  "A Tale of Two Cities",
		Author: &models.Author{Firstname: "John", Lastname: "Doe"},
	})
	books = append(books, models.Book{
		ID:     "2",
		Isbn:   "9456846859",
		Title:  "The Busy Road",
		Author: &models.Author{Firstname: "Richard", Lastname: "Smith"},
	})

	// Register endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8000", r))
}
