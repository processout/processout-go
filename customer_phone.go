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

// CustomerPhone represents the CustomerPhone API object
type CustomerPhone struct {
	// Number is the phone number of the customer
	Number *string `json:"number,omitempty"`
	// DialingCode is the phone number dialing code of the customer
	DialingCode *string `json:"dialing_code,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CustomerPhone object and its
// children
func (s *CustomerPhone) SetClient(c *ProcessOut) *CustomerPhone {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CustomerPhone) Prefill(c *CustomerPhone) *CustomerPhone {
	if c == nil {
		return s
	}

	s.Number = c.Number
	s.DialingCode = c.DialingCode

	return s
}

// dummyCustomerPhone is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCustomerPhone() {
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
