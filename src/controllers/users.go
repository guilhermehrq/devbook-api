package controllers

import "net/http"

// CreateUser inserts a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}

// GetUsers gets all users inserted in the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users..."))
}

// GetUser gets the user by the ID given
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user by ID..."))
}

// UpdateUser updates the user by the ID given
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user..."))
}

// DeleteUser delets the user by the ID given
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user..."))
}
