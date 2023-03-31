package button

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
	table = "buttons"
)

var (
	fields = []string{
		"id",
		"name",
		"color",
		"shape",
		"details",
		"created_at",
		"updated_at",
	}
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
)

type Button struct {
	db *pgxpool.Pool
}

// New returns a new Button storage
func New(db *pgxpool.Pool) Button {
	return Button{db: db}
}

// Create creates a model.Button
func (b Button) Create(m *model.Button) error {
	_, err := b.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Name,
		m.Color,
		m.Shape,
		m.Details,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (b Button) GetByID(ID uuid.UUID) (model.Button, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := b.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return b.scanRow(row)
}

func (b Button) GetAll() (model.Buttons, error) {
	rows, err := b.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Buttons{}
	for rows.Next() {
		m, err := b.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (b Button) Update(m *model.Button) error {
	if len(m.Details) == 0 || m.Details == nil {
		p, err := b.GetByID(m.ID)
		if err != nil {
			return fmt.Errorf("the id does not exist: %d", m.ID)
		}
		m.Details = p.Details
	}

	res, err := b.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Name,
		m.Color,
		m.Shape,
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

func (b Button) scanRow(s pgx.Row) (model.Button, error) {
	m := model.Button{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Color,
		&m.Shape,
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
