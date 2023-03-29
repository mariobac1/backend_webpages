package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/mariobac1/backend_webpages/infrastructure/postgres"
	"github.com/mariobac1/backend_webpages/model"
)

const (
	table = "users"
)

var (
	fields = []string{
		"id",
		"name",
		"email",
		"password",
		"avatar",
		"details",
		"created_at",
		"updated_at",
	}
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
)

type User struct {
	db *pgxpool.Pool
}

// New returns a new User storage
func New(db *pgxpool.Pool) User {
	return User{db: db}
}

// Create creates a model.User
func (u User) Create(m *model.User) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Name,
		m.Email,
		m.Password,
		m.Avatar,
		m.Details,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u User) GetByID(ID uuid.UUID) (model.User, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := u.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return u.scanRow(row, false)
}

func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlGetAll + " WHERE email = $1"
	row := u.db.QueryRow(
		context.Background(),
		query,
		email,
	)

	return u.scanRow(row, true)
}

func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := model.Users{}
	for rows.Next() {
		m, err := u.scanRow(rows, false)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) Update(m *model.User) error {
	// 	if len(m.Details) == 0 || m.Details == nil {
	// 		p, err := u.GetByID(m.ID)
	// 		if err != nil {
	// 			return fmt.Errorf("the id does not exist: %d", m.ID)
	// 		}
	// 		m.Details = p.Details
	// 	}

	// 	res, err := u.db.Exec(
	// 		context.Background(),
	// 		noAdminPsqlUpdate,
	// 		m.Password,
	// 		m.CreatedAt,
	// 		m.ID,
	// 	)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if res.RowsAffected() == 0 {
	// 		return fmt.Errorf("the id does not exist: %d", m.ID)
	// 	}

	return nil

}

// func (u User) AdminUpdate(m *model.User) error {
// 	if len(m.Details) == 0 || m.Details == nil {
// 		p, err := u.GetByID(m.ID)
// 		if err != nil {
// 			return fmt.Errorf("the id does not exist: %d", m.ID)
// 		}
// 		m.Details = p.Details
// 	}

// 	res, err := u.db.Exec(
// 		context.Background(),
// 		psqlUpdate,
// 		m.Name,
// 		m.Email,
// 		m.Password,
// 		m.Avatar,
// 		m.Details,
// 		m.CreatedAt,
// 		m.ID,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if res.RowsAffected() == 0 {
// 		return fmt.Errorf("the id does not exist: %d", m.ID)
// 	}

// 	return nil
// }

func (u User) scanRow(s pgx.Row, withPassword bool) (model.User, error) {
	m := model.User{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Email,
		&m.Password,
		&m.Avatar,
		&m.Details,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Int64

	if !withPassword {
		m.Password = ""
	}

	return m, nil
}
