package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Unique_Id  string             `json:"unique_id" bson:"unique_id" validate:"required,alphanum,len=12"`
	Name       string             `json:"name" bson:"name" validate:"required"`
	Address    string             `json:"address" bson:"address"`
	Contact_No string             `json:"contact_no" bson:"contact_no" validate:"required,e164"`
	Email      string             `json:"email" bson:"email" validate:"required,email"`
}
