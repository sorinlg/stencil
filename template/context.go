package template

import (
	"os"
	"strings"
)

type Context struct {
	Env map[string]string
}

func NewContext() *Context {
	c := &Context{}
	c.Env = c.env()

	return c
}

func (c *Context) env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}
