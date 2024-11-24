package repository

import "clean-polytech/internal/domain/model"

type UserRepository interface {
	SaveUser(student *model.User) error
	GetAllUsers() ([]*model.User, error)
}
