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

// UnsupportedFeatureBypass represents the UnsupportedFeatureBypass API object
type UnsupportedFeatureBypass struct {
	// IncrementalAuthorization is the indicates whether to fallback to normal authorization if incremental is not supported
	IncrementalAuthorization *bool `json:"incremental_authorization,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the UnsupportedFeatureBypass object and its
// children
func (s *UnsupportedFeatureBypass) SetClient(c *ProcessOut) *UnsupportedFeatureBypass {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *UnsupportedFeatureBypass) Prefill(c *UnsupportedFeatureBypass) *UnsupportedFeatureBypass {
	if c == nil {
		return s
	}

	s.IncrementalAuthorization = c.IncrementalAuthorization

	return s
}

// dummyUnsupportedFeatureBypass is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyUnsupportedFeatureBypass() {
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
