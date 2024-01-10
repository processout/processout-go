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

// InvoiceShipping represents the InvoiceShipping API object
type InvoiceShipping struct {
	// Amount is the amount of the shipping
	Amount *string `json:"amount,omitempty"`
	// Method is the delivery method
	Method *string `json:"method,omitempty"`
	// Provider is the delivery provider
	Provider *string `json:"provider,omitempty"`
	// Delay is the shipping delay
	Delay *string `json:"delay,omitempty"`
	// Address1 is the address where the shipment will be delivered
	Address1 *string `json:"address1,omitempty"`
	// Address2 is the secondary address where the shipment will be delivered
	Address2 *string `json:"address2,omitempty"`
	// City is the city where the shipment will be delivered
	City *string `json:"city,omitempty"`
	// State is the state where the shipment will be delivered
	State *string `json:"state,omitempty"`
	// CountryCode is the country code where the shipment will be delivered
	CountryCode *string `json:"country_code,omitempty"`
	// Zip is the zIP where the shipment will be delivered
	Zip *string `json:"zip,omitempty"`
	// PhoneNumber is the shipment full phone number, consisting of a combined dialing code and phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Phone is the phone number for the shipment
	Phone *InvoiceShippingPhone `json:"phone,omitempty"`
	// ExpectsShippingAt is the date at which the shipment is expected to be sent
	ExpectsShippingAt *time.Time `json:"expects_shipping_at,omitempty"`
	// RelayStoreName is the relay store name
	RelayStoreName *string `json:"relay_store_name,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoiceShipping object and its
// children
func (s *InvoiceShipping) SetClient(c *ProcessOut) *InvoiceShipping {
	if s == nil {
		return s
	}
	s.client = c
	if s.Phone != nil {
		s.Phone.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceShipping) Prefill(c *InvoiceShipping) *InvoiceShipping {
	if c == nil {
		return s
	}

	s.Amount = c.Amount
	s.Method = c.Method
	s.Provider = c.Provider
	s.Delay = c.Delay
	s.Address1 = c.Address1
	s.Address2 = c.Address2
	s.City = c.City
	s.State = c.State
	s.CountryCode = c.CountryCode
	s.Zip = c.Zip
	s.PhoneNumber = c.PhoneNumber
	s.Phone = c.Phone
	s.ExpectsShippingAt = c.ExpectsShippingAt
	s.RelayStoreName = c.RelayStoreName

	return s
}

// dummyInvoiceShipping is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceShipping() {
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
