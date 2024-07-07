package service

import (
	"github.com/slack-go/slack"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserInput struct {
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	ValidationErrors map[string]string `json:"-"`
}

func (s *Service) RegisterUser(input *RegisterUserInput) error {
	input.ValidationErrors = make(map[string]string)

	if input.Username == "" {
		input.ValidationErrors["username"] = "must be provided"
	}

	if len(input.ValidationErrors) > 0 {
		return ErrFailedValidation
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)

	if err != nil {
		return err
	}

	_, err = s.DB.Exec("INSERT INTO (username, hashed_password) VALUES ($1,$2)", input.Username, hashedPassword)
	if err != nil {
		return err
	}

	msg := slack.WebhookMessage{
		Username: "robot",
		Channel:  "#general",
		Text:     "A new user has signed up!",
	}

	return slack.PostWebhook(s.SlackWebhookURL, &msg)
}
