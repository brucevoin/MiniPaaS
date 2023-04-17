package port

import "mini-paas/app/domain/model"

type UserRepository interface {
	SaveUser(user *model.User)
	UpdateUser(user *model.User)
}
