package render

import (
	"html/template"
	"io"
)

var tpl *template.Template

func Init(template *template.Template) {
	tpl = template
}

func RenderTemplate(w io.Writer, templateName string, viewData interface{}) error {
	return tpl.ExecuteTemplate(w, templateName, viewData)
}
