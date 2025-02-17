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

// CardUpdateRequest represents the CardUpdateRequest API object
type CardUpdateRequest struct {
	// PreferredScheme is the customer preferred scheme, such as carte bancaire vs visa. Can be set to none to clear the previous value
	PreferredScheme *string `json:"preferred_scheme,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CardUpdateRequest object and its
// children
func (s *CardUpdateRequest) SetClient(c *ProcessOut) *CardUpdateRequest {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CardUpdateRequest) Prefill(c *CardUpdateRequest) *CardUpdateRequest {
	if c == nil {
		return s
	}

	s.PreferredScheme = c.PreferredScheme

	return s
}

// CardUpdateRequestUpdateParameters is the structure representing the
// additional parameters used to call CardUpdateRequest.Update
type CardUpdateRequestUpdateParameters struct {
	*Options
	*CardUpdateRequest
}

// Update allows you to update a card by its ID.
func (s CardUpdateRequest) Update(cardID string, options ...CardUpdateRequestUpdateParameters) (*CardUpdateRequest, error) {
	if s.client == nil {
		panic("Please use the client.NewCardUpdateRequest() method to create a new CardUpdateRequest object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CardUpdateRequestUpdateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.CardUpdateRequest)

	type Response struct {
		CardUpdateRequest *CardUpdateRequest `json:"card"`
		HasMore           bool               `json:"has_more"`
		Success           bool               `json:"success"`
		Message           string             `json:"message"`
		Code              string             `json:"error_type"`
	}

	data := struct {
		*Options
		PreferredScheme interface{} `json:"preferred_scheme"`
	}{
		Options:         opt.Options,
		PreferredScheme: s.PreferredScheme,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/cards/" + url.QueryEscape(cardID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	payload.CardUpdateRequest.SetClient(s.client)
	return payload.CardUpdateRequest, nil
}

// dummyCardUpdateRequest is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCardUpdateRequest() {
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
