package model

import "testing"

func TestAddParameters(t *testing.T) {
	parameters := []Parameter{
		{ParameterKey: "name", ParameterValue: "mytestaaaaaaa"},
		{ParameterKey: "image", ParameterValue: "nginx:1.21.3"},
		{ParameterKey: "memory", ParameterValue: "512Mi"},
		{ParameterKey: "cpu", ParameterValue: "0.5"},
		{ParameterKey: "port", ParameterValue: "[8088, 8099]"},
		{ParameterKey: "url", ParameterValue: "http://172.16.120.13:12001/callback/collect/"},
	}
	var template = Template{Id: "123", Content: "xxx", Parameters: parameters}
	var component = Component{Name: "component1", Template: template, Parameters: parameters}
	//component.addParameter(parameters)
	component.Rendering()
}
