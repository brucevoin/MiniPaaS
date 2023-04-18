package model

type Component struct {
	Name       string
	Template   Template
	Parameters []Parameter
}

func (u *Component) addParameter(parameters []Parameter) {
	u.Parameters = append(u.Parameters, parameters...)
}

func (u *Component) Rendering() string {
	return u.Template.FillTemplate(u.Parameters)
}
