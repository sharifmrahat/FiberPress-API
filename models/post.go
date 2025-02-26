package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paragraph struct {
    Heading     *string `bson:"heading,omitempty" validate:"omitempty,max=255"` // Optional, max length 255
    Content     *string `bson:"content,omitempty" validate:"omitempty"`         // Optional
    Description string  `bson:"description" validate:"required"`                // Required
}

type Post struct {
    ID              primitive.ObjectID   `bson:"_id,omitempty"`
    Title           string               `bson:"title" validate:"required,max=255"`
    Intro           string               `bson:"intro,omitempty" validate:"omitempty,max=500"`
    Thumbnail       string               `bson:"thumbnail,omitempty" validate:"omitempty,url"`
    Paragraphs      []Paragraph          `bson:"paragraphs" validate:"required,dive"`
    Conclusion      *string              `bson:"conclusion,omitempty" validate:"omitempty"`
    Tags            []string             `bson:"tags" validate:"required,dive,max=50"`
    CategoryID      primitive.ObjectID   `bson:"categoryId" validate:"required"`
    UserID          primitive.ObjectID   `bson:"userId" validate:"required"` // References User ID
    Status          string               `bson:"status" validate:"required,oneof=Pending Published Archived"`
    ViewCount       int64                `bson:"viewCount"`
    IsTrending      bool                 `bson:"isTrending"`
    RequestDeletion bool                 `bson:"requestDeletion"`
    IsActive        bool                 `bson:"isActive"`
    IsDeleted       bool                 `bson:"isDeleted"`
    CreatedAt       time.Time            `bson:"createdAt,omitempty"`
    UpdatedAt       time.Time            `bson:"updatedAt,omitempty"`
}
