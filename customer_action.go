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

// CustomerAction represents the CustomerAction API object
type CustomerAction struct {
	// Type is the customer action type (such as url)
	Type *string `json:"type,omitempty"`
	// Value is the value of the customer action. If type is an URL, URL to which you should redirect your customer
	Value *string `json:"value,omitempty"`
	// Metadata is the metadata related to the customer action, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CustomerAction object and its
// children
func (s *CustomerAction) SetClient(c *ProcessOut) *CustomerAction {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CustomerAction) Prefill(c *CustomerAction) *CustomerAction {
	if c == nil {
		return s
	}

	s.Type = c.Type
	s.Value = c.Value
	s.Metadata = c.Metadata

	return s
}

// dummyCustomerAction is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCustomerAction() {
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
