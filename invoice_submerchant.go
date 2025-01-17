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

// InvoiceSubmerchant represents the InvoiceSubmerchant API object
type InvoiceSubmerchant struct {
	// ID is the iD of the invoice submerchant
	ID *string `json:"id,omitempty"`
	// Name is the name of the submerchant
	Name *string `json:"name,omitempty"`
	// Reference is the submerchant's reference ID
	Reference *string `json:"reference,omitempty"`
	// Mcc is the submerchant's MCC (Merchant Category Code).
	Mcc *string `json:"mcc,omitempty"`
	// PhoneNumber is the submerchant's phone number
	PhoneNumber *SubmerchantPhoneNumber `json:"phone_number,omitempty"`
	// Email is the email address
	Email *string `json:"email,omitempty"`
	// Address is the submerchant's address
	Address *SubmerchantAddress `json:"address,omitempty"`
	// TaxReference is the tax reference
	TaxReference *string `json:"tax_reference,omitempty"`
	// ServiceEstablishmentNumber is the service establishment number
	ServiceEstablishmentNumber *string `json:"service_establishment_number,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *InvoiceSubmerchant) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the InvoiceSubmerchant object and its
// children
func (s *InvoiceSubmerchant) SetClient(c *ProcessOut) *InvoiceSubmerchant {
	if s == nil {
		return s
	}
	s.client = c
	if s.PhoneNumber != nil {
		s.PhoneNumber.SetClient(c)
	}
	if s.Address != nil {
		s.Address.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceSubmerchant) Prefill(c *InvoiceSubmerchant) *InvoiceSubmerchant {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Name = c.Name
	s.Reference = c.Reference
	s.Mcc = c.Mcc
	s.PhoneNumber = c.PhoneNumber
	s.Email = c.Email
	s.Address = c.Address
	s.TaxReference = c.TaxReference
	s.ServiceEstablishmentNumber = c.ServiceEstablishmentNumber

	return s
}

// dummyInvoiceSubmerchant is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceSubmerchant() {
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
