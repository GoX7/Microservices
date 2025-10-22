package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type (
	EventsResponse struct {
		Events []*Event `json:"events"`
	}
	Event struct {
		ID          string `json:"id"`
		Login       string `json:"login"`
		Event       string `json:"event"`
		Createstamp int64  `json:"createstamp"`
	}
)

func (events *EventService) Ping() bool {
	response, err := http.Get("http://" + events.Addr + "/ping")
	if err != nil {
		return false
	}

	if response.StatusCode != 200 {
		return false
	} else {
		return true
	}
}

func (events *EventService) Register(login string, event string) (string, error) {
	body, _ := json.Marshal(map[string]string{"login": login, "event": event})
	resp, err := http.Post("http://"+events.Addr+"/register", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 201 {
		return "", errors.New(string(data))
	}

	return string(data), nil
}

func (events *EventService) GetEvents(login string) (*EventsResponse, error) {
	response, err := http.Get("http://" + events.Addr + "/events/" + login)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	var event EventsResponse
	if err := json.NewDecoder(response.Body).Decode(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (events *EventService) GetEvent(id string, login string) (*Event, error) {
	response, err := http.Get("http://" + events.Addr + "/events/" + login + "/" + id)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	var event Event
	if err := json.NewDecoder(response.Body).Decode(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (events *EventService) DeleteEvent(id string, login string) error {
	request, err := http.NewRequest("DELETE", "http://"+events.Addr+"/events/"+login+"/"+id, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}
