package controllers

import "net/http"

// criar usuario insere um usuario no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user!"))
}

// buscando todos os usuarios salvo no banco
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users!"))
}

// busca um usuario salvo no banco
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user!"))
}

// edita um usuario salvo no banco
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user!"))
}

// deleta um usuario salvo no banco
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user!"))
}
