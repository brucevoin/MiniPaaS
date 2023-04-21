package model

import (
	"fmt"
	"io/ioutil"
	"testing"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

func TestCUEFill(t *testing.T) {
	var ctx = cuecontext.New()
	var file = "/Users/fhc/work/MiniPaaS/conf/template.cue"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	templateString := string(data)
	template := ctx.CompileString(templateString)

	parameterString := `parameter: {
    name: ["fff"]
    age: [123,345]
}`

	parameterString = "parameter: {\n  age: [123,345]\n  name: [\"component-1\"]\n}\n"
	parameter := ctx.CompileString(parameterString)

	result := template.Fill(parameter)

	var buf []byte
	buf, err2 := yaml.Encode(result.Value())
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(buf))
}

func TestRenderingApp(t *testing.T) {
	compparameters := []Parameter{
		{ParameterKey: "name", ParameterValue: "mytestaaaaaaa"},
		{ParameterKey: "image", ParameterValue: "nginx:1.21.3"},
		{ParameterKey: "memory", ParameterValue: "512Mi"},
		{ParameterKey: "cpu", ParameterValue: "0.5"},
		{ParameterKey: "port", ParameterValue: "[8088, 8099]"},
		{ParameterKey: "url", ParameterValue: "http://172.16.120.13:12001/callback/collect/"},
	}
	var compfile = "/Users/fhc/work/MiniPaaS/conf/component-template.cue"
	compsrc, err := ioutil.ReadFile(compfile)
	if err != nil {
		panic(err)
	}
	var comptemplate = Template{Id: "123", Content: string(compsrc), Parameters: compparameters}
	component := Component{Name: "comp1", Template: comptemplate, Parameters: compparameters}

	appparameters := []Parameter{
		{ParameterKey: "name", ParameterValue: "mytestaaaaaaa"},
		{ParameterKey: "url", ParameterValue: "http://172.16.120.13:12001/callback/collect/"},
		{ParameterKey: "components", ParameterValue: "comp1"},
	}

	var file = "/Users/fhc/work/MiniPaaS/conf/application-template.cue"
	appsrc, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var apptemplate = Template{Id: "123", Content: string(appsrc), Parameters: appparameters}

	app := Application{Name: "testApp", Project: "pro1", Description: "", Template: apptemplate, Parameters: appparameters, Components: []Component{component}}

	app.Rendering()

}
func TestRenderingComponent(t *testing.T) {
	parameters := []Parameter{
		{ParameterKey: "name", ParameterValue: "mytestaaaaaaa"},
		{ParameterKey: "image", ParameterValue: "nginx:1.21.3"},
		{ParameterKey: "memory", ParameterValue: "512Mi"},
		{ParameterKey: "cpu", ParameterValue: "0.5"},
		{ParameterKey: "port", ParameterValue: "[8088, 8099]"},
		{ParameterKey: "url", ParameterValue: "http://172.16.120.13:12001/callback/collect/"},
	}

	var file = "/Users/fhc/work/MiniPaaS/conf/component-template.cue"
	src, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var template = Template{Id: "123", Content: string(src), Parameters: parameters}
	component := Component{Name: "comp1", Template: template, Parameters: parameters}
	result := component.Rendering()
	fmt.Println(result)

}

func TestBuildParamaters(t *testing.T) {
	parameters := []Parameter{
		{ParameterKey: "name", ParameterValue: "mytestaaaaaaa"},
		{ParameterKey: "image", ParameterValue: "nginx:1.21.3"},
		{ParameterKey: "memory", ParameterValue: "512Mi"},
		{ParameterKey: "cpu", ParameterValue: "0.5"},
		{ParameterKey: "port", ParameterValue: "[8088, 8099]"},
		{ParameterKey: "url", ParameterValue: "http://172.16.120.13:12001/callback/collect/"},
	}
	var result = Template.BuildParamString(Template{}, parameters)
	fmt.Printf(result)
}
