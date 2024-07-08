package main

import (
	"net/http"

	_ "bookstore/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	locale := r.URL.Query().Get(":locale")

	var lang language.Tag

	switch locale {
	case "en-gb":
		lang = language.MustParse("en-GB")

	case "de-de":
		lang = language.MustParse("de-DE")
	case "fr-ch":
		lang = language.MustParse("fr-CH")

	default:
		http.NotFound(w, r)
		return
	}

	var totalBookCount = 1_252_794

	p := message.NewPrinter(lang)
	p.Fprintf(w, "Welcome!\n")

	p.Fprintf(w, "%d books available\n", totalBookCount)
}
