package router

import (
	"bookstore/api"

	"github.com/gorilla/mux"
)

// # Router Function
func Router() *mux.Router {
	// # Create a new router object
	router := mux.NewRouter()

	// # Root Route
	router.HandleFunc("/", api.RootRoute).Methods("GET") // # Get the server message

	// # Create a subrouter for books routes
	books := router.PathPrefix("/books").Subrouter()

	// # Add other routes to the subrouter
	books.HandleFunc("", api.CreateBook).Methods("POST")        // # Create a new book
	books.HandleFunc("", api.GetBooks).Methods("GET")           // # Get all books
	books.HandleFunc("", api.DeleteBooks).Methods("DELETE")     // # Delete all books
	books.HandleFunc("/{id}", api.GetBook).Methods("GET")       // # Get a single book
	books.HandleFunc("/{id}", api.UpdateBook).Methods("PUT")    // # Update a book
	books.HandleFunc("/{id}", api.DeleteBook).Methods("DELETE") // # Delete a book

	// # Return Router Object
	return router
}
