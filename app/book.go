package app

import (
	"context"
	"errors"
	"fmt"

	"bookstore/db"
	"bookstore/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// # Insert One Book
func InsertBook(book *model.Book, ctx context.Context) (string, error) {
	// # Get MongoDB Collection Object
	collection := db.GetCollection()

	// # Insert Book into MongoDB Collection
	res, err := collection.InsertOne(ctx, *book)
	if err != nil {
		err = errors.New("Failed to insert book: " + err.Error())
		return "", err
	}

	// # Convert Inserted ID from ObjectID to String
	bookID := res.InsertedID.(primitive.ObjectID).Hex()

	// # Create Success Message
	successMsg := fmt.Sprintf("Inserted 1 book with ID: %v", bookID)

	// # Return Success Message
	return successMsg, nil
}

// # Update One Book
func UpdateBook(bookID string, updates *bson.M, ctx context.Context) (string, error) {
	// # Get MongoDB Collection Object
	collection := db.GetCollection()

	// # Convert Book ID from String to ObjectID
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		err = errors.New("Invalid Book ID: " + err.Error())
		return "", err
	}

	// # Filter to Find Book by ID
	filter := bson.M{"_id": id}

	// # Update Book in MongoDB Collection
	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": updates})
	if err != nil {
		err = errors.New("Failed to update book: " + err.Error())
		return "", err
	}

	// # Create Success Message
	successMsg := fmt.Sprintf("Updated 1 book with ID: %v and Modified Count: %v", bookID, res.ModifiedCount)

	// # Return Success Message
	return successMsg, nil
}

// # Delete One Book
func DeleteBook(bookID string, ctx context.Context) (string, error) {
	// # Get MongoDB Collection Object
	collection := db.GetCollection()

	// # Convert Book ID from String to ObjectID
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		err = errors.New("Invalid Book ID: " + err.Error())
		return "", err
	}

	// # Filter to Find Book by ID
	filter := bson.M{"_id": id}

	// # Check if Book Exists in MongoDB Collection
	check := collection.FindOne(ctx, filter)
	if check.Err() != nil {
		err = errors.New("Book with ID: " + bookID + " not found")
		return "", err
	}

	// # Delete Book from MongoDB Collection
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		err = errors.New("Failed to delete book: " + err.Error())
		return "", err
	}

	// # Create Success Message
	successMsg := fmt.Sprintf("Deleted 1 book with ID: %v and Deleted Count: %v", bookID, res.DeletedCount)

	// # Return Success Message
	return successMsg, nil
}

// # Delete All Books
func DeleteBooks(ctx context.Context) (string, error) {
	// # Get MongoDB Collection Object
	collection := db.GetCollection()

	// # Filter to Delete All Books
	filter := bson.D{{}}

	// # Check if Books Exist in MongoDB Collection
	check := collection.FindOne(ctx, filter)
	if check.Err() != nil {
		err := errors.New("No books found")
		return "", err
	}

	// # Delete All Books from MongoDB Collection
	res, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		err = errors.New("Failed to delete books: " + err.Error())
		return "", err
	}

	// # Create Success Message
	successMsg := fmt.Sprintf("Number of books deleted: %v", res.DeletedCount)

	// # Return Success Message
	return successMsg, nil

}

// # Get One Book
func GetBook(bookID string, ctx context.Context) (*model.Book, error) {
	// # Get MongoDB Collection Object
	collection := db.GetCollection()

	// # Convert Book ID from String to ObjectID
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		err = errors.New("Invalid Book ID: " + err.Error())
		return nil, err
	}

	// # Filter to Find Book by ID
	filter := bson.M{"_id": id}

	// # Create Book Object
	var book model.Book

	// # Find Book from MongoDB Collection
	err = collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		err = errors.New("Failed to get book with ID: " + bookID + " " + err.Error())
		return nil, err
	}

	// # Return Book
	return &book, nil
}

// # Get All Books
func GetBooks(ctx context.Context) (*[]model.Book, error) {
	// # Get MongoDB Collection Object
	collection := db.GetCollection()

	// # Filter to Find All Books
	filter := bson.D{{}}

	// # Find All Books from MongoDB Collection
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		err = errors.New("Failed to get books: " + err.Error())
		return nil, err
	}

	// # Create Books Slice
	var books []model.Book

	// # Iterate Over Books
	for cur.Next(ctx) {
		// # Create Book Object
		var book model.Book

		// # Decode Book
		err := cur.Decode(&book)
		if err != nil {
			err = errors.New("Failed to decode book: " + err.Error())
			return nil, err
		}

		// # Append Book to Books Slice
		books = append(books, book)
	}

	// # Close Cursor
	defer cur.Close(ctx)

	// # Return Books Slice
	return &books, nil
}
