package localizer

import (
	_ "bookstore/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Localizer struct {
	ID      string
	printer *message.Printer
}

var locales = []Localizer{
	{ID: "de-de",
		printer: message.NewPrinter(language.MustParse("de-DE"))},
	{
		ID:      "en-gb",
		printer: message.NewPrinter(language.MustParse("en-GB")),
	},
	{ID: "fr-ch",
		printer: message.NewPrinter(language.MustParse("fr-CH")),
	},
}

func Get(id string) (Localizer, bool) {
	for _, locale := range locales {
		if id == locale.ID {
			return locale, true
		}
	}

	return Localizer{}, false
}

func (l *Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
