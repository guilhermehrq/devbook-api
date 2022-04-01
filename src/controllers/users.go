package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser inserts a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Reading body of request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Validating user
	err = user.Prepare("CREATE")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	// Creating a user
	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers gets all users inserted in the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Reading the query params
	userToSearch := strings.ToLower(r.URL.Query().Get("search"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	// Searching the users
	repository := repositories.NewUserRepository(db)
	users, err := repository.Search(userToSearch)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUser gets the user by the ID given
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Reading the parameters of request
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	// Searching the user by the ID given
	repository := repositories.NewUserRepository(db)
	user, err := repository.GetByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Validating if any user was found
	if user.ID == 0 {
		err = errors.New(fmt.Sprintf("User ID %d not found!", userID))

		responses.Error(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates the user by the ID given
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Reading the parameters of request
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Reading body of request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Validating user
	err = user.Prepare("UPDATE")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	// Validating if the user exists in the database
	repository := repositories.NewUserRepository(db)

	userValidate, err := repository.GetByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userValidate.ID == 0 {
		err = errors.New(fmt.Sprintf("User ID %d not found!", userID))

		responses.Error(w, http.StatusNotFound, err)
		return
	}

	// Updating the user
	err = repository.Update(userID, user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser delets the user by the ID given
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	userValidate, err := repository.GetByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userValidate.ID == 0 {
		err = errors.New(fmt.Sprintf("User ID %d not found!", userID))

		responses.Error(w, http.StatusNotFound, err)
		return
	}

	err = repository.Delete(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
