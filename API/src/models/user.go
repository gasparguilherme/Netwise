package models

import "time"

// user representa um usuario utilizando a rede social
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedIn time.Time `json:"CreatedIn"`
}
