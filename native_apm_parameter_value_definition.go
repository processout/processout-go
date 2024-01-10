package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v4/errors"
)

// NativeAPMParameterValueDefinition represents the NativeAPMParameterValueDefinition API object
type NativeAPMParameterValueDefinition struct {
	// Value is the native APM parameter value
	Value *string `json:"value,omitempty"`
	// Default is the native APM parameter default value flag
	Default *bool `json:"default,omitempty"`
	// DisplayName is the native APM parameter value display name
	DisplayName *string `json:"display_name,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMParameterValueDefinition object and its
// children
func (s *NativeAPMParameterValueDefinition) SetClient(c *ProcessOut) *NativeAPMParameterValueDefinition {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMParameterValueDefinition) Prefill(c *NativeAPMParameterValueDefinition) *NativeAPMParameterValueDefinition {
	if c == nil {
		return s
	}

	s.Value = c.Value
	s.Default = c.Default
	s.DisplayName = c.DisplayName

	return s
}

// dummyNativeAPMParameterValueDefinition is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMParameterValueDefinition() {
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
