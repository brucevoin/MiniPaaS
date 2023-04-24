package application

import "mini-paas/app/application/dto"

type UserServiceImpl struct {
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) Create(user *dto.UserDTO) (*dto.UserDTO, error) {
	return user, nil
}
