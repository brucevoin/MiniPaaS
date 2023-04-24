package application

import (
	"fmt"
	model "mini-paas/app/domain/model/app"
	port "mini-paas/app/domain/port"
)

type ApplicationService struct {
	ApplicationRepository port.ApplicationRepository `inject:"ApplicationRepository"`
}

func (a ApplicationService) NewApplicationService() *ApplicationService {
	//applicationRepo := repository.NewApplicationRepository()
	return &ApplicationService{
		//repo: applicationRepo,
	}
}

func (s *ApplicationService) CreateApplication(application *port.ApplicationDTO) (*port.ApplicationDTO, error) {
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

func (s *ApplicationService) GetApplication(id string) (*port.ApplicationDTO, error) {
	return nil, nil
}
func (s *ApplicationService) UpdateApplication(application *port.ApplicationDTO) (*port.ApplicationDTO, error) {
	return nil, nil
}
func (s *ApplicationService) DeleteApplication(id string) error {
	return nil
}
func (s *ApplicationService) ListApplications() ([]port.ApplicationDTO, error) {
	return nil, nil
}
func (s *ApplicationService) RunApplication(application *port.ApplicationDTO) error {
	return nil
}
