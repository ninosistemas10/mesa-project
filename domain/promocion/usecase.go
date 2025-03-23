package promocion

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/delivery/model"
)

type Promocion struct {
	storage Storage
}

func New(s Storage) Promocion {
	return Promocion{storage: s}
}

func (p Promocion) Create(m *model.Promocion) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.UUID.NewUUID()", err)
	}

	if len(m.Categoria) == 0 {
		m.Categoria = "General"
	}

	m.ID = ID
	if len(m.Image) == 0 {
		m.Image = ""
	}

	if m.Features == nil {
		m.Features = []byte("{}")
	}

	m.CreatedAt = time.Now().Unix()

	err = p.storage.Create(m)
	if err != nil {
		return err
	}

	return nil

}

func (p Promocion) Update(m *model.Promocion) error {
	if m.HasID() {
		return fmt.Errorf("Update HasID")
	}

	if len(m.Image) == 0 {
		m.Image = ""
	}
	m.UpdatedAt = time.Now().Unix()

	err := p.storage.Update(m)
	if err != nil {
		return err
	}

	return nil
}

func (c Promocion) UpdateImage(ID uuid.UUID, imagePath string) error {
	// Verificar si el ID es válido
	if ID == uuid.Nil {
		return fmt.Errorf("invalid ID")
	}

	// Intentar actualizar la imagen en la base de datos
	err := c.storage.UpdateImage(ID, imagePath)
	if err != nil {
		return err
	}

	return nil
}

func (p Promocion) Delete(ID uuid.UUID) error {
	err := p.storage.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (p Promocion) GetByID(ID uuid.UUID) (model.Promocion, error) {
	promocion, err := p.storage.GetByID(ID)
	if err != nil {
		return model.Promocion{}, fmt.Errorf("promocion: %w", err)
	}

	return promocion, nil
}

func (p Promocion) GetAll() (model.Promociones, error) {
	promociones, err := p.storage.GetAll()
	if err != nil {
		return model.Promociones{}, fmt.Errorf("promociones: %w", err)
	}
	if len(promociones) == 0 {
		return model.Promociones{}, nil
	}

	return promociones, nil
}
