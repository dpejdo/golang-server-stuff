package main

import (
	"errors"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/get", getCookies)
	mux.HandleFunc("/set", setCookies)

	http.ListenAndServe(":3000", mux)

}

func getCookies(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.Write([]byte(cookie.Value))

}

func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	w.Write([]byte("cookie set"))
}
