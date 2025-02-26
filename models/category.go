package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
    ID          primitive.ObjectID   `bson:"_id,omitempty"`
    Name        string               `bson:"name" validate:"required,max=255"`
    ImageURL    string               `bson:"imageUrl,omitempty" validate:"omitempty,url"`
    Description string               `bson:"description,omitempty" validate:"omitempty,max=500"`
    PostIDs     []primitive.ObjectID `bson:"postIds,omitempty" validate:"dive"`
    ViewCount   int64                `bson:"viewCount"`
    IsTrending  bool                 `bson:"isTrending"`
    IsActive    bool                 `bson:"isActive"`
    IsDeleted   bool                 `bson:"isDeleted"`
    CreatedAt   time.Time            `bson:"createdAt,omitempty"`
    UpdatedAt   time.Time            `bson:"updatedAt,omitempty"`
}
