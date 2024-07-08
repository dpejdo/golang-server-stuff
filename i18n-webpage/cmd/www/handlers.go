package main

import (
	"fmt"
	"net/http"

	"bookstore/internal/localizer"
	_ "bookstore/internal/translations"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	l, ok := localizer.Get(r.URL.Query().Get(":locale"))
	if !ok {
		http.NotFound(w, r)
		return
	}

	var totalBookCount = 1_252

	fmt.Fprintln(w, l.Translate("Welcome!"))

	fmt.Fprintln(w, l.Translate("%d books available", totalBookCount))

	fmt.Fprintln(w, l.Translate("Launching soon"))
}
