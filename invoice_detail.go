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

// InvoiceDetail represents the InvoiceDetail API object
type InvoiceDetail struct {
	// Name is the name of the invoice detail
	Name *string `json:"name,omitempty"`
	// Type is the type of the invoice detail. Can be a string containing anything, up to 30 characters
	Type *string `json:"type,omitempty"`
	// Amount is the amount represented by the invoice detail
	Amount *string `json:"amount,omitempty"`
	// Quantity is the quantity of items represented by the invoice detail
	Quantity *int `json:"quantity,omitempty"`
	// Metadata is the metadata related to the invoice detail, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoiceDetail object and its
// children
func (s *InvoiceDetail) SetClient(c *ProcessOut) *InvoiceDetail {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceDetail) Prefill(c *InvoiceDetail) *InvoiceDetail {
	if c == nil {
		return s
	}

	s.Name = c.Name
	s.Type = c.Type
	s.Amount = c.Amount
	s.Quantity = c.Quantity
	s.Metadata = c.Metadata

	return s
}

// dummyInvoiceDetail is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceDetail() {
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
