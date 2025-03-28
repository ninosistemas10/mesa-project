package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	ID 		    uuid.UUID       	`json:"id"`
	Nombre		string				`json:"nombre"`
	Email     	string          	`json:"email"`
	Password  	string          	`json:"password"`
	IsAdmin   	bool            	`json:"is_admin"`
	Images		json.RawMessage		`json:"images"`
	Details   	json.RawMessage		`json:"details"`
	CreatedAt 	int64   	        `json:"created_at"`
	UpdatedAt 	int64	          	`json:"updated_at"`
}

type Users []User
