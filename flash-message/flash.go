package main

import (
	"encoding/base64"
	"net/http"
	"time"
)

func setFlash(w http.ResponseWriter, name string, value string) {

	http.SetCookie(w, &http.Cookie{Name: name, Value: encode([]byte(value))})
}

func getFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, err
		default:
			return nil, err
		}
	}

	value, err := decode(cookie.Value)
	if err != nil {
		return nil, err
	}
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return value, nil

}

func encode(value []byte) string {
	return base64.URLEncoding.EncodeToString(value)
}

func decode(value string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(value)
}
