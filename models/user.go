package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name            string               `bson:"name" json:"name" validate:"required,min=3,max=50"`
	Email           string               `bson:"email" json:"email" validate:"required,email"`
	Password        string               `bson:"password" json:"password" validate:"required,min=8"`
	Role            string               `bson:"role" json:"role" validate:"required,oneof=admin author"`
	Posts           []primitive.ObjectID `bson:"posts,omitempty" json:"posts"`
	IsVerified      bool                 `bson:"isVerified" json:"isVerified"`
	RequestDeletion bool                 `bson:"requestDeletion" json:"requestDeletion"`
	IsActive        bool                 `bson:"isActive" json:"isActive"`
	IsDeleted       bool                 `bson:"isDeleted" json:"isDeleted"`
	CreatedAt       time.Time            `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt       time.Time            `bson:"updatedAt,omitempty" json:"updatedAt"`
}
