package promocionproducto

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/delivery/infrastructure/postgres"
	"github.com/ninosistemas10/delivery/model"
)

const table = "promocionProducto"

var fields = []string{
	"id",
	"promocion_id",
	"producto_id",
	"cantidad",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type PromocionProducto struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) PromocionProducto {
	return PromocionProducto{}
}

func (pp PromocionProducto) Create(m *model.PromocionProducto) error {
	_, err := pp.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.PromocionID,
		m.ProductoID,
		m.Cantidad,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (pp PromocionProducto) Update(m *model.PromocionProducto) error {
	_, err := pp.db.Exec(
		context.Background(),
		psqlUpdate,
		m.ProductoID,
		m.Cantidad,
		m.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (pp PromocionProducto) Delete(ID uuid.UUID) error {
	_, err := pp.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (pp PromocionProducto) GetByID(ID uuid.UUID) (model.PromocionProducto, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := pp.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return pp.scanRow(row)
}

func (pp PromocionProducto) scanRow(s pgx.Row) (model.PromocionProducto, error) {
	m := model.PromocionProducto{}

	updateAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.PromocionID,
		&m.ProductoID,
		&m.Cantidad,
		&m.CreatedAt,
		&updateAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updateAtNull.Int64
	return m, nil
}
