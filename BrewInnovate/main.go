package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *sql.DB

func main() {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server up & running...")

	log.Fatal(http.ListenAndServe(":8090", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi from main")
	w.Write([]byte("up and running"))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book Book
	err := db.QueryRow("SELECT * FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow("INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id", book.Title, book.Author).Scan(&book.ID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("UPDATE books SET title = $1, author = $2 WHERE id = $3", book.Title, book.Author, id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Book with ID %d has been deleted", id)
}
