package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type (
	user struct {
		Id    string `json:"id"`
		Login string `json:"login"`
	}
)

func (auth *AuthService) Ping() bool {
	response, err := http.Get("http://" + auth.Addr + "/ping")
	if err != nil {
		return false
	}

	if response.StatusCode != 200 {
		return false
	} else {
		return true
	}
}

func (auth *AuthService) Register(login string, password string) (string, error) {
	body, _ := json.Marshal(map[string]string{"login": login, "password": password})
	response, err := http.Post("http://"+auth.Addr+"/register", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	text, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 201 {
		return "", errors.New(string(text))
	}

	return string(text), nil
}

func (auth *AuthService) Login(login string, password string) (string, error) {
	body, _ := json.Marshal(map[string]string{"login": login, "password": password})
	response, err := http.Post("http://"+auth.Addr+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	text, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New(string(text))
	}

	return string(text), nil
}

func (auth *AuthService) LoginToken(token string) (*user, error) {
	response, err := http.Post("http://"+auth.Addr+"/token", "text", bytes.NewBuffer([]byte(token)))
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		text, _ := io.ReadAll(response.Body)
		return nil, errors.New(string(text))
	}

	var user user
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
