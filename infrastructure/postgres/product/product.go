package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/mariobac1/backend_webpages/infrastructure/postgres"
	"github.com/mariobac1/backend_webpages/model"
)

const (
	table = "products"
)

var (
	fields = []string{
		"id",
		"name",
		"price",
		"description",
		"details",
		"created_at",
		"updated_at",
	}
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
)

type Product struct {
	db *pgxpool.Pool
}

// New returns a new Product storage
func New(db *pgxpool.Pool) Product {
	return Product{db: db}
}

// Create creates a model.Product
func (p Product) Create(m *model.Product) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Name,
		m.Price,
		m.Description,
		m.Details,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := p.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return p.scanRow(row)
}

func (p Product) GetAll() (model.Products, error) {
	rows, err := p.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Products{}
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (p Product) Update(m *model.Product) error {
	if len(m.Details) == 0 || m.Details == nil {
		p, err := p.GetByID(m.ID)
		if err != nil {
			return fmt.Errorf("the id does not exist: %d", m.ID)
		}
		m.Details = p.Details
	}

	res, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Name,
		m.Price,
		m.Description,
		m.Details,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("the id does not exist: %d", m.ID)
	}

	return nil
}

func (p Product) Delete(ID uuid.UUID) error {
	res, err := p.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)

	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("the id does not exist: %d", ID)
	}

	return nil

}

func (p Product) scanRow(s pgx.Row) (model.Product, error) {
	m := model.Product{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Price,
		&m.Description,
		&m.Details,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Int64

	return m, nil
}
