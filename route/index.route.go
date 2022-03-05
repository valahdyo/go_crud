package route

import (
	"log"
	"net/http"
	"go-crud/handler"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	//using router gorilla mux
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/users", handler.AllUsers).Methods("GET")
	r.HandleFunc("/api/user", handler.AddUser).Methods("POST")
	r.HandleFunc("/api/auth/login", handler.LoginUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", handler.UpdateUser).Methods("PATCH")
	r.HandleFunc("/api/user/{id}", handler.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}