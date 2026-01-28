package repository

import "github.com/Akakazkz/go-task-manager-api/internal/model"

type UserRepository interface {
	Create(user *model.User) error
	ExistsByEmail(email string) bool
	List() ([]*model.User, error)
}
