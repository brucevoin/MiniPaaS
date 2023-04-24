package port

import "mini-paas/app/application/dto"

type ApplicationService interface {
	GetApplication(id string) (*dto.ApplicationDTO, error)
	CreateApplication(application *dto.ApplicationDTO) (*dto.ApplicationDTO, error)
	UpdateApplication(application *dto.ApplicationDTO) (*dto.ApplicationDTO, error)
	DeleteApplication(id string) error
	ListApplications() ([]dto.ApplicationDTO, error)
	RunApplication(application *dto.ApplicationDTO) error
}
