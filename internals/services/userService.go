package services

import (
	"bugounty/internals/dto"
	"bugounty/internals/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (h *UserService) CreateUser(userData dto.UserDto) {

}
