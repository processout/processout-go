package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// Events manages the Event struct
type Events struct {
	p *ProcessOut
}

type Event struct {
	// Data : Data associated to the event, in the form of a dictionary
	Data map[string]string `json:"data"`
	// Date : The date at which the event was fired
	Date time.Time `json:"date"`
	// ID : Id of the event
	ID string `json:"id"`
	// Name : Name of the event
	Name string `json:"name"`
	// Sandbox : Whether or not the event was fired in the sandbox environment
	Sandbox bool `json:"sandbox"`
}

// Pull : Get the 15 oldest events pending processing.
func (e Events) Pull() ([]*Event, error) {
	type Response struct {
		Events  []*Event `json:"events"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
	}

	path := "/events"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(e.p.projectID, e.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return payload.Events, nil
}

// SetAllProcessed : Set all the events as processed.
func (e Events) SetAllProcessed() error {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/events"

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", e.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(e.p.projectID, e.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return err
	}

	if !payload.Success {
		return errors.New(payload.Message)
	}
	return nil
}

// Find : Get the information related to the specific event.
func (e Events) Find(ID string) (*Event, error) {
	type Response struct {
		Event   `json:"event"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/events/{id}"
	path = strings.Replace(path, "{id}", ID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(e.p.projectID, e.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Event, nil
}

// MarkProcessed : Set the specific event as processed.
func (e Events) MarkProcessed(event *Event) error {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/events/{id}"
	path = strings.Replace(path, "{id}", event.ID, -1)

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", e.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(e.p.projectID, e.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return err
	}

	if !payload.Success {
		return errors.New(payload.Message)
	}
	return nil
}

// dummyEvent is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyEvent() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
	}
	errors.New("")
}
