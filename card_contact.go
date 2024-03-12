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

// CardContact represents the CardContact API object
type CardContact struct {
	// Address1 is the address line of the card holder
	Address1 *string `json:"address1,omitempty"`
	// Address2 is the secondary address line of the card holder
	Address2 *string `json:"address2,omitempty"`
	// City is the city of the card holder
	City *string `json:"city,omitempty"`
	// State is the state of the card holder
	State *string `json:"state,omitempty"`
	// CountryCode is the country code of the card holder (ISO-3166, 2 characters format)
	CountryCode *string `json:"country_code,omitempty"`
	// Zip is the zIP code of the card holder
	Zip *string `json:"zip,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CardContact object and its
// children
func (s *CardContact) SetClient(c *ProcessOut) *CardContact {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CardContact) Prefill(c *CardContact) *CardContact {
	if c == nil {
		return s
	}

	s.Address1 = c.Address1
	s.Address2 = c.Address2
	s.City = c.City
	s.State = c.State
	s.CountryCode = c.CountryCode
	s.Zip = c.Zip

	return s
}

// dummyCardContact is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCardContact() {
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
