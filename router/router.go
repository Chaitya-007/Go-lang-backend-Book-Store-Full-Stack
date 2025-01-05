package router

import (
	"bookstore/api"

	"github.com/gorilla/mux"
)

// # Router
func Router() *mux.Router {
	// # Create a new router object
	router := mux.NewRouter()

	// # Root Route
	router.HandleFunc("/", api.RootRoute).Methods("GET") // # Get server welcome message

	// # API Base Path
	BasePath := "/api/v1"

	// # API Book Resource Path
	BookResPath := "/book"

	// # API Books Resource Path
	BooksResPath := "/books"

	// # API Endpoints Slice
	endpoints := []string{BasePath + BookResPath, BasePath + BooksResPath}

	// # Create a subrouter for the book route
	book := router.PathPrefix(endpoints[0]).Subrouter()

	// # Create a subrouter for the books route
	books := router.PathPrefix(endpoints[1]).Subrouter()

	// # Add routes to the subrouter
	book.HandleFunc("", api.CreateBook).Methods("POST")        // # Create a new book
	books.HandleFunc("", api.GetBooks).Methods("GET")          // # Get all books
	books.HandleFunc("", api.DeleteBooks).Methods("DELETE")    // # Delete all books
	book.HandleFunc("/{id}", api.GetBook).Methods("GET")       // # Get a single book
	book.HandleFunc("/{id}", api.UpdateBook).Methods("PUT")    // # Update a book
	book.HandleFunc("/{id}", api.DeleteBook).Methods("DELETE") // # Delete a book

	// # Return Router Object
	return router
}
