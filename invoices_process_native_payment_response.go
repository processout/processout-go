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

// InvoicesProcessNativePaymentResponse represents the InvoicesProcessNativePaymentResponse API object
type InvoicesProcessNativePaymentResponse struct {
	// Transaction is the transaction linked to this Native APM
	Transaction *Transaction `json:"transaction,omitempty"`
	// NativeApm is the native APM response
	NativeApm *NativeAPMResponse `json:"native_apm,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoicesProcessNativePaymentResponse object and its
// children
func (s *InvoicesProcessNativePaymentResponse) SetClient(c *ProcessOut) *InvoicesProcessNativePaymentResponse {
	if s == nil {
		return s
	}
	s.client = c
	if s.Transaction != nil {
		s.Transaction.SetClient(c)
	}
	if s.NativeApm != nil {
		s.NativeApm.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoicesProcessNativePaymentResponse) Prefill(c *InvoicesProcessNativePaymentResponse) *InvoicesProcessNativePaymentResponse {
	if c == nil {
		return s
	}

	s.Transaction = c.Transaction
	s.NativeApm = c.NativeApm

	return s
}

// dummyInvoicesProcessNativePaymentResponse is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoicesProcessNativePaymentResponse() {
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
