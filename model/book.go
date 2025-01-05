package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// # Book Model
type Book struct {
	ID          *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       *string             `bson:"title" json:"title" validate:"required"`
	Author      *string             `bson:"author" json:"author" validate:"required"`
	PublishYear *int                `bson:"publishYear" json:"publishYear" validate:"required"`
	CreatedAt   *time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   *time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
