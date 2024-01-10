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

// InvoiceExternalFraudTools represents the InvoiceExternalFraudTools API object
type InvoiceExternalFraudTools struct {
	// Forter is the forter
	Forter *string `json:"forter,omitempty"`
	// Ravelin is the ravelin
	Ravelin *string `json:"ravelin,omitempty"`
	// Signifyd is the signifyd
	Signifyd *string `json:"signifyd,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoiceExternalFraudTools object and its
// children
func (s *InvoiceExternalFraudTools) SetClient(c *ProcessOut) *InvoiceExternalFraudTools {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceExternalFraudTools) Prefill(c *InvoiceExternalFraudTools) *InvoiceExternalFraudTools {
	if c == nil {
		return s
	}

	s.Forter = c.Forter
	s.Ravelin = c.Ravelin
	s.Signifyd = c.Signifyd

	return s
}

// dummyInvoiceExternalFraudTools is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceExternalFraudTools() {
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
