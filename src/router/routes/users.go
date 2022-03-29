package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	// CREATE USER
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Func:           controllers.CreateUser,
		AuthIsRequired: false,
	},
	// GET USERS
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Func:           controllers.GetUsers,
		AuthIsRequired: false,
	},
	// GET USER BY ID
	{
		URI:            "/users/{userID}",
		Method:         http.MethodGet,
		Func:           controllers.GetUser,
		AuthIsRequired: false,
	},
	// UPDATE USER
	{
		URI:            "/users/{userID}",
		Method:         http.MethodPut,
		Func:           controllers.UpdateUser,
		AuthIsRequired: false,
	},
	// DELETE USER
	{
		URI:            "/users/{userID}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		AuthIsRequired: false,
	},
}
