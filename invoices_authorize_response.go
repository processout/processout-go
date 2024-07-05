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

// InvoicesAuthorizeResponse represents the InvoicesAuthorizeResponse API object
type InvoicesAuthorizeResponse struct {
	// Transaction is the transaction linked to the invoice
	Transaction *Transaction `json:"transaction,omitempty"`
	// CustomerAction is the customer action to be performed
	CustomerAction *CustomerAction `json:"customer_action,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoicesAuthorizeResponse object and its
// children
func (s *InvoicesAuthorizeResponse) SetClient(c *ProcessOut) *InvoicesAuthorizeResponse {
	if s == nil {
		return s
	}
	s.client = c
	if s.Transaction != nil {
		s.Transaction.SetClient(c)
	}
	if s.CustomerAction != nil {
		s.CustomerAction.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoicesAuthorizeResponse) Prefill(c *InvoicesAuthorizeResponse) *InvoicesAuthorizeResponse {
	if c == nil {
		return s
	}

	s.Transaction = c.Transaction
	s.CustomerAction = c.CustomerAction

	return s
}

// dummyInvoicesAuthorizeResponse is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoicesAuthorizeResponse() {
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
