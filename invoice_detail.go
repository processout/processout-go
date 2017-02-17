package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// InvoiceDetail represents the InvoiceDetail API object
type InvoiceDetail struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// Type is the type of the invoice detail. Can be a string containing anything, up to 30 characters
	Type string `json:"type,omitempty"`
	// Amount is the amount represented by the invoice detail
	Amount string `json:"amount,omitempty"`
	// Metadata is the metadata related to the invoice detail, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata,omitempty"`
}

// SetClient sets the client for the InvoiceDetail object and its
// children
func (s *InvoiceDetail) SetClient(c *ProcessOut) {
	if s == nil {
		return
	}
	s.Client = c
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
	}
	errors.New(nil, "", "")
}
