package router

import (
	"github.com/gasparguilherme/Netwise/api/src/router/routes"
	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
