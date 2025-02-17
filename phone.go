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

// Phone represents the Phone API object
type Phone struct {
	// Number is the phone number (without dialing code)
	Number *string `json:"number,omitempty"`
	// DialingCode is the phone number dialing code
	DialingCode *string `json:"dialing_code,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the Phone object and its
// children
func (s *Phone) SetClient(c *ProcessOut) *Phone {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Phone) Prefill(c *Phone) *Phone {
	if c == nil {
		return s
	}

	s.Number = c.Number
	s.DialingCode = c.DialingCode

	return s
}

// dummyPhone is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPhone() {
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
