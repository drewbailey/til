package til

import (
	"text/template"
)

func TemplateProcessor() *template.Template {
	return template.New("til")
}
