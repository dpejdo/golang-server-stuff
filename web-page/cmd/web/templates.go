package main

import (
	"web-page.hg6p.com/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
