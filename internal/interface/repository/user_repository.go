package repository

import "github.com/tryuuu/go-task-api/internal/domain/model"

type UserRepository interface {
	Create(user *model.User) error
}
