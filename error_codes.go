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

// ErrorCodes represents the ErrorCodes API object
type ErrorCodes struct {
	// Gateway is the error codes from gateways by category.
	Gateway *CategoryErrorCodes `json:"gateway,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ErrorCodes object and its
// children
func (s *ErrorCodes) SetClient(c *ProcessOut) *ErrorCodes {
	if s == nil {
		return s
	}
	s.client = c
	if s.Gateway != nil {
		s.Gateway.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ErrorCodes) Prefill(c *ErrorCodes) *ErrorCodes {
	if c == nil {
		return s
	}

	s.Gateway = c.Gateway

	return s
}

// ErrorCodesAllParameters is the structure representing the
// additional parameters used to call ErrorCodes.All
type ErrorCodesAllParameters struct {
	*Options
	*ErrorCodes
}

// All allows you to get all error codes.
func (s ErrorCodes) All(options ...ErrorCodesAllParameters) (*ErrorCodes, error) {
	return s.AllWithContext(context.Background(), options...)
}

// All allows you to get all error codes., passes the provided context to the request
func (s ErrorCodes) AllWithContext(ctx context.Context, options ...ErrorCodesAllParameters) (*ErrorCodes, error) {
	if s.client == nil {
		panic("Please use the client.NewErrorCodes() method to create a new ErrorCodes object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ErrorCodesAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ErrorCodes)

	type Response struct {
		ErrorCodes *ErrorCodes `json:""`
		HasMore    bool        `json:"has_more"`
		Success    bool        `json:"success"`
		Message    string      `json:"message"`
		Code       string      `json:"error_type"`
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

	path := "/error-codes"

	req, err := http.NewRequestWithContext(
		ctx,
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

	payload.ErrorCodes.SetClient(s.client)
	return payload.ErrorCodes, nil
}

// dummyErrorCodes is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyErrorCodes() {
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
