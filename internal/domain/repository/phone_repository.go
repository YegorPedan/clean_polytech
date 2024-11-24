package repository

import "clean-polytech/internal/domain/model"

type PhoneRepository interface {
	SavePhone(phone *model.Smartphone) error
}
