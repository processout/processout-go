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

// SubmerchantAddress represents the SubmerchantAddress API object
type SubmerchantAddress struct {
	// Line1 is the address line 1
	Line1 *string `json:"line1,omitempty"`
	// Line2 is the address line 2
	Line2 *string `json:"line2,omitempty"`
	// City is the city
	City *string `json:"city,omitempty"`
	// State is the state
	State *string `json:"state,omitempty"`
	// CountryCode is the country code
	CountryCode *string `json:"country_code,omitempty"`
	// Zip is the zIP code
	Zip *string `json:"zip,omitempty"`
	// County is the county (US specific)
	County *string `json:"county,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the SubmerchantAddress object and its
// children
func (s *SubmerchantAddress) SetClient(c *ProcessOut) *SubmerchantAddress {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *SubmerchantAddress) Prefill(c *SubmerchantAddress) *SubmerchantAddress {
	if c == nil {
		return s
	}

	s.Line1 = c.Line1
	s.Line2 = c.Line2
	s.City = c.City
	s.State = c.State
	s.CountryCode = c.CountryCode
	s.Zip = c.Zip
	s.County = c.County

	return s
}

// dummySubmerchantAddress is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummySubmerchantAddress() {
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
