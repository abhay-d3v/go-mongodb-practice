package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/user", controllers.AddUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	fmt.Println("Server is starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
