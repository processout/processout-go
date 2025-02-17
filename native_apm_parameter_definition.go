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

// NativeAPMParameterDefinition represents the NativeAPMParameterDefinition API object
type NativeAPMParameterDefinition struct {
	// Key is the native APM parameter value key
	Key *string `json:"key,omitempty"`
	// Type is the nativeAPM parameter value type
	Type *string `json:"type,omitempty"`
	// Required is the nativeAPM parameter value requirement
	Required *bool `json:"required,omitempty"`
	// Length is the nativeAPM parameter value length
	Length *int `json:"length,omitempty"`
	// DisplayName is the native APM parameter display name
	DisplayName *string `json:"display_name,omitempty"`
	// AvailableValues is the native APM parameter available input values
	AvailableValues *[]*NativeAPMParameterValueDefinition `json:"available_values,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMParameterDefinition object and its
// children
func (s *NativeAPMParameterDefinition) SetClient(c *ProcessOut) *NativeAPMParameterDefinition {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMParameterDefinition) Prefill(c *NativeAPMParameterDefinition) *NativeAPMParameterDefinition {
	if c == nil {
		return s
	}

	s.Key = c.Key
	s.Type = c.Type
	s.Required = c.Required
	s.Length = c.Length
	s.DisplayName = c.DisplayName
	s.AvailableValues = c.AvailableValues

	return s
}

// dummyNativeAPMParameterDefinition is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMParameterDefinition() {
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
