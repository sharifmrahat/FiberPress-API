package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
    ID         primitive.ObjectID   `bson:"_id,omitempty"`
    Name       string               `bson:"name"`
    ImageURL   string               `bson:"imageUrl,omitempty"`
	Description string              `bson:"description,omitempty"`
    PostIDs    []primitive.ObjectID `bson:"postIds,omitempty"` // References Posts
    ViewCount  int64                `bson:"viewCount"`
    IsTrending bool                 `bson:"isTrending"`
    IsActive   bool                 `bson:"isActive"`
    IsDeleted  bool                 `bson:"isDeleted"`
}
