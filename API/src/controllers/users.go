package controllers

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/Netwise/api/src/database"
	"github.com/gasparguilherme/Netwise/api/src/models"
)

// criar usuario insere um usuario no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		slog.Error("unable to unterpret JSON", "error", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return

	}

	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

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
