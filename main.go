package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Fatal(err)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params from request

	for _, book := range books {
		if book.ID == params["id"] {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode(&Book{})
	if err != nil {
		log.Fatal(err)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id
	books = append(books, book)

	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Fatal(err)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)

			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
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
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8000", r))
}
