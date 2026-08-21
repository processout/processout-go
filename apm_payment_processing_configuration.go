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

// APMPaymentProcessingConfiguration represents the APMPaymentProcessingConfiguration API object
type APMPaymentProcessingConfiguration struct {
	// ReturnRedirectType is the type of redirection performed once the customer returns from the Alternative Payment Method (APM) flow.
	ReturnRedirectType *string `json:"return_redirect_type,omitempty"`
	// PreferredFinalizationMode is the preferred finalization mode requested for the Alternative Payment Method (APM) flow.
	PreferredFinalizationMode *string `json:"preferred_finalization_mode,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the APMPaymentProcessingConfiguration object and its
// children
func (s *APMPaymentProcessingConfiguration) SetClient(c *ProcessOut) *APMPaymentProcessingConfiguration {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *APMPaymentProcessingConfiguration) Prefill(c *APMPaymentProcessingConfiguration) *APMPaymentProcessingConfiguration {
	if c == nil {
		return s
	}

	s.ReturnRedirectType = c.ReturnRedirectType
	s.PreferredFinalizationMode = c.PreferredFinalizationMode

	return s
}

// dummyAPMPaymentProcessingConfiguration is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyAPMPaymentProcessingConfiguration() {
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
