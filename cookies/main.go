package main

import (
	"bytes"
	"cookie/internal/cookies"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var secretKey []byte

type User struct {
	Name string
	Age  int
}

func main() {
	var err error

	gob.Register(&User{})

	secretKey, err = hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/get", getCookies)
	mux.HandleFunc("/set", setCookies)

	http.ListenAndServe(":3000", mux)

}

func setCookies(w http.ResponseWriter, _ *http.Request) {

	user := User{"yasuo", 23}

	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    buf.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	err = cookies.EncryptedWrite(w, cookie, secretKey)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Cookie set"))

}

func getCookies(w http.ResponseWriter, r *http.Request) {
	gobEncodedValue, err := cookies.EncryptedRead(r, "exampleCookie", secretKey)
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

	var user User
	reader := strings.NewReader(gobEncodedValue)

	if err := gob.NewDecoder(reader).Decode(&user); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Name: %q\n", user.Name)
	fmt.Fprintf(w, "Age: %d\n", user.Age)
	response := fmt.Sprintf("Name: %q\nAge: %d\n", user.Name, user.Age)
	w.Write([]byte(response))

}
