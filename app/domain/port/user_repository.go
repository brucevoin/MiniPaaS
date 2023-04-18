package port

import model "mini-paas/app/domain/model/user"

type UserRepository interface {
	SaveUser(user *model.User)
	UpdateUser(user *model.User)
}
