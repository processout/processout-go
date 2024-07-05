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

// CardShipping represents the CardShipping API object
type CardShipping struct {
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
	// Phone is the shipping phone number
	Phone *Phone `json:"phone,omitempty"`
	// FirstName is the first name of the card shipping
	FirstName *string `json:"first_name,omitempty"`
	// LastName is the last name of the card shipping
	LastName *string `json:"last_name,omitempty"`
	// Email is the email of the card shipping
	Email *string `json:"email,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CardShipping object and its
// children
func (s *CardShipping) SetClient(c *ProcessOut) *CardShipping {
	if s == nil {
		return s
	}
	s.client = c
	if s.Phone != nil {
		s.Phone.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CardShipping) Prefill(c *CardShipping) *CardShipping {
	if c == nil {
		return s
	}

	s.Address1 = c.Address1
	s.Address2 = c.Address2
	s.City = c.City
	s.State = c.State
	s.CountryCode = c.CountryCode
	s.Zip = c.Zip
	s.Phone = c.Phone
	s.FirstName = c.FirstName
	s.LastName = c.LastName
	s.Email = c.Email

	return s
}

// dummyCardShipping is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCardShipping() {
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
