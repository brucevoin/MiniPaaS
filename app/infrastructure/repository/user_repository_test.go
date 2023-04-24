package repository

import (
	"mini-paas/app/application/port"
	model "mini-paas/app/domain/model/user"
	"testing"
)

func TestAddUser(t *testing.T) {
	var user = model.User{}
	var userRepository port.UserRepository = &UserRepository{}
	userRepository.SaveUser(&user)

}
