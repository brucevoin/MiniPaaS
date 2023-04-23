package repository

import (
	"fmt"
	model "mini-paas/app/domain/model/app"
)

type ApplicationRepository struct {
}

func NewApplicationRepository() *ApplicationRepository {
	fmt.Println("creating applicationRepository")
	return &ApplicationRepository{}
}

func (r *ApplicationRepository) Create(application *model.Application) (*model.Application, error) {
	fmt.Println("ApplicationRepository Create")
	return application, nil
}
