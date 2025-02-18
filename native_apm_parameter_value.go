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

// NativeAPMParameterValue represents the NativeAPMParameterValue API object
type NativeAPMParameterValue struct {
	// Key is the native APM parameter value key
	Key *string `json:"key,omitempty"`
	// Value is the native APM parameter value value
	Value *string `json:"value,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMParameterValue object and its
// children
func (s *NativeAPMParameterValue) SetClient(c *ProcessOut) *NativeAPMParameterValue {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMParameterValue) Prefill(c *NativeAPMParameterValue) *NativeAPMParameterValue {
	if c == nil {
		return s
	}

	s.Key = c.Key
	s.Value = c.Value

	return s
}

// dummyNativeAPMParameterValue is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMParameterValue() {
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
