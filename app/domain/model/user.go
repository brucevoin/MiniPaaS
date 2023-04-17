package model

import "errors"

type User struct {
	Name    string
	Passwd  string
	Roles   []Role
	IsValid bool
}

func (u *User) ChangePwd(newPwd string) error {
	if u.Passwd == newPwd {
		return errors.New("新密码不能与旧密码相同")
	}
	u.Passwd = newPwd
	return nil
}

func (u User) Login() error {
	u.IsValid = true
	return nil
}
