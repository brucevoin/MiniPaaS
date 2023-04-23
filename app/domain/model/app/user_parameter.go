package model

type UserParameter struct {
	ParameterKey   string
	ParameterValue string
	ParameterType  parameterType
}

type parameterType int

const (
	PARAMETER_TYPE_STRING  parameterType = 0
	PARAMETER_TYPE_NUMBER  parameterType = 1
	PARAMETER_TYPE_BOOLEAN parameterType = 2
	PARAMETER_TYPE_LIST    parameterType = 3
)
