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

// Balance represents the Balance API object
type Balance struct {
	// Amount is the amount available
	Amount *string `json:"amount,omitempty"`
	// Currency is the currency the balance is in
	Currency *string `json:"currency,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the Balance object and its
// children
func (s *Balance) SetClient(c *ProcessOut) *Balance {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Balance) Prefill(c *Balance) *Balance {
	if c == nil {
		return s
	}

	s.Amount = c.Amount
	s.Currency = c.Currency

	return s
}

// dummyBalance is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyBalance() {
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
