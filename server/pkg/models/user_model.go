package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	UserCollection = "users"
)

type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" validate:""`
	Username string             `bson:"username,omitempty" validate:"required,gte=4,lte=18"`
	Email    string             `bson:"email,omitempty" validate:"required,email"`
	Password string             `bson:"password,omitempty" validate:"required"`
}
