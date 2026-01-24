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

func (r *PostgresUserRepository) ExistsByEmail(email string) bool {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	_ = r.db.QueryRow(query, email).Scan(&exists)
	return exists
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
