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

// InvoiceShippingPhone represents the InvoiceShippingPhone API object
type InvoiceShippingPhone struct {
	// Number is the phone number for the shipment
	Number *string `json:"number,omitempty"`
	// DialingCode is the phone number dialing code for the shipment
	DialingCode *string `json:"dialing_code,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoiceShippingPhone object and its
// children
func (s *InvoiceShippingPhone) SetClient(c *ProcessOut) *InvoiceShippingPhone {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceShippingPhone) Prefill(c *InvoiceShippingPhone) *InvoiceShippingPhone {
	if c == nil {
		return s
	}

	s.Number = c.Number
	s.DialingCode = c.DialingCode

	return s
}

// dummyInvoiceShippingPhone is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceShippingPhone() {
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
