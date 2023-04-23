package port

import model "mini-paas/app/domain/model/app"

type ApplicationRepository interface {
	Create(application *model.Application) (*model.Application, error)
}
