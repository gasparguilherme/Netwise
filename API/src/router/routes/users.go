package routes

import "net/http"

var routesUsers = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Handler: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "/users/{usersID}",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "/users/{usersID}",
		Method: http.MethodPut,
		Handler: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuth: false,
	},
	{
		URI:    "/users/{usersID}",
		Method: http.MethodDelete,
		Handler: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuth: false,
	},
}
