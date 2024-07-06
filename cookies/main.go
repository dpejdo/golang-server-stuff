package main

import (
	"cookie/internal/cookies"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
)

var secretKey []byte

func main() {
	var err error

	secretKey, err = hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/get", getCookies)
	mux.HandleFunc("/set", setCookies)

	http.ListenAndServe(":3000", mux)

}

func getCookies(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.SignedRead(r, "exampleCookie", secretKey)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, cookies.ErrInvalidValue):
			http.Error(w, "invalid cookie", http.StatusBadRequest)
		default:
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.Write([]byte(cookie))

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

	err := cookies.SignedWrite(w, cookie, secretKey)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte("cookie set"))
}
