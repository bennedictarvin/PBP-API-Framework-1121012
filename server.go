package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

// Book adalah struktur data untuk mewakili buku
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Book 1", Author: "Author 1"},
	{ID: "2", Title: "Book 2", Author: "Author 2"},
	// Tambahkan contoh buku lainnya di sini jika diperlukan
}

func main() {
	// Initialize Martini
	m := martini.Classic()

	// Define CRUD endpoints
	m.Get("/books", getBooks)
	m.Get("/books/:id", getBook)
	m.Post("/books", createBook)
	m.Put("/books/:id", updateBook)
	m.Delete("/books/:id", deleteBook)

	// Start server
	m.Run()
}

// getBooks mengembalikan daftar semua buku
func getBooks() string {
	// Convert books slice to JSON
	booksJSON, err := json.Marshal(books)
	if err != nil {
		return "Error retrieving books"
	}
	return string(booksJSON)
}

// getBook mengembalikan buku dengan ID yang sesuai
func getBook(params martini.Params) string {
	bookID := params["id"]
	for _, book := range books {
		if book.ID == bookID {
			bookJSON, err := json.Marshal(book)
			if err != nil {
				return "Error retrieving book"
			}
			return string(bookJSON)
		}
	}
	return "Book not found"
}

// createBook menambahkan buku baru
func createBook(res http.ResponseWriter, req *http.Request) string {
	decoder := json.NewDecoder(req.Body)
	var newBook Book
	err := decoder.Decode(&newBook)
	if err != nil {
		http.Error(res, "Error decoding request body", http.StatusBadRequest)
		return ""
	}
	// Generate unique ID for the new book (you might use a UUID library for this in a real-world scenario)
	newBook.ID = strconv.Itoa(len(books) + 1)
	books = append(books, newBook)
	return "Book added successfully"
}

// updateBook memperbarui buku dengan ID yang sesuai
func updateBook(params martini.Params, req *http.Request) string {
	bookID := params["id"]
	decoder := json.NewDecoder(req.Body)
	var updatedBook Book
	err := decoder.Decode(&updatedBook)
	if err != nil {
		return "Error decoding request body"
	}
	for i, book := range books {
		if book.ID == bookID {
			books[i] = updatedBook
			return "Book updated successfully"
		}
	}
	return "Book not found"
}

// deleteBook menghapus buku dengan ID yang sesuai
func deleteBook(params martini.Params) string {
	bookID := params["id"]
	for i, book := range books {
		if book.ID == bookID {
			books = append(books[:i], books[i+1:]...)
			return "Book deleted successfully"
		}
	}
	return "Book not found"
}
