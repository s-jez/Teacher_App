package config

import (
	"html/template"
)

var TPL *template.Template

func LoadTemplate() {
	// Load templates from path
	TPL = template.Must(template.ParseGlob("assets/*"))
}
