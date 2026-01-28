package service

import (
	"strings"
	"time"

	"github.com/Akakazkz/go-task-manager-api/internal/model"
	"github.com/Akakazkz/go-task-manager-api/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) List() ([]*model.User, error) {
	return s.repo.List()
}

func (s *userService) Create(email, password string) (*model.User, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, ErrInvalidInput
	}

	if s.repo.ExistsByEmail(email) {
		return nil, ErrUserExists
	}
	user := &model.User{
		Email:     email,
		Password:  password,
		Role:      model.RoleUser,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
