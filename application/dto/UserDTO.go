package dto

import (
	"github.com/Nortriptyline/Eirene/domain"
)

type UserDTO struct {
	ID          string `json:"id"`
	ClientToken string `json:"client_token"`
}

type UserLightDTO struct {
	ID          string `json:"id"`
	ClientToken string `json:"client_token"`
}

func ToUserDTO(user *domain.User) *UserDTO {
	return &UserDTO{
		ID:          user.ID.String(),
		ClientToken: user.ClientToken,
	}
}

func ToUserLightDTO(User *domain.User) *UserLightDTO {
	return &UserLightDTO{
		ID:          User.ID.String(),
		ClientToken: User.ClientToken,
	}
}
