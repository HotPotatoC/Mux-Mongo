package userhandler

import (
	"context"
	"encoding/json"
	"log"
	"mux-mongo/database"
	platform "mux-mongo/platform/user"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database = database.Connect()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	users := []platform.User{}

	collection := db.Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user platform.User
		_ = cursor.Decode(&user)

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	defer r.Body.Close()
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement GET/:id user
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement POST user
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement PUT user
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement DELETE user
}
