package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"web-page.hg6p.com/internal/models"
	"web-page.hg6p.com/ui"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
	Flash       string
	IsAuth      bool
	CSRFToken   string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(funcMap).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}
	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var funcMap = template.FuncMap{
	"humanDate": humanDate,
}
