package handler

import (
	"encoding/json"
	"go-crud/db"
	"go-crud/model/entity"
	"go-crud/model/request"
	"go-crud/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users []entity.User
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash password
	hashedPass, _ := utils.HashingPassword(user.Password)
	user.Password = hashedPass

	// Create user
	db.Create(&user)

	respondJSON(w, http.StatusCreated, user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var login request.LoginRequest
	json.NewDecoder(r.Body).Decode(&login)

	// Check user
	var user entity.User
	err := db.First(&user, "email = ?", login.Email).Error
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, struct{ Msg string }{"Email not found"})
		return
	}

	// Check password
	isValid := utils.CheckPasswordHash(login.Password, user.Password)
	if !isValid {
		respondJSON(w, http.StatusInternalServerError, struct{ Msg string }{"User or password didn't match"})
		return
	}

	respondJSON(w, http.StatusOK, struct{ Msg string }{"Login success"})

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Get id param
	params := mux.Vars(r)
	id := params["id"]

	userRequest := new(request.UserUpdateRequest)
	json.NewDecoder(r.Body).Decode(&userRequest)

	// Check user
	var user entity.User
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		respondJSON(w, http.StatusInternalServerError, struct{ Msg string }{"User not found"})
		return
	}

	// Update detail user
	if userRequest.FullName != "" {
		user.FullName = userRequest.FullName
	}
	if userRequest.Password != "" {
		hashedPassword, _ := utils.HashingPassword(userRequest.Password)
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
		respondJSON(w, http.StatusInternalServerError, struct{ Msg string }{"Failed to update user"})

		return
	}
	respondJSON(w, http.StatusOK, struct{ Msg string }{"Update user success"})

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Check id params
	params := mux.Vars(r)
	id := params["id"]

	// Check user
	var user entity.User
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		respondJSON(w, http.StatusInternalServerError, struct{ Msg string }{"User not found"})
		return
	}

	// Delete user
	errDel := db.Delete(&user).Error
	if errDel != nil {
		respondJSON(w, http.StatusInternalServerError, struct{ Msg string }{"Failed to delete user"})
		return
	}

	respondJSON(w, http.StatusOK, struct{ Msg string }{"Delete success"})

}
