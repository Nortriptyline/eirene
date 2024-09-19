package repository

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindByToken(*domain.User) (*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
}
