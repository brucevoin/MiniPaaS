package port

import "mini-paas/app/application/dto"

type ApplicationInterface interface {
	GetApplication(name string) (*dto.ApplicationDTO, error)
}
