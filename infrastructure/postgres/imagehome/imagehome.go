package imagehome

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
	table = "imagehome"
)

var (
	fields = []string{
		"id",
		"name",
		"color",
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

type ImageHome struct {
	db *pgxpool.Pool
}

// New returns a new ImageHome storage
func New(db *pgxpool.Pool) ImageHome {
	return ImageHome{db: db}
}

// Create creates a model.ImageHome
func (i ImageHome) Create(m *model.ImageHome) error {
	_, err := i.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Name,
		m.Color,
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

func (i ImageHome) GetByID(ID uuid.UUID) (model.ImageHome, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := i.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return i.scanRow(row)
}

func (i ImageHome) GetAll() (model.ImageHomes, error) {
	rows, err := i.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.ImageHomes{}
	for rows.Next() {
		m, err := i.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (i ImageHome) Update(m *model.ImageHome) error {
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
		m.Color,
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

func (i ImageHome) scanRow(s pgx.Row) (model.ImageHome, error) {
	m := model.ImageHome{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Color,
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
