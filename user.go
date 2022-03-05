package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Password  string
	FullName  string
	BirthDate string
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user User
	json.NewDecoder(r.Body).Decode(&user)
	hashedPass, _ := HashingPassword(user.Password)
	user.Password = hashedPass
	fmt.Println(hashedPass)
	db.Create(&user)

	json.NewEncoder(w).Encode(&user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var login LoginRequest
	json.NewDecoder(r.Body).Decode(&login)
	var user User
	err := db.First(&user, "email = ?", login.Email).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	isValid := CheckPasswordHash(login.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode("Login Success")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	params := mux.Vars(r)
	id := params["id"]
	userRequest := new(UserUpdateRequest)
	json.NewDecoder(r.Body).Decode(&userRequest)

	var user User
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if userRequest.FullName != "" {
		user.FullName = userRequest.FullName
	}
	if userRequest.Password != "" {
		hashedPassword, _ := HashingPassword(userRequest.Password)
		user.Password = hashedPassword
	}
	if userRequest.Email != "" {
		user.Email = userRequest.Email
	}
	if userRequest.BirthDate != "" {
		user.BirthDate = userRequest.BirthDate
	}

	errUpdate := db.Save(&user).Error
	if errUpdate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode("Update Success")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	params := mux.Vars(r)
	id := params["id"]

	var user User
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	errDel := db.Delete(&user).Error
	if errDel != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode("Delete Success")

}

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedByte), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
