package port

type ApplicationDTO struct {
}

type ApplicationService interface {
	GetApplication(id string) (*ApplicationDTO, error)
	CreateApplication(application *ApplicationDTO) (*ApplicationDTO, error)
	UpdateApplication(application *ApplicationDTO) (*ApplicationDTO, error)
	DeleteApplication(id string) error
	ListApplications() ([]ApplicationDTO, error)
	RunApplication(application *ApplicationDTO) error
}
