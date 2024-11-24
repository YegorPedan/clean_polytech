package repository

import "clean-polytech/internal/domain/model"

type UserRepository interface {
	Save(student *model.User) error
	GetUser(id int) (*model.User, error)
}
