package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"server/models"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Pinged your deployment. You successfully connected to MongoDB!"))
	if err != nil {
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDoc := bson.D{{"name", user.Name}, {"email", user.Email}}

	// Insert POST data
	result, err := models.UsersCollection.InsertOne(context.TODO(), userDoc)
	if err != nil {
		http.Error(w, "Error while insertion", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Marshal the response to JSON
	jsonResponse, err := json.Marshal(map[string]interface{}{
		"user_id": result.InsertedID,
	})
	if err != nil {
		http.Error(w, "Error while marshaling response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Set the content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	reqVars := mux.Vars(r)

	// Check if the ID is in the correct format
	ID, err := primitive.ObjectIDFromHex(reqVars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Decode JSON
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filter := bson.D{{"_id", ID}}
	updatedQuery := bson.D{{"$set", bson.D{
		{"name", user.Name},
		{"email", user.Email},
	},
	}}

	// Update PUT data
	result, err := models.UsersCollection.UpdateOne(context.TODO(), filter, updatedQuery)
	if err != nil {
		http.Error(w, "Error while updating", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Marshal the response to JSON
	jsonResponse, err := json.Marshal(map[string]interface{}{
		"modified_count": result.ModifiedCount,
		"user_id":        ID,
	})
	if err != nil {
		http.Error(w, "Error while marshaling response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Set the content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	reqVars := mux.Vars(r)

	// Check if the ID is in the correct format
	ID, err := primitive.ObjectIDFromHex(reqVars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	filter := bson.D{{"_id", ID}}

	// Delete document
	result, err := models.UsersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, "Error while deleting", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Marshal the response to JSON
	jsonResponse, err := json.Marshal(map[string]interface{}{
		"deleted_count": result.DeletedCount,
		"user_id":       ID,
	})
	if err != nil {
		http.Error(w, "Error while marshaling response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Set the content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
