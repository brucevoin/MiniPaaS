package model

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

func (a *Application) Rendering() string{
	//rendering basic information
	//rendering components
	//rendering workflow
	return ""
}
