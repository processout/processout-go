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

// ApplePayAlternativeMerchantCertificates represents the ApplePayAlternativeMerchantCertificates API object
type ApplePayAlternativeMerchantCertificates struct {
	// AlternativeMerchantCertificates is the alternative merchant certificates available
	AlternativeMerchantCertificates *[]*AlternativeMerchantCertificate `json:"alternative_merchant_certificates,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ApplePayAlternativeMerchantCertificates object and its
// children
func (s *ApplePayAlternativeMerchantCertificates) SetClient(c *ProcessOut) *ApplePayAlternativeMerchantCertificates {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ApplePayAlternativeMerchantCertificates) Prefill(c *ApplePayAlternativeMerchantCertificates) *ApplePayAlternativeMerchantCertificates {
	if c == nil {
		return s
	}

	s.AlternativeMerchantCertificates = c.AlternativeMerchantCertificates

	return s
}

// ApplePayAlternativeMerchantCertificatesFetchParameters is the structure representing the
// additional parameters used to call ApplePayAlternativeMerchantCertificates.Fetch
type ApplePayAlternativeMerchantCertificatesFetchParameters struct {
	*Options
	*ApplePayAlternativeMerchantCertificates
}

// Fetch allows you to fetch the project's alternative certificates by ID
func (s ApplePayAlternativeMerchantCertificates) Fetch(options ...ApplePayAlternativeMerchantCertificatesFetchParameters) (*ApplePayAlternativeMerchantCertificates, error) {
	if s.client == nil {
		panic("Please use the client.NewApplePayAlternativeMerchantCertificates() method to create a new ApplePayAlternativeMerchantCertificates object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ApplePayAlternativeMerchantCertificatesFetchParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ApplePayAlternativeMerchantCertificates)

	type Response struct {
		ApplePayAlternativeMerchantCertificates *ApplePayAlternativeMerchantCertificates `json:"applepay_certificates"`
		HasMore                                 bool                                     `json:"has_more"`
		Success                                 bool                                     `json:"success"`
		Message                                 string                                   `json:"message"`
		Code                                    string                                   `json:"error_type"`
	}

	data := struct {
		*Options
	}{
		Options: opt.Options,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/projects/applepay/alternative-merchant-certificates"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.ApplePayAlternativeMerchantCertificates.SetClient(s.client)
	return payload.ApplePayAlternativeMerchantCertificates, nil
}

// dummyApplePayAlternativeMerchantCertificates is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyApplePayAlternativeMerchantCertificates() {
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
