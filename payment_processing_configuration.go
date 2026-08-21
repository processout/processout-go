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

// PaymentProcessingConfiguration represents the PaymentProcessingConfiguration API object
type PaymentProcessingConfiguration struct {
	// BypassUnsupportedSplitPayments is the payment processing should bypass unsupported split payments validation when no payment gateway supports it.
	BypassUnsupportedSplitPayments *bool `json:"bypass_unsupported_split_payments,omitempty"`
	// ApmPaymentConfig is the alternative Payment Method (APM) specific payment processing cofiguration.
	ApmPaymentConfig *APMPaymentProcessingConfiguration `json:"apm_payment_config,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the PaymentProcessingConfiguration object and its
// children
func (s *PaymentProcessingConfiguration) SetClient(c *ProcessOut) *PaymentProcessingConfiguration {
	if s == nil {
		return s
	}
	s.client = c
	if s.ApmPaymentConfig != nil {
		s.ApmPaymentConfig.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *PaymentProcessingConfiguration) Prefill(c *PaymentProcessingConfiguration) *PaymentProcessingConfiguration {
	if c == nil {
		return s
	}

	s.BypassUnsupportedSplitPayments = c.BypassUnsupportedSplitPayments
	s.ApmPaymentConfig = c.ApmPaymentConfig

	return s
}

// dummyPaymentProcessingConfiguration is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPaymentProcessingConfiguration() {
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
