package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
)

// Event represents the Event API object
type Event struct {
	// ID is the iD of the event
	ID *string `json:"id,omitempty"`
	// Project is the project to which the event belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the event belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Name is the name of the event
	Name *string `json:"name,omitempty"`
	// Data is the data object associated to the event
	Data interface{} `json:"data,omitempty"`
	// Sandbox is the define whether or not the event is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// FiredAt is the date at which the event was fired
	FiredAt *time.Time `json:"fired_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Event) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Event object and its
// children
func (s *Event) SetClient(c *ProcessOut) *Event {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Event) Prefill(c *Event) *Event {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Name = c.Name
	s.Data = c.Data
	s.Sandbox = c.Sandbox
	s.FiredAt = c.FiredAt

	return s
}

// EventFetchWebhooksParameters is the structure representing the
// additional parameters used to call Event.FetchWebhooks
type EventFetchWebhooksParameters struct {
	*Options
	*Event
}

// FetchWebhooks allows you to get all the webhooks of the event.
func (s Event) FetchWebhooks(options ...EventFetchWebhooksParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewEvent() method to create a new Event object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := EventFetchWebhooksParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Event)

	type Response struct {
		Webhooks []*Webhook `json:"webhooks"`

		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
	}{
		Options: opt.Options,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/events/" + url.QueryEscape(*s.ID) + "/webhooks"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	webhooksList := []Identifiable{}
	for _, o := range payload.Webhooks {
		webhooksList = append(webhooksList, o.SetClient(s.client))
	}
	webhooksIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    webhooksList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"webhooks"`
				HasMore bool            `json:"has_more"`
			}{}
			if err := json.NewDecoder(b).Decode(&r); err != nil {
				return false, err
			}
			if err := json.Unmarshal(r.Data, i); err != nil {
				return false, err
			}
			return r.HasMore, nil
		},
		client:      s.client,
		hasMoreNext: payload.HasMore,
		hasMorePrev: false,
	}
	return webhooksIterator, nil
}

// EventAllParameters is the structure representing the
// additional parameters used to call Event.All
type EventAllParameters struct {
	*Options
	*Event
}

// All allows you to get all the events.
func (s Event) All(options ...EventAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewEvent() method to create a new Event object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := EventAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Event)

	type Response struct {
		Events []*Event `json:"events"`

		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
	}{
		Options: opt.Options,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/events"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	eventsList := []Identifiable{}
	for _, o := range payload.Events {
		eventsList = append(eventsList, o.SetClient(s.client))
	}
	eventsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    eventsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"events"`
				HasMore bool            `json:"has_more"`
			}{}
			if err := json.NewDecoder(b).Decode(&r); err != nil {
				return false, err
			}
			if err := json.Unmarshal(r.Data, i); err != nil {
				return false, err
			}
			return r.HasMore, nil
		},
		client:      s.client,
		hasMoreNext: payload.HasMore,
		hasMorePrev: false,
	}
	return eventsIterator, nil
}

// EventFindParameters is the structure representing the
// additional parameters used to call Event.Find
type EventFindParameters struct {
	*Options
	*Event
}

// Find allows you to find an event by its ID.
func (s Event) Find(eventID string, options ...EventFindParameters) (*Event, error) {
	if s.client == nil {
		panic("Please use the client.NewEvent() method to create a new Event object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := EventFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Event)

	type Response struct {
		Event   *Event `json:"event"`
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
	}{
		Options: opt.Options,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/events/" + url.QueryEscape(eventID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Event.SetClient(s.client)
	return payload.Event, nil
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
