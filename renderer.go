package main

import (
	"html/template"
	"slices"
)

func GetTemplateRenderer() *template.Template {
	funcMap := template.FuncMap{
		"containsString": slices.Contains[[]string, string],
	}

	return template.Must(
		template.New("").Funcs(funcMap).ParseGlob("view/**/*.html"),
	)
}
