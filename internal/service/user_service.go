package service

import "github.com/Akakazkz/go-task-manager-api/internal/model"

type UserService interface {
	Create(email, password string) (*model.User, error)
	List() ([]*model.User, error)
	Login(email, password string) (string, error)
}
