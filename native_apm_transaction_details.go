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

// NativeAPMTransactionDetails represents the NativeAPMTransactionDetails API object
type NativeAPMTransactionDetails struct {
	// Gateway is the native APM Gateway details
	Gateway *NativeAPMTransactionDetailsGateway `json:"gateway,omitempty"`
	// Invoice is the native APM Invoice details
	Invoice *NativeAPMTransactionDetailsInvoice `json:"invoice,omitempty"`
	// Parameters is the native APM Parameter details
	Parameters *[]*NativeAPMParameterDefinition `json:"parameters,omitempty"`
	// State is the native APM Transaction State
	State *string `json:"state,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMTransactionDetails object and its
// children
func (s *NativeAPMTransactionDetails) SetClient(c *ProcessOut) *NativeAPMTransactionDetails {
	if s == nil {
		return s
	}
	s.client = c
	if s.Gateway != nil {
		s.Gateway.SetClient(c)
	}
	if s.Invoice != nil {
		s.Invoice.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMTransactionDetails) Prefill(c *NativeAPMTransactionDetails) *NativeAPMTransactionDetails {
	if c == nil {
		return s
	}

	s.Gateway = c.Gateway
	s.Invoice = c.Invoice
	s.Parameters = c.Parameters
	s.State = c.State

	return s
}

// dummyNativeAPMTransactionDetails is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMTransactionDetails() {
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
