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

// InvoiceBilling represents the InvoiceBilling API object
type InvoiceBilling struct {
	// Address1 is the address of the cardholder
	Address1 *string `json:"address1,omitempty"`
	// Address2 is the secondary address of the cardholder
	Address2 *string `json:"address2,omitempty"`
	// City is the city of the cardholder
	City *string `json:"city,omitempty"`
	// State is the state of the cardholder
	State *string `json:"state,omitempty"`
	// CountryCode is the country code of the cardholder
	CountryCode *string `json:"country_code,omitempty"`
	// Zip is the zIP of the cardholder
	Zip *string `json:"zip,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoiceBilling object and its
// children
func (s *InvoiceBilling) SetClient(c *ProcessOut) *InvoiceBilling {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceBilling) Prefill(c *InvoiceBilling) *InvoiceBilling {
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

// dummyInvoiceBilling is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceBilling() {
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
