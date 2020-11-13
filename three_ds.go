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

// ThreeDS represents the ThreeDS API object
type ThreeDS struct {
	// Version is the version of the 3DS
	Version *string `json:"Version,omitempty"`
	// Status is the current status of the authentication
	Status *string `json:"Status,omitempty"`
	// Fingerprinted is the true if a fingerprint has occured
	Fingerprinted *bool `json:"fingerprinted,omitempty"`
	// Challenged is the true if a challenge has occured
	Challenged *bool `json:"challenged,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ThreeDS object and its
// children
func (s *ThreeDS) SetClient(c *ProcessOut) *ThreeDS {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ThreeDS) Prefill(c *ThreeDS) *ThreeDS {
	if c == nil {
		return s
	}

	s.Version = c.Version
	s.Status = c.Status
	s.Fingerprinted = c.Fingerprinted
	s.Challenged = c.Challenged

	return s
}

// dummyThreeDS is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyThreeDS() {
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
