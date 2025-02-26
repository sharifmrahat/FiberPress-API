package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID                primitive.ObjectID   `bson:"_id,omitempty"`
    Name              string               `bson:"name"`
    Email             string               `bson:"email"`
    Password          string               `bson:"password"`
    Role              string               `bson:"role"` // "Admin" or "Author"
    Posts             []primitive.ObjectID `bson:"posts,omitempty"` // References Post IDs
    IsVerified        bool                 `bson:"isVerified"`
    RequestDeletion   bool                 `bson:"requestDeletion"`
    IsActive          bool                 `bson:"isActive"`  // Soft delete (default: true)
    IsDeleted         bool                 `bson:"isDeleted"` // Soft delete (default: false)
}
