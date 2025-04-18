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

// InvoiceRisk represents the InvoiceRisk API object
type InvoiceRisk struct {
	// Score is the scoring of the invoice
	Score *string `json:"score,omitempty"`
	// IsLegit is the define whether or not the invoice is legit
	IsLegit *bool `json:"is_legit,omitempty"`
	// SkipGatewayRules is the skip payment gateway fraud engine rules (on compatible gateways only.)
	SkipGatewayRules *bool `json:"skip_gateway_rules,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the InvoiceRisk object and its
// children
func (s *InvoiceRisk) SetClient(c *ProcessOut) *InvoiceRisk {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *InvoiceRisk) Prefill(c *InvoiceRisk) *InvoiceRisk {
	if c == nil {
		return s
	}

	s.Score = c.Score
	s.IsLegit = c.IsLegit
	s.SkipGatewayRules = c.SkipGatewayRules

	return s
}

// dummyInvoiceRisk is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoiceRisk() {
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
