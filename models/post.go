package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Paragraph struct {
    Heading     *string `bson:"heading,omitempty"`   // Optional
    Content     *string `bson:"content,omitempty"`   // Optional
    Description string  `bson:"description"`         // Required
}

type Post struct {
    ID              primitive.ObjectID   `bson:"_id,omitempty"`
    Title           string               `bson:"title"`
    Intro           string               `bson:"intro,omitempty"`
    Thumbnail       string               `bson:"thumbnail,omitempty"`
    Paragraphs      []Paragraph          `bson:"paragraphs"`
    Conclusion      *string              `bson:"conclusion,omitempty"` // Optional
    Tags            []string             `bson:"tags"`
    CategoryID      primitive.ObjectID   `bson:"categoryId"`
    UserID          primitive.ObjectID   `bson:"userId"` // References User ID
    Status          string               `bson:"status"` // Pending | Published | Archived
    ViewCount       int64                `bson:"viewCount"`
    IsTrending      bool                 `bson:"isTrending"`
    RequestDeletion bool                 `bson:"requestDeletion"`
    IsActive        bool                 `bson:"isActive"`
    IsDeleted       bool                 `bson:"isDeleted"`
}
