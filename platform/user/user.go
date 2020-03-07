package platform

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Index        rune               `json:"index" bson:"index"`
	Guid         string             `json:"guid" bson:"guid"`
	IsActive     bool               `json:"isActive" bson:"isActive"`
	Balance      string             `json:"balance" bson:"balance"`
	Picture      string             `json:"picture" bson:"picture"`
	Age          rune               `json:"age" bson:"age"`
	EyeColor     string             `json:"eyeColor" bson:"eyeColor"`
	Name         string             `json:"name" bson:"name"`
	Gender       string             `json:"gender" bson:"gender"`
	Company      string             `json:"company" bson:"company"`
	Email        string             `json:"email" bson:"email"`
	Phone        string             `json:"phone" bson:"phone"`
	Address      string             `json:"address" bson:"address"`
	About        string             `json:"about" bson:"about"`
	Registered   string             `json:"registered" bson:"registered"`
	Latitude     float32            `json:"latitude" bson:"latitude"`
	Longitude    float32            `json:"longitude" bson:"longitude"`
	Greeting     string             `json:"greeting" bson:"greeting"`
	FavoriteFood string             `json:"favoriteFood" bson:"favoriteFood"`
}
