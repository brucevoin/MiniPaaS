package model

type Component struct {
	Name       string
	Template   Template
	Parameters []UserParameter
}

func (u *Component) addParameter(parameters []UserParameter) {
	u.Parameters = append(u.Parameters, parameters...)
}

func (u *Component) Rendering() string {
	//TODO
	return u.Template.FillTemplate(u.Parameters)
}
