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

// InvoiceTax represents the InvoiceTax API object
type InvoiceTax struct {
	// Amount is the amount of the tax for an invoice
	Amount *string `json:"amount,omitempty"`
	// Rate is the rate of the tax for an invoice
	Rate *string `json:"rate,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoiceTax object and its
// children
func (s *InvoiceTax) SetClient(c *ProcessOut) *InvoiceTax {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceTax) Prefill(c *InvoiceTax) *InvoiceTax {
	if c == nil {
		return s
	}

	s.Amount = c.Amount
	s.Rate = c.Rate

	return s
}

// dummyInvoiceTax is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceTax() {
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
