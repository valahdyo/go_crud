package main

type LoginRequest struct {
	Email    string
	Password string
}

type UserUpdateRequest struct {
	FullName  string
	Email     string
	Password  string
	BirthDate string
}
