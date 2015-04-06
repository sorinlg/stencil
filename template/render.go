package template

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"

	"github.com/sorinlg/stencil/error"
)

func Render(template_path string, destination_path string) {
	// load custom template functions
	funcMap := template.FuncMap{
		"split":   strings.Split,
		"list":    List,
		"exclude": ListExclude,
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
	error.Check(err)

	// prepare destination
	destination := os.Stdout
	if destination_path != "" {
		destination, err = os.Create(destination_path)
		error.Check(err)
		defer destination.Close()
	}

	// execute template
	err = tmpl.ExecuteTemplate(destination, template_name, &Context{})
	error.Check(err)
}
