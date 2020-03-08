package platform

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Guid         string             `json:"guid,omitempty" bson:"guid,omitempty"`
	Picture      string             `json:"picture,omitempty" bson:"picture,omitempty"`
	Age          rune               `json:"age,omitempty" bson:"age,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Gender       string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Company      string             `json:"company,omitempty" bson:"company,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Registered   string             `json:"registered,omitempty" bson:"registered,omitempty"`
}
