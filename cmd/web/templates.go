package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"snippetbox.shrishail.dev/internal/models"
	"snippetbox.shrishail.dev/ui"
)

type templateData struct {
	Snippet         models.Snippet
	Snippets        []models.Snippet
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func templateCache() (map[string]*template.Template, error) {

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	cache := map[string]*template.Template{}

	for _, page := range pages {
		fileName := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err := template.New(fileName).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[fileName] = ts
	}

	return cache, nil
}
