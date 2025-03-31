package repository

import "github.com/tryuuu/go-task-api/internal/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
