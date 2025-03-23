package promocionproducto

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)

type UseCase interface {
	Create(m *model.PromocionProducto) error
	Update(m *model.PromocionProducto) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.PromocionProducto, error)
	//GetAll() (model.PromocionProductoList, error)
}

type Storage interface {
	Create(m *model.PromocionProducto) error
	Update(m *model.PromocionProducto) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.PromocionProducto, error)
	// GetAll() (model.PromocionProductoList, error)
}
