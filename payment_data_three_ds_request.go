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

// PaymentDataThreeDSRequest represents the PaymentDataThreeDSRequest API object
type PaymentDataThreeDSRequest struct {
	// AcsURL is the uRL of the ACS
	AcsURL *string `json:"acs_url,omitempty"`
	// Pareq is the pAReq used during the 3DS authentication
	Pareq *string `json:"pareq,omitempty"`
	// Md is the mD used during the 3DS authentication
	Md *string `json:"md,omitempty"`
	// TermURL is the uRL of the 3DS term
	TermURL *string `json:"term_url,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the PaymentDataThreeDSRequest object and its
// children
func (s *PaymentDataThreeDSRequest) SetClient(c *ProcessOut) *PaymentDataThreeDSRequest {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *PaymentDataThreeDSRequest) Prefill(c *PaymentDataThreeDSRequest) *PaymentDataThreeDSRequest {
	if c == nil {
		return s
	}

	s.AcsURL = c.AcsURL
	s.Pareq = c.Pareq
	s.Md = c.Md
	s.TermURL = c.TermURL

	return s
}

// dummyPaymentDataThreeDSRequest is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPaymentDataThreeDSRequest() {
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
