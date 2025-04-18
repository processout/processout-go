package processout

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
)

// AlternativeMerchantCertificate represents the AlternativeMerchantCertificate API object
type AlternativeMerchantCertificate struct {
	// ID is the id of the alternative merchant certificate
	ID *string `json:"id,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *AlternativeMerchantCertificate) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the AlternativeMerchantCertificate object and its
// children
func (s *AlternativeMerchantCertificate) SetClient(c *ProcessOut) *AlternativeMerchantCertificate {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *AlternativeMerchantCertificate) Prefill(c *AlternativeMerchantCertificate) *AlternativeMerchantCertificate {
	if c == nil {
		return s
	}

	s.ID = c.ID

	return s
}

// AlternativeMerchantCertificateSaveParameters is the structure representing the
// additional parameters used to call AlternativeMerchantCertificate.Save
type AlternativeMerchantCertificateSaveParameters struct {
	*Options
	*AlternativeMerchantCertificate
}

// Save allows you to save new alternative apple pay certificates
func (s AlternativeMerchantCertificate) Save(options ...AlternativeMerchantCertificateSaveParameters) (*AlternativeMerchantCertificate, error) {
	return s.SaveWithContext(context.Background(), options...)
}

// Save allows you to save new alternative apple pay certificates, passes the provided context to the request
func (s AlternativeMerchantCertificate) SaveWithContext(ctx context.Context, options ...AlternativeMerchantCertificateSaveParameters) (*AlternativeMerchantCertificate, error) {
	if s.client == nil {
		panic("Please use the client.NewAlternativeMerchantCertificate() method to create a new AlternativeMerchantCertificate object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AlternativeMerchantCertificateSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.AlternativeMerchantCertificate)

	type Response struct {
		AlternativeMerchantCertificate *AlternativeMerchantCertificate `json:"alternative_merchant_certificate"`
		HasMore                        bool                            `json:"has_more"`
		Success                        bool                            `json:"success"`
		Message                        string                          `json:"message"`
		Code                           string                          `json:"error_type"`
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

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
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

	payload.AlternativeMerchantCertificate.SetClient(s.client)
	return payload.AlternativeMerchantCertificate, nil
}

// AlternativeMerchantCertificateDeleteParameters is the structure representing the
// additional parameters used to call AlternativeMerchantCertificate.Delete
type AlternativeMerchantCertificateDeleteParameters struct {
	*Options
	*AlternativeMerchantCertificate
}

// Delete allows you to delete a given alternative merchant certificate
func (s AlternativeMerchantCertificate) Delete(options ...AlternativeMerchantCertificateDeleteParameters) error {
	return s.DeleteWithContext(context.Background(), options...)
}

// Delete allows you to delete a given alternative merchant certificate, passes the provided context to the request
func (s AlternativeMerchantCertificate) DeleteWithContext(ctx context.Context, options ...AlternativeMerchantCertificateDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewAlternativeMerchantCertificate() method to create a new AlternativeMerchantCertificate object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AlternativeMerchantCertificateDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.AlternativeMerchantCertificate)

	type Response struct {
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
	}{
		Options: opt.Options,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "/projects/applepay/alternative-merchant-certificates/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return erri
	}

	return nil
}

// dummyAlternativeMerchantCertificate is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyAlternativeMerchantCertificate() {
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
