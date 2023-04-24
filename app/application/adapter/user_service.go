package adapter

import "mini-paas/app/application/dto"

//Implementation of interface port.UserService

type UserServiceImpl struct {
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) Create(user *dto.UserDTO) (*dto.UserDTO, error) {
	return user, nil
}
