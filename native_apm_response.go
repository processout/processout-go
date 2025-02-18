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

// NativeAPMResponse represents the NativeAPMResponse API object
type NativeAPMResponse struct {
	// State is the native APM response state
	State *string `json:"state,omitempty"`
	// ParameterDefinitions is the native APM parameter values description
	ParameterDefinitions *[]*NativeAPMParameterDefinition `json:"parameter_definitions,omitempty"`
	// ParameterValues is the native APM parameter values
	ParameterValues *[]*NativeAPMParameterValue `json:"parameter_values,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMResponse object and its
// children
func (s *NativeAPMResponse) SetClient(c *ProcessOut) *NativeAPMResponse {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMResponse) Prefill(c *NativeAPMResponse) *NativeAPMResponse {
	if c == nil {
		return s
	}

	s.State = c.State
	s.ParameterDefinitions = c.ParameterDefinitions
	s.ParameterValues = c.ParameterValues

	return s
}

// dummyNativeAPMResponse is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMResponse() {
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
