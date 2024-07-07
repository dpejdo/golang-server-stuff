package main

import (
	"net/http"

	"fet-service/internal/service"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input service.RegisterUserInput

	err := app.decodeJson(r.Body, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	err = app.service.RegisterUser(&input)

	if err != nil {
		if errors.is(err, service.ErrFailedValidation) {
			app.failedValidation(w, r, input.ValidationErrors)

		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
