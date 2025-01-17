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

// SubmerchantPhoneNumber represents the SubmerchantPhoneNumber API object
type SubmerchantPhoneNumber struct {
	// DialingCode is the dialing code of the phone number
	DialingCode *string `json:"dialing_code,omitempty"`
	// Number is the phone number
	Number *string `json:"number,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the SubmerchantPhoneNumber object and its
// children
func (s *SubmerchantPhoneNumber) SetClient(c *ProcessOut) *SubmerchantPhoneNumber {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *SubmerchantPhoneNumber) Prefill(c *SubmerchantPhoneNumber) *SubmerchantPhoneNumber {
	if c == nil {
		return s
	}

	s.DialingCode = c.DialingCode
	s.Number = c.Number

	return s
}

// dummySubmerchantPhoneNumber is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummySubmerchantPhoneNumber() {
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
