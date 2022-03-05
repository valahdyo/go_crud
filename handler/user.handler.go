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
	hashedPass, _ := utils.HashingPassword(user.Password)
	user.Password = hashedPass
	db.Create(&user)

	respondJSON(w, http.StatusCreated, user)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var login request.LoginRequest
	json.NewDecoder(r.Body).Decode(&login)
	var user entity.User
	err := db.First(&user, "email = ?", login.Email).Error
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, struct {Msg string} {"Email not found"})
		return
	}

	isValid := utils.CheckPasswordHash(login.Password, user.Password)
	if !isValid {
		respondJSON(w, http.StatusInternalServerError, "User or password didn't match")
		return
	}

	respondJSON(w, http.StatusOK, struct {Msg string} {"Login success"})

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	params := mux.Vars(r)
	id := params["id"]
	userRequest := new(request.UserUpdateRequest)
	json.NewDecoder(r.Body).Decode(&userRequest)

	var user entity.User
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		respondJSON(w, http.StatusInternalServerError, struct {Msg string} {"User not found"})
		return
	}

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
		respondJSON(w, http.StatusInternalServerError, struct {Msg string} {"Failed to update user"})

		return
	}
	respondJSON(w, http.StatusOK, struct {Msg string} {"Update user success"})

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := db.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	params := mux.Vars(r)
	id := params["id"]

	var user entity.User
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		respondJSON(w, http.StatusInternalServerError, struct {Msg string} {"User not found"})
		return
	}

	errDel := db.Delete(&user).Error
	if errDel != nil {
		respondJSON(w, http.StatusInternalServerError, struct {Msg string} {"Failed to delete user"})
		return
	}

	respondJSON(w, http.StatusOK, struct {Msg string} {"Delete success"})

}
