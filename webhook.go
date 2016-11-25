package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Webhook represents the Webhook API object
type Webhook struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the recurring invoice
	ID string `json:"id,omitempty"`
	// Project is the project to which the webhook belongs
	Project *Project `json:"project,omitempty"`
	// Event is the event the webhook is linked to
	Event *Event `json:"event,omitempty"`
	// RequestURL is the uRL to which the webhook will be posted
	RequestURL string `json:"request_url,omitempty"`
	// RequestMethod is the method used to send the webhook (GET or POST)
	RequestMethod string `json:"request_method,omitempty"`
	// ResponseBody is the the response body the webhook received when sending its payload
	ResponseBody string `json:"response_body,omitempty"`
	// ResponseCode is the the response code the webhook received when sending its payload
	ResponseCode string `json:"response_code,omitempty"`
	// ResponseHeaders is the the response headers the webhook received when sending its payload
	ResponseHeaders string `json:"response_headers,omitempty"`
	// ResponseTimeMs is the the time it took for the webhook to send its payload
	ResponseTimeMs int `json:"response_time_ms,omitempty"`
	// Status is the the status of the webhook. 0: pending, 1: success, 2: error
	Status int `json:"status,omitempty"`
	// CreatedAt is the date at which the webhook was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ReleaseAt is the date at webhook will be/is released
	ReleaseAt *time.Time `json:"release_at,omitempty"`
}

// SetClient sets the client for the Webhook object and its
// children
func (s *Webhook) SetClient(c *ProcessOut) {
	if s == nil {
		return
	}
	s.Client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Event != nil {
		s.Event.SetClient(c)
	}
}

// dummyWebhook is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyWebhook() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New(nil, "", "")
}
