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

// NativeAPMTransactionDetailsInvoice represents the NativeAPMTransactionDetailsInvoice API object
type NativeAPMTransactionDetailsInvoice struct {
	// Amount is the native APM Invoice amount
	Amount *string `json:"amount,omitempty"`
	// CurrencyCode is the native APM Invoice currency code
	CurrencyCode *string `json:"currency_code,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMTransactionDetailsInvoice object and its
// children
func (s *NativeAPMTransactionDetailsInvoice) SetClient(c *ProcessOut) *NativeAPMTransactionDetailsInvoice {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMTransactionDetailsInvoice) Prefill(c *NativeAPMTransactionDetailsInvoice) *NativeAPMTransactionDetailsInvoice {
	if c == nil {
		return s
	}

	s.Amount = c.Amount
	s.CurrencyCode = c.CurrencyCode

	return s
}

// dummyNativeAPMTransactionDetailsInvoice is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMTransactionDetailsInvoice() {
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
