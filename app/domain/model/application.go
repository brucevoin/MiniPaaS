package model

type Application struct {
	Name        string
	Project     string
	Description string
	Template    Template
	Parameters  []Parameter
	Components  []Component
}
