package models

type Login struct {
	Email           string               `bson:"email" json:"email" validate:"required,email"`
	Password        string               `bson:"password" json:"password" validate:"required,min=8"`
}

type Register struct {
	Name            string               `bson:"name" json:"name" validate:"required,min=3,max=50"`
	Email           string               `bson:"email" json:"email" validate:"required,email"`
	Password        string               `bson:"password" json:"password" validate:"required,min=8"`
	Role            string               `bson:"role" json:"role" validate:"required,oneof=admin author"`
}
