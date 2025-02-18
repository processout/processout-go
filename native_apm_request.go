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

// NativeAPMRequest represents the NativeAPMRequest API object
type NativeAPMRequest struct {
	// ParameterValues is the native APM parameter values
	ParameterValues *[]*NativeAPMParameterValue `json:"parameter_values,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMRequest object and its
// children
func (s *NativeAPMRequest) SetClient(c *ProcessOut) *NativeAPMRequest {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMRequest) Prefill(c *NativeAPMRequest) *NativeAPMRequest {
	if c == nil {
		return s
	}

	s.ParameterValues = c.ParameterValues

	return s
}

// dummyNativeAPMRequest is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMRequest() {
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
