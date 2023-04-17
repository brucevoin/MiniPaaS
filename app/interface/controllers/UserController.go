package controllers

import (
	domain "mini-paas/app/domain/model"

	"github.com/revel/revel"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) Index() revel.Result {
	users := []string{"Alice", "Bob", "Charlie"}
	return c.RenderJSON(users)
}

func (c UserController) Show(id string) revel.Result {
	// user := struct {
	// 	ID   string `json:"id"`
	// 	Name string `json:"name"`
	// }{
	// 	ID:   id,
	// 	Name: "Alice",
	// }
	var user = domain.User{Name: "Alice", Passwd: "abcd1234"}
	user.ChangePwd("Abcd1234")
	return c.RenderJSON(user)
}
