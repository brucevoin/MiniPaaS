package controllers

import (
	"fmt"
	"mini-paas/app"
	"mini-paas/app/application/port"

	"github.com/revel/revel"
)

type ApplicationController struct {
	*revel.Controller
}

func (c ApplicationController) Applications() revel.Result {
	applicationService := app.IOCContainer.GetResource("applicationService")
	s, ok := applicationService.(port.ApplicationService)
	if !ok {
		fmt.Println("Get application service from IOCContainer failed")
	}
	s.CreateApplication(nil)
	fmt.Println("s.createApplication success")
	users := []string{"Alice", "Bob", "Charlie"}
	return c.RenderJSON(users)
}
