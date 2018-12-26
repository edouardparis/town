package templates

import "text/template"

var HTMLTemplates *template.Template

func init() {
	var err error
	HTMLTemplates, err = template.ParseGlob("templates/html/*.html")
	if err != nil {
		panic(err)
	}
}
