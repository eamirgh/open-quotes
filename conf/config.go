package conf

import "html/template"

var Tpl *template.Template

func Init() {
	Tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}