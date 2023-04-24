package port

import "mini-paas/app/application/dto"

type UserService interface {
	Create(user *dto.UserDTO) (*dto.UserDTO, error)
}
