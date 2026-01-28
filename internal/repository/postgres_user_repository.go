package repository

import (
	"database/sql"

	"github.com/Akakazkz/go-task-manager-api/internal/model"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) List() ([]*model.User, error) {
	rows, err := r.db.Query(`
		SELECT id, email, password, role, created_at
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var u model.User
		if err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.Password,
			&u.Role,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *PostgresUserRepository) ExistsByEmail(email string) bool {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	_ = r.db.QueryRow(query, email).Scan(&exists)
	return exists
}

func (r *PostgresUserRepository) GetByEmail(email string) (*model.User, error) {
	var u model.User

	query := `
		SELECT id, email, password, role, created_at
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PostgresUserRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users (email, password, role, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
	).Scan(&user.ID)
}
