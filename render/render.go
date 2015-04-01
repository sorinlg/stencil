package render

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Context struct {
}

func (c *Context) Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func Render(template_path string, destination_path string) {
	// load custom template functions
	funcMap := template.FuncMap{
		"split": strings.Split,
	}

	// init template object and add custom functions
	template_name := filepath.Base(template_path)
	tmpl := template.New(template_name).Funcs(funcMap)

	// debug
	log.WithFields(log.Fields{
		"template":    template_path,
		"destination": destination_path,
	}).Debug("Input")

	// parse template
	tmpl, err := tmpl.ParseFiles(template_path)
	check(err)

	// prepare destination
	destination := os.Stdout
	if destination_path != "" {
		destination, err = os.Create(destination_path)
		check(err)
		defer destination.Close()
	}

	// execute template
	err = tmpl.ExecuteTemplate(destination, template_name, &Context{})
	check(err)
}
