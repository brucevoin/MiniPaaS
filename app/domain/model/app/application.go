package model

import (
	"fmt"
	"strings"
)

type Application struct {
	Name        string
	Project     string
	Description string
	Template    Template
	Parameters  []Parameter
	Components  []Component
}

func (a *Application) AddComponent(component *Component) {
	a.Components = append(a.Components, *component)

}
func (a *Application) RemoveComponent(component *Component) {
	var index int = -1
	for i, elem := range a.Components {
		if elem.Name == component.Name {
			index = i
		}
	}
	if index == -1 {
		//TODO Error Can not find component
	} else {
		a.Components = append(a.Components[:index], a.Components[index+1:]...)
	}
}

func (a *Application) Rendering() string {

	componentsString := a.Components[0].Rendering()
	fmt.Printf(componentsString)
	lines := strings.Split(componentsString, "\n")
	componentsString = string("")
	for _, s := range lines {
		if s == "" {
			continue
		}
		s = "    " + s
		componentsString += s + "\n"
	}

	a.Parameters = []Parameter{
		{ParameterKey: "name", ParameterValue: "mytestaaaaaaa"},
		{ParameterKey: "url", ParameterValue: "http://172.16.120.13:12001/callback/collect/"},
	}
	a.Parameters = append(a.Parameters, Parameter{ParameterKey: "components", ParameterValue: `["component-1"]`})

	appString := a.Template.FillTemplate(a.Parameters)
	appString = strings.Replace(appString, "component-1", componentsString, -1)

	fmt.Printf(appString)
	//rendering basic information

	//rendering components
	//rendering policy
	//rendering workflow
	return ""
}
