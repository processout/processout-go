package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Token represents the Token API object
type Token struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the customer token
	ID string `json:"id,omitempty"`
	// Customer is the customer owning the token
	Customer *Customer `json:"customer,omitempty"`
	// CustomerID is the iD of the customer linked to the token, if any
	CustomerID string `json:"customer_id,omitempty"`
	// Card is the card used to create this token, if any
	Card *Card `json:"card,omitempty"`
	// Type is the type of the token. Can be card or gateway_token
	Type string `json:"type,omitempty"`
	// Metadata is the metadata related to the token, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata,omitempty"`
	// IsSubscriptionOnly is the define whether or not the customer token is used on a recurring invoice
	IsSubscriptionOnly bool `json:"is_subscription_only,omitempty"`
	// CreatedAt is the date at which the customer token was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// SetClient sets the client for the Token object and its
// children
func (s *Token) SetClient(c *ProcessOut) {
	if s == nil {
		return
	}
	s.Client = c
	if s.Customer != nil {
		s.Customer.SetClient(c)
	}
	if s.Card != nil {
		s.Card.SetClient(c)
	}
}

// Find allows you to find a customer's token by its ID.
func (s Token) Find(customerID, tokenID string, options ...Options) (*Token, error) {
	if s.Client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Token   *Token `json:"token"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(customerID) + "/tokens/" + url.QueryEscape(tokenID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Token.SetClient(s.Client)
	return payload.Token, nil
}

// Create allows you to create a new token for the given customer ID.
func (s Token) Create(customerID, source string, options ...Options) (*Token, error) {
	if s.Client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Token   *Token `json:"token"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"metadata":    s.Metadata,
		"source":      source,
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(customerID) + "/tokens"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Token.SetClient(s.Client)
	return payload.Token, nil
}

// CreateFromRequest allows you to create a new token for the given customer ID from an authorization request
func (s Token) CreateFromRequest(customerID, source, target string, options ...Options) (*Token, error) {
	if s.Client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Token   *Token `json:"token"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"metadata":    s.Metadata,
		"source":      source,
		"target":      target,
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(customerID) + "/tokens"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Token.SetClient(s.Client)
	return payload.Token, nil
}

// Delete allows you to delete a customer token
func (s Token) Delete(options ...Options) error {
	if s.Client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "customers/" + url.QueryEscape(s.CustomerID) + "/tokens/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
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

// dummyToken is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyToken() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New(nil, "", "")
}
