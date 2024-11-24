package user

import (
	"clean-polytech/internal/domain/model"
	"clean-polytech/internal/domain/repository"
	"context"
	"github.com/google/uuid"
	"time"
)

type SaveUser struct {
	userRepo  repository.UserRepository
	phoneRepo repository.PhoneRepository
}

func SaveNewUser(userRepository repository.UserRepository, phoneRepository repository.PhoneRepository) *SaveUser {
	return &SaveUser{
		userRepo:  userRepository,
		phoneRepo: phoneRepository,
	}
}

func (s *SaveUser) Execute(ctx context.Context, name string, familyName string, phoneModel string) error {
	phoneID := uuid.New().String()
	userID := uuid.New().String()
	phone := &model.Smartphone{
		ID:             phoneID,
		Model:          phoneModel,
		Charge:         charge,
		ConnectionTime: time.Now(),
		UserID:         userID,
	}
	if err := s.phoneRepo.SavePhone(phone); err != nil {
		return err
	}
	user := &model.User{
		ID:         userID,
		Name:       name,
		FamilyName: familyName,
		PhoneID:    phoneID,
		Phone:      phone,
	}

	return s.userRepo.SaveUser(user)
}