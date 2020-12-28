package conf

import "html/template"

var Locales = []string{"en_US"}
var Tpl *template.Template

func Init() {
	Tpl = template.Must(template.ParseGlob("resources/template/*.gohtml"))
}
