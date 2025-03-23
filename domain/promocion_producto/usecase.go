package promocionproducto

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)

type PromocionProducto struct {
	storage Storage
}

func New(s Storage) PromocionProducto {
	return PromocionProducto{storage: s}
}

func (pp PromocionProducto) Create(m *model.PromocionProducto) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	m.ID = ID

	m.CreatedAt = time.Now().Unix()

	err = pp.storage.Create(m)
	if err != nil {
		return err
	}

	return nil
}

func (pp PromocionProducto) Update(m *model.PromocionProducto) error {

	if !m.HasID() {
		return fmt.Errorf("Update HasID")
	}

	m.UpdatedAt = time.Now().Unix()

	err := pp.storage.Update(m)
	if err != nil {
		return err
	}

	return nil
}

func (pp PromocionProducto) Delete(ID uuid.UUID) error {

	err := pp.storage.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (pp PromocionProducto) GetByID(ID uuid.UUID) (model.PromocionProducto, error) {

	promocionProducto, err := pp.storage.GetByID(ID)
	if err != nil {
		return model.PromocionProducto{}, fmt.Errorf("promocionProducto: %w", err)
	}

	return promocionProducto, nil
}
