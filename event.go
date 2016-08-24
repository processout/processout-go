package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Events manages the Event struct
type Events struct {
	p *ProcessOut
}

type Event struct {
	// ID : ID of the event
	ID string `json:"id"`
	// Name : Name of the event
	Name string `json:"name"`
	// Data : Data object associated to the event
	Data interface{} `json:"data"`
	// Sandbox : Define whether or not the event is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// FiredAt : Date at which the event was fired
	FiredAt time.Time `json:"fired_at"`
}

// Webhooks : Get all the webhooks of the event.
func (s Events) Webhooks(event *Event) ([]*Webhook, error) {

	type Response struct {
		Webhooks []*Webhook `json:"webhooks"`
		Success  bool       `json:"success"`
		Message  string     `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/events/" + url.QueryEscape(event.ID) + "/webhooks"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

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
	return payload.Webhooks, nil
}

// All : Get all the events.
func (s Events) All() ([]*Event, error) {

	type Response struct {
		Events  []*Event `json:"events"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
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
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

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

// Find : Find an event by its ID.
func (s Events) Find(eventID string) (*Event, error) {

	type Response struct {
		Event   `json:"event"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/events/" + url.QueryEscape(eventID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

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
		f url.URL
	}
	errors.New("")
}
