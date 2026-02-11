package repositories

import (
	"database/sql"

	"github.com/gasparguilherme/Netwise/api/src/models"
)

// usuarios representa um repositorio de ususarios
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository cria um repositorio de usuarios
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// create insere um usuario no banco de dados
func (u UserRepository) Create(user models.User) (uint64, error) {
	return 0, nil
}
