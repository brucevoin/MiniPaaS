package application

import (
	"fmt"
	"mini-paas/app/application/dto"
	"mini-paas/app/application/port"
	model "mini-paas/app/domain/model/app"
)

//Impl

type ApplicationServiceImpl struct {
	ApplicationRepository port.ApplicationRepository `inject:"ApplicationRepository"`
}

func NewApplicationService() *ApplicationServiceImpl {
	//applicationRepo := repository.NewApplicationRepository()
	return &ApplicationServiceImpl{
		//repo: applicationRepo,
	}
}

func (s *ApplicationServiceImpl) CreateApplication(application *dto.ApplicationDTO) (*dto.ApplicationDTO, error) {
	applicationDO := model.Application{Name: "test", Project: "testProject"}
	//DO something with application
	if applicationDO.Components == nil {
		fmt.Println("the application Components is nil")
		//panic(application)
	}
	if s.ApplicationRepository == nil {
		fmt.Println("the repo is nil")
		panic(s.ApplicationRepository)
	}

	s.ApplicationRepository.Create(&applicationDO)
	fmt.Println("Application created successfully")
	return nil, nil
}

func (s *ApplicationServiceImpl) GetApplication(id string) (*dto.ApplicationDTO, error) {
	return nil, nil
}
func (s *ApplicationServiceImpl) UpdateApplication(application *dto.ApplicationDTO) (*dto.ApplicationDTO, error) {
	return nil, nil
}
func (s *ApplicationServiceImpl) DeleteApplication(id string) error {
	return nil
}
func (s *ApplicationServiceImpl) ListApplications() ([]dto.ApplicationDTO, error) {
	return nil, nil
}
func (s *ApplicationServiceImpl) RunApplication(application *dto.ApplicationDTO) error {
	return nil
}
