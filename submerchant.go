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

// Submerchant represents the Submerchant API object
type Submerchant struct {
	// ID is the iD of the submerchant at ProcessOut
	ID *string `json:"id,omitempty"`
	// Name is the legal name of the submerchant
	Name *string `json:"name,omitempty"`
	// CreatedAt is the time at which the submerchant was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Submerchant) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Submerchant object and its
// children
func (s *Submerchant) SetClient(c *ProcessOut) *Submerchant {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Submerchant) Prefill(c *Submerchant) *Submerchant {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Name = c.Name
	s.CreatedAt = c.CreatedAt

	return s
}

// dummySubmerchant is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummySubmerchant() {
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
