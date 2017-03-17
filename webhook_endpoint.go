package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// WebhookEndpoint represents the WebhookEndpoint API object
type WebhookEndpoint struct {
	Identifier

	// Project is the project to which the webhook endpoint belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the webhook belongs
	ProjectID string `json:"project_id,omitempty"`
	// URL is the uRL to which the webhook endpoint points to
	URL string `json:"url,omitempty"`
	// EventsWhitelist is the slice of string representing the whitelisted events posted to the endpoint
	EventsWhitelist interface{} `json:"events_whitelist,omitempty"`
	// Sandbox is the define whether or not the webhook endpoint is in sandbox environment
	Sandbox bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the webhook endpoint was created
	CreatedAt time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the WebhookEndpoint object and its
// children
func (s *WebhookEndpoint) SetClient(c *ProcessOut) *WebhookEndpoint {
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
func (s *WebhookEndpoint) Prefill(c *WebhookEndpoint) *WebhookEndpoint {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.URL = c.URL
	s.EventsWhitelist = c.EventsWhitelist
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// dummyWebhookEndpoint is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyWebhookEndpoint() {
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
