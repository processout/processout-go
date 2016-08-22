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

// Activitys manages the Activity struct
type Activities struct {
	p *ProcessOut
}

type Activity struct {
	// ID : ID of the activity
	ID string `json:"id"`
	// Title : Title of the activity
	Title string `json:"title"`
	// Content : Content of the activity
	Content string `json:"content"`
	// Level : Level of the activity
	Level int `json:"level"`
	// CreatedAt : Date at which the transaction was created
	CreatedAt time.Time `json:"created_at"`
}

// All : Get all the project activities.
func (s Activities) All() ([]*Activity, error) {

	type Response struct {
		Activities []*Activity `json:"activities"`
		Success    bool        `json:"success"`
		Message    string      `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/activities"

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
	return payload.Activities, nil
}

// Find : Find a specific activity and fetch its data.
func (s Activities) Find(activityID string) (*Activity, error) {

	type Response struct {
		Activity `json:"activity"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/activities/" + url.QueryEscape(activityID) + ""

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
	return &payload.Activity, nil
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
	}
	errors.New("")
}
