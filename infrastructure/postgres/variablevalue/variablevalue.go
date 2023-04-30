package variablevalue

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
	table = "variablevalues"
)

var (
	fields = []string{
		"id",
		"name",
		"title",
		"paragraph",
		"color",
		"bgcolor",
		"font",
		"icon",
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

type VariableValue struct {
	db *pgxpool.Pool
}

// New returns a new VariableValue storage
func New(db *pgxpool.Pool) VariableValue {
	return VariableValue{db: db}
}

// Create creates a model.VariableValue
func (i VariableValue) Create(m *model.VariableValue) error {
	_, err := i.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Name,
		m.Title,
		m.Paragraph,
		m.Color,
		m.BgColor,
		m.Font,
		m.Icon,
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

func (i VariableValue) GetByID(ID uuid.UUID) (model.VariableValue, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := i.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return i.scanRow(row)
}

func (i VariableValue) GetAll() (model.VariableValues, error) {
	rows, err := i.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.VariableValues{}
	for rows.Next() {
		m, err := i.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (i VariableValue) Update(m *model.VariableValue) error {
	if len(m.Details) == 0 || m.Details == nil {
		im, err := i.GetByID(m.ID)
		if err != nil {
			return fmt.Errorf("the id does not exist: %d", m.ID)
		}
		m.Details = im.Details
	}

	res, err := i.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Name,
		m.Title,
		m.Paragraph,
		m.Color,
		m.BgColor,
		m.Font,
		m.Icon,
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

func (i VariableValue) scanRow(s pgx.Row) (model.VariableValue, error) {
	m := model.VariableValue{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Title,
		&m.Paragraph,
		&m.Color,
		&m.BgColor,
		&m.Font,
		&m.Icon,
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
