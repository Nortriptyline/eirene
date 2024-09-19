package service

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/domain/repository"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetOrCreateUser(token string) (*domain.User, error) {
	user := &domain.User{
		ClientToken: token,
	}

	user, err := s.userRepo.FindByToken(user)

	if err == gorm.ErrRecordNotFound {
		user, err = s.userRepo.Save(user)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
