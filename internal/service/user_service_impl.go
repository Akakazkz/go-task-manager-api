package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Akakazkz/go-task-manager-api/internal/model"
)

type userService struct {
	users  map[string]*model.User
	nextID int64
}

func NewUserService() UserService {
	return &userService{
		users:  make(map[string]*model.User),
		nextID: 1,
	}
}

func (s *userService) Create(email, password string) (*model.User, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	if _, exists := s.users[email]; exists {
		return nil, errors.New("user already exists")
	}

	user := &model.User{
		ID:        s.nextID,
		Email:     email,
		Password:  password,
		Role:      model.RoleUser,
		CreatedAt: time.Now(),
	}

	s.users[email] = user
	s.nextID++

	return user, nil
}
