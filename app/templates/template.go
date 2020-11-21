package templates

import (
	"html/template"
	"net/http"
)

var TPL *template.Template = template.Must(template.ParseGlob("app/templates/*.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	TPL.ExecuteTemplate(w, tmpl+".html", data)
}
