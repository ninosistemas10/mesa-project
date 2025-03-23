package model

import (
	"github.com/google/uuid"
)

// PromocionProducto representa la relación entre promociones y productos
type PromocionProducto struct {
	ID          uuid.UUID `json:"id"`
	PromocionID uuid.UUID `json:"promocion_id"`
	ProductoID  uuid.UUID `json:"producto_id"`
	Cantidad    int       `json:"cantidad"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at,omitempty"`
}

// HasID verifica si la relación tiene un ID asignado
func (pp PromocionProducto) HasID() bool {
	return pp.ID != uuid.Nil
}

// PromocionProductoList representa una lista de relaciones entre promociones y productos
type PromocionProductoList []PromocionProducto

// IsEmpty verifica si la lista está vacía
// func (pp PromocionProductoList) IsEmpty() bool {
// 	return len(pp) == 0
// }
