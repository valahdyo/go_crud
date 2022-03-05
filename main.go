package main

import (
	"fmt"
	"log"
	"net/http"
	"go-crud/handler"
	"go-crud/db/migration"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleRequest() {
	//using router gorilla mux
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/api/users", handler.AllUsers).Methods("GET")
	r.HandleFunc("/api/user", handler.AddUser).Methods("POST")
	r.HandleFunc("/api/auth/login", handler.LoginUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", handler.UpdateUser).Methods("PATCH")
	r.HandleFunc("/api/user/{id}", handler.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func main() {
	fmt.Println("Go run on port 5000")
	migration.Migrate()
	handleRequest()
}
