package app

import (
	"fmt"
	"mini-paas/app/application"
	"mini-paas/app/infrastructure/repository"
)

type Context struct {
	Resources map[string]interface{}
}

func NewContext() *Context {
	return &Context{
		Resources: make(map[string]interface{}),
	}
}

func (c *Context) Register(key string, value interface{}) {
	c.Resources[key] = value
}

func (c *Context) GetResource(key string) interface{} {
	return c.Resources[key]
}

func (c *Context) Init() {
	fmt.Println("Init application resources ...")
	applicationService := application.ApplicationService.NewApplicationService(application.ApplicationService{})
	c.Register("applicationService", applicationService)

	applicationRepository := repository.NewApplicationRepository()
	c.Register("applicationRepository", applicationRepository)
	fmt.Println("Init application resources end")
}
