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

// Activity represents the Activity API object
type Activity struct {
	// ID is the iD of the activity
	ID *string `json:"id,omitempty"`
	// Project is the project to which the activity belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the activity belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Title is the title of the activity
	Title *string `json:"title,omitempty"`
	// Content is the content of the activity
	Content *string `json:"content,omitempty"`
	// Level is the level of the activity
	Level *int `json:"level,omitempty"`
	// CreatedAt is the date at which the transaction was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Activity) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Activity object and its
// children
func (s *Activity) SetClient(c *ProcessOut) *Activity {
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
func (s *Activity) Prefill(c *Activity) *Activity {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Title = c.Title
	s.Content = c.Content
	s.Level = c.Level
	s.CreatedAt = c.CreatedAt

	return s
}

// ActivityAllParameters is the structure representing the
// additional parameters used to call Activity.All
type ActivityAllParameters struct {
	*Options
	*Activity
}

// All allows you to get all the project activities.
func (s Activity) All(options ...ActivityAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewActivity() method to create a new Activity object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ActivityAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Activity)

	type Response struct {
		Activities []*Activity `json:"activities"`

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

	path := "/activities"

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

	activitiesList := []Identifiable{}
	for _, o := range payload.Activities {
		activitiesList = append(activitiesList, o.SetClient(s.client))
	}
	activitiesIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    activitiesList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"activities"`
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
	return activitiesIterator, nil
}

// ActivityFindParameters is the structure representing the
// additional parameters used to call Activity.Find
type ActivityFindParameters struct {
	*Options
	*Activity
}

// Find allows you to find a specific activity and fetch its data.
func (s Activity) Find(activityID string, options ...ActivityFindParameters) (*Activity, error) {
	if s.client == nil {
		panic("Please use the client.NewActivity() method to create a new Activity object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ActivityFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Activity)

	type Response struct {
		Activity *Activity `json:"activity"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
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

	path := "/activities/" + url.QueryEscape(activityID) + ""

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

	payload.Activity.SetClient(s.client)
	return payload.Activity, nil
}

// dummyActivity is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyActivity() {
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
