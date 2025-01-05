package api

import (
	"bookstore/app"
	"bookstore/model"
	"bookstore/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// # Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// # Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// # Create Book Object
	var book model.Book

	// # Create New Decoder for Request Body
	decoder := json.NewDecoder(r.Body)

	// # Decode Request Body to Book Object
	err := decoder.Decode(&book)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	// # Get Current Time
	time, err := utils.GetCurrentTime()
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Set CreatedAt and UpdatedAt Time
	book.CreatedAt = time
	book.UpdatedAt = time

	// # Insert Book into Database
	res, err := app.InsertBook(&book, context.Background())
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Send Success Response
	utils.SendSuccess(w, res)
}

// # Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// # Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// # Get Params from Request URL
	params := mux.Vars(r)

	// # Create Updates Object for Book
	var updates bson.M

	// # Create New Decoder for Request Body
	decoder := json.NewDecoder(r.Body)

	// # Decode Request Body to Updates Object
	err := decoder.Decode(&updates)
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Get Current Time
	time, err := utils.GetCurrentTime()
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Set UpdatedAt Time
	updates["updatedAt"] = time

	// # Update Book in Database
	res, err := app.UpdateBook(params["id"], &updates, context.Background())
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Send Success Response
	utils.SendSuccess(w, res)
}

// # Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// # Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// # Get Params from Request URL
	params := mux.Vars(r)

	// # Delete Book from Database
	res, err := app.DeleteBook(params["id"], context.Background())
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Send Success Response
	utils.SendSuccess(w, res)
}

// # Delete All Books
func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	// # Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// # Delete All Books from Database
	res, err := app.DeleteBooks(context.Background())
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Send Success Response
	utils.SendSuccess(w, res)
}

// # Get Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	// # Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	// # Get Params from Request URL
	params := mux.Vars(r)

	// # Get Book from Database
	book, err := app.GetBook(params["id"], context.Background())
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Send Success Response
	utils.SendSuccess(w, book)
}

// # Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// # Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	// # Get All Books from Database
	books, err := app.GetBooks(context.Background())
	if err != nil {
		// # Send Error Response
		utils.SendError(w, err)
		return
	}

	// # Send Success Response
	utils.SendSuccess(w, books)
}
