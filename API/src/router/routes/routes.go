package routes

import "net/http"

// Route representa todas as rotas da API
type Route struct {
	URI          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}
