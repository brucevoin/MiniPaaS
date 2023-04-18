package repository

import (
	model "mini-paas/app/domain/model/user"
	"mini-paas/app/domain/port"
	"testing"
)

func TestAddUser(t *testing.T) {
	var user = model.User{}
	var userRepository port.UserRepository = &UserRepository{}
	userRepository.SaveUser(&user)

}
