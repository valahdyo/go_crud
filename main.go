package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleRequest() {
	//using router gorilla mux
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/api/users", AllUsers).Methods("GET")
	r.HandleFunc("/api/user", AddUser).Methods("POST")
	r.HandleFunc("/api/auth/login", LoginUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/api/user/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func main() {
	fmt.Println("Go run on port 5000")
	Migrate()
	handleRequest()
}
