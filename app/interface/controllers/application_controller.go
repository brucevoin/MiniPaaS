package controllers

import (
	"fmt"
	"mini-paas/app"
	"mini-paas/app/application"

	"github.com/revel/revel"
)

type ApplicationController struct {
	*revel.Controller
}

func (c ApplicationController) Applications() revel.Result {
	applicationService := app.IOCContainer.GetResource("applicationService")
	s, ok := applicationService.(application.ApplicationService)
	if !ok {
		fmt.Println("类型断言失败")
	}
	s.CreateApplication(nil)
	fmt.Println("s.createApplication success")
	users := []string{"Alice", "Bob", "Charlie"}
	return c.RenderJSON(users)
}
