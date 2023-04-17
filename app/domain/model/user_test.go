package model

import (
	"fmt"
	"testing"
)

func TestChangePwd(t *testing.T) {
	var user = User{Name: "Alice", Passwd: "abcd1234"}
	var result = user.ChangePwd("Abcd12345")
	if result != nil {
		t.Errorf("user.ChangePwd(\"Abcd123456\") returned %s, expected %s", result.Error(), "新密码不能与旧密码相同")
	} else if user.Passwd != "Abcd12345" {
		t.Errorf("change Pwd filed")
	}
	fmt.Println(user)
}

func TestLogin(t *testing.T) {
	var user = User{Name: "Alice", Passwd: "abcd1234", IsValid: false}
	user.Login()
	fmt.Println(user)
}
