package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Promocion struct {
	ID              uuid.UUID       `json:"id"`
	Nombre          string          `json:"nombre"`
	Slug            string          `json:"slug"`
	Description     string          `json:"description"`
	Image           string          `json:"image"`
	Precio          float64         `json:"precio"`
	Features        json.RawMessage `json:"features"`
	Categoria       string          `json:"categoria"`
	StockDisponible int             `json:"stock_disponible"`
	Activo          bool            `json:"activo"`
	CreatedAt       int64           `json:"created_at"`
	UpdatedAt       int64           `json:"updated_at"`
}

func (p Promocion) HasID() bool {
	return p.ID != uuid.Nil
}

type Promociones []Promocion

func (p Promociones) IdEmpty() bool {
	return len(p) == 0
}
