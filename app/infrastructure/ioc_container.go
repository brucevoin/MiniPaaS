package infrastructure

import "reflect"

type Container struct {
	resources map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		resources: make(map[string]interface{}),
	}
}

func (c *Container) Register(key string, value interface{}) {
	c.resources[key] = value
}

func (c *Container) GetResource(key string) interface{} {
	return c.resources[key]
}

func (c *Container) Inject(object interface{}) {
	v := reflect.ValueOf(object).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		tagName := t.Field(i).Tag.Get("inject")
		if tagName == "" {
			continue
		}

		resource := c.GetResource(tagName)
		if needInject(resource) {
			c.Inject(resource)
		}
		field.Set(reflect.ValueOf(resource))
	}
}
func needInject(object interface{}) bool {
	v := reflect.ValueOf(object).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		tagName := t.Field(i).Tag.Get("inject")
		if tagName != "" {
			return true
		}
	}

	return false
}

// 只支持一级依赖自动注入
func (c *Container) InjectAll() error {
	for _, v := range c.resources {
		if needInject(v) {
			c.Inject(v)
		}
	}
	return nil
}
