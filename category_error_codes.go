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

// CategoryErrorCodes represents the CategoryErrorCodes API object
type CategoryErrorCodes struct {
	// Generic is the generic error codes.
	Generic *[]string `json:"generic,omitempty"`
	// Service is the service related error codes.
	Service *[]string `json:"service,omitempty"`
	// Gateway is the gateway related error codes.
	Gateway *[]string `json:"gateway,omitempty"`
	// Card is the card related error codes.
	Card *[]string `json:"card,omitempty"`
	// Check is the check related error codes.
	Check *[]string `json:"check,omitempty"`
	// Shipping is the shipping related error codes.
	Shipping *[]string `json:"shipping,omitempty"`
	// Customer is the customer related error codes.
	Customer *[]string `json:"customer,omitempty"`
	// Payment is the payment related error codes.
	Payment *[]string `json:"payment,omitempty"`
	// Refund is the refund related error codes.
	Refund *[]string `json:"refund,omitempty"`
	// Wallet is the wallet related error codes.
	Wallet *[]string `json:"wallet,omitempty"`
	// Request is the request related error codes.
	Request *[]string `json:"request,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CategoryErrorCodes object and its
// children
func (s *CategoryErrorCodes) SetClient(c *ProcessOut) *CategoryErrorCodes {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CategoryErrorCodes) Prefill(c *CategoryErrorCodes) *CategoryErrorCodes {
	if c == nil {
		return s
	}

	s.Generic = c.Generic
	s.Service = c.Service
	s.Gateway = c.Gateway
	s.Card = c.Card
	s.Check = c.Check
	s.Shipping = c.Shipping
	s.Customer = c.Customer
	s.Payment = c.Payment
	s.Refund = c.Refund
	s.Wallet = c.Wallet
	s.Request = c.Request

	return s
}

// dummyCategoryErrorCodes is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCategoryErrorCodes() {
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
