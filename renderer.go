package main

import (
	"html/template"
	"os"
	"slices"
)

func GetTemplateRenderer() *template.Template {
	funcMap := template.FuncMap{
		"containsString": slices.Contains[[]string, string],
		"isProd":         isProd,
		"plusPlus":       plusPlus,
	}

	return template.Must(
		template.New("").Funcs(funcMap).ParseGlob("view/**/*.html"),
	)
}

func isProd() bool {
	return os.Getenv("GIN_MODE") == "release"
}

func plusPlus(n int) int {
	return n + 1
}
