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

// BalancesCustomerAction represents the BalancesCustomerAction API object
type BalancesCustomerAction struct {
	// Type is the customer action type (such as url)
	Type *string `json:"type,omitempty"`
	// Value is the value of the customer action. If type is an URL, URL to which you should redirect your customer
	Value *string `json:"value,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the BalancesCustomerAction object and its
// children
func (s *BalancesCustomerAction) SetClient(c *ProcessOut) *BalancesCustomerAction {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *BalancesCustomerAction) Prefill(c *BalancesCustomerAction) *BalancesCustomerAction {
	if c == nil {
		return s
	}

	s.Type = c.Type
	s.Value = c.Value

	return s
}

// dummyBalancesCustomerAction is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyBalancesCustomerAction() {
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
