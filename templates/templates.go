package templates

import "embed"

//go:embed *.tmpl
var TemplatesFs embed.FS

func GetTemplateFs() *embed.FS {
	return &TemplatesFs
}
