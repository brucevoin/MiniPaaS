package model

import (
	"io/ioutil"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

type Template struct {
	Id         string
	Content    string
	Parameters []Parameter //模板定义的参数
}

// []Parameter用户输入的参数
func (u *Template) FillTemplate(parameters []Parameter) string {
	var parameterString = u.BuildParamString(parameters)
	ctx, instance := u.getTemplateInstance()
	param := ctx.CompileString(parameterString)
	if param.Err() != nil {
		panic(param.Err())
	}
	instance = instance.Fill(param)

	var buf []byte
	buf, err := yaml.Encode(instance.Value())
	if err != nil {
		panic(err)
	}
	originalStr := string(buf)

	var substring = trimParameter(originalStr)
	return substring
}

func (u Template) getTemplateInstance() (*cue.Context, cue.Value) {
	var file = "/Users/fhc/work/mini-paas/conf/application-template.cue"
	src, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var ctx = cuecontext.New()
	instance := ctx.CompileBytes(src)
	//从字符串读取
	//instance := ctx.CompileString(src)
	if instance.Err() != nil {
		panic(instance.Err())
	}
	return ctx, instance

}

func (u Template) BuildParamString(parameters []Parameter) string {
	var sb strings.Builder
	sb.WriteString("parameter: {\n")
	for _, p := range parameters {
		//TODO Check if Parameters is valid
		sb.WriteString("  " + p.ParameterKey + ": \"" + p.ParameterValue + "\"\n")
	}
	sb.WriteString("}\n")
	return sb.String()
}

func trimParameter(originalStr string) string {
	var splitter = "\"---\""
	idx := strings.Index(originalStr, splitter)
	var substring = originalStr[idx+len(splitter):]
	return substring

}
