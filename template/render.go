package template

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	log "github.com/Sirupsen/logrus"

	e "github.com/sorinlg/stencil/err"
)

func Render(template_path string, destination_path string) {
	// load custom template functions
	funcMap := template.FuncMap{
		"GenListFromString":  GenListFromString,
		"RemoveItemFromList": RemoveItemFromList,
		"Default":            Default,
		"Panic":              Panic,
	}

	// init template object and add custom functions
	template_name := filepath.Base(template_path)
	tmpl := template.New(template_name).Funcs(funcMap)

	// debug
	log.WithFields(log.Fields{
		"template":    template_path,
		"destination": destination_path,
	}).Debug("Input")

	// read template file
	file, err := ioutil.ReadFile(template_path)
	e.Check(err)

	// remove escaped endlines
	file_contents := string(file)
	re := regexp.MustCompile("\\\\[[:space:]]*")
	file_contents = re.ReplaceAllString(file_contents, "")

	// parse template
	tmpl, err = tmpl.Parse(file_contents)
	e.Check(err)

	// prepare destination
	destination := os.Stdout
	if destination_path != "" {
		destination, err = os.Create(destination_path)
		e.Check(err)
		defer destination.Close()
	}

	// execute template
	err = tmpl.ExecuteTemplate(destination, template_name, *NewContext())
	// err = tmpl.ExecuteTemplate(destination, template_name, &Context{})
	e.Check(err)
}
