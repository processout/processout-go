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

// InvoiceDevice represents the InvoiceDevice API object
type InvoiceDevice struct {
	// Channel is the channel of the device
	Channel *string `json:"channel,omitempty"`
	// IpAddress is the iP address of the device
	IpAddress *string `json:"ip_address,omitempty"`
	// ID is the iD of the device
	ID *string `json:"id,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *InvoiceDevice) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the InvoiceDevice object and its
// children
func (s *InvoiceDevice) SetClient(c *ProcessOut) *InvoiceDevice {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceDevice) Prefill(c *InvoiceDevice) *InvoiceDevice {
	if c == nil {
		return s
	}

	s.Channel = c.Channel
	s.IpAddress = c.IpAddress
	s.ID = c.ID

	return s
}

// dummyInvoiceDevice is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceDevice() {
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
