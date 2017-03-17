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

// APIVersion represents the APIVersion API object
type APIVersion struct {
	// Name is the name used to identify the API version
	Name string `json:"name,omitempty"`
	// Description is the description of the API version. Can contain a changelog
	Description string `json:"description,omitempty"`
	// CreatedAt is the date at which the API version was released
	CreatedAt time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the APIVersion object and its
// children
func (s *APIVersion) SetClient(c *ProcessOut) *APIVersion {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *APIVersion) Prefill(c *APIVersion) *APIVersion {
	if c == nil {
		return s
	}

	s.Name = c.Name
	s.Description = c.Description
	s.CreatedAt = c.CreatedAt

	return s
}

// dummyAPIVersion is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyAPIVersion() {
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
