package promocion

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/delivery/infrastructure/postgres"
	"github.com/ninosistemas10/delivery/model"
)

const table = "promocion"

var field = []string{
	"id",
	"nombre",
	"description",
	"image",
	"precio",
	"features",
	"categoria",
	"stock_disponible",
	"activo",
	"created_at",
	"updated_at",
}

var (
	psqlInsert      = postgres.BuildSQLInsert(table, field)
	psqlUpdate      = postgres.BuildSQLUpdateByID(table, field)
	psqlDelete      = postgres.BuildSQLDelete(table)
	psqlGetAll      = postgres.BuildSQLSelect(table, field)
	psqlUpdateImage = `UPDATE promocion SET image = $1, updated_at = $2 WHERE id = $3` // Nueva consulta
)

type Promocion struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Promocion {
	return Promocion{db: db}
}

func (p Promocion) Create(m *model.Promocion) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Nombre,
		m.Description,
		m.Image,
		m.Precio,
		m.Features,
		m.Categoria,
		m.StockDisponible,
		m.Activo,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Promocion) Update(m *model.Promocion) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Nombre,
		m.Description,
		m.Image,
		m.Precio,
		m.Features,
		m.Categoria,
		m.StockDisponible,
		m.Activo,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c Promocion) UpdateImage(ID uuid.UUID, imagePath string) error {
	_, err := c.db.Exec(
		context.Background(),
		psqlUpdateImage,
		imagePath,
		time.Now().Unix(), // Registrar hora de actualización
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Promocion) Delete(ID uuid.UUID) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p Promocion) GetByID(ID uuid.UUID) (model.Promocion, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := p.db.QueryRow(
		context.Background(),
		query,
		ID,
	)
	return p.scanRow(row)
}

func (p Promocion) GetAll() (model.Promociones, error) {
	rows, err := p.db.Query(
		context.Background(),
		psqlGetAll,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Promociones
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
}

func (p Promocion) scanRow(s pgx.Row) (model.Promocion, error) {
	m := model.Promocion{}
	updatedNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Nombre,
		&m.Description,
		&m.Image,
		&m.Precio,
		&m.Features,
		&m.Categoria,
		&m.StockDisponible,
		&m.Activo,
		&m.CreatedAt,
		&updatedNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedNull.Int64
	return m, nil
}
