package routes

import (
	"net/http"

	"github.com/gasparguilherme/Netwise/api/src/controllers"
)

var routesUsers = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Handler:      controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Handler:      controllers.GetAllUsers,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{usersID}",
		Method:       http.MethodGet,
		Handler:      controllers.GetUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{usersID}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{usersID}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeleteUser,
		RequiresAuth: false,
	},
}
