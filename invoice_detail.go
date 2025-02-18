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

// InvoiceDetail represents the InvoiceDetail API object
type InvoiceDetail struct {
	// ID is the iD of the invoice detail
	ID *string `json:"id,omitempty"`
	// Name is the name of the invoice detail
	Name *string `json:"name,omitempty"`
	// Type is the type of the invoice detail. Can be a string containing anything, up to 30 characters
	Type *string `json:"type,omitempty"`
	// Amount is the amount represented by the invoice detail
	Amount *string `json:"amount,omitempty"`
	// Quantity is the quantity of items represented by the invoice detail
	Quantity *int `json:"quantity,omitempty"`
	// Metadata is the metadata related to the invoice detail, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Reference is the reference of the product
	Reference *string `json:"reference,omitempty"`
	// Description is the description of the invoice detail
	Description *string `json:"description,omitempty"`
	// Brand is the brand of the product
	Brand *string `json:"brand,omitempty"`
	// Model is the model of the product
	Model *string `json:"model,omitempty"`
	// DiscountAmount is the discount amount represented by the invoice detail
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// Condition is the condition of the product
	Condition *string `json:"condition,omitempty"`
	// MarketplaceMerchant is the marketplace merchant of the invoice detail
	MarketplaceMerchant *string `json:"marketplace_merchant,omitempty"`
	// MarketplaceMerchantIsBusiness is the define whether or not the marketplace merchant is a business
	MarketplaceMerchantIsBusiness *bool `json:"marketplace_merchant_is_business,omitempty"`
	// MarketplaceMerchantCreatedAt is the date at which the merchant was created
	MarketplaceMerchantCreatedAt *time.Time `json:"marketplace_merchant_created_at,omitempty"`
	// Category is the category of the product
	Category *string `json:"category,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *InvoiceDetail) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the InvoiceDetail object and its
// children
func (s *InvoiceDetail) SetClient(c *ProcessOut) *InvoiceDetail {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceDetail) Prefill(c *InvoiceDetail) *InvoiceDetail {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Name = c.Name
	s.Type = c.Type
	s.Amount = c.Amount
	s.Quantity = c.Quantity
	s.Metadata = c.Metadata
	s.Reference = c.Reference
	s.Description = c.Description
	s.Brand = c.Brand
	s.Model = c.Model
	s.DiscountAmount = c.DiscountAmount
	s.Condition = c.Condition
	s.MarketplaceMerchant = c.MarketplaceMerchant
	s.MarketplaceMerchantIsBusiness = c.MarketplaceMerchantIsBusiness
	s.MarketplaceMerchantCreatedAt = c.MarketplaceMerchantCreatedAt
	s.Category = c.Category

	return s
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
