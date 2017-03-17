package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Token represents the Token API object
type Token struct {
	Identifier

	// Customer is the customer owning the token
	Customer *Customer `json:"customer,omitempty"`
	// CustomerID is the iD of the customer linked to the token
	CustomerID string `json:"customer_id,omitempty"`
	// GatewayConfiguration is the gateway configuration to which the token is linked, if any
	GatewayConfiguration *GatewayConfiguration `json:"gateway_configuration,omitempty"`
	// GatewayConfigurationID is the iD of the gateway configuration to which the token is linked, if any
	GatewayConfigurationID *string `json:"gateway_configuration_id,omitempty"`
	// Card is the card used to create this token, if any
	Card *Card `json:"card,omitempty"`
	// CardID is the iD of the card used to create the token, if any
	CardID *string `json:"card_id,omitempty"`
	// Type is the type of the token. Can be card or gateway_token
	Type string `json:"type,omitempty"`
	// Metadata is the metadata related to the token, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata,omitempty"`
	// IsSubscriptionOnly is the define whether or not the customer token is used on a recurring invoice
	IsSubscriptionOnly bool `json:"is_subscription_only,omitempty"`
	// IsDefault is the true if the card it the default card of the customer, false otherwise
	IsDefault bool `json:"is_default,omitempty"`
	// CreatedAt is the date at which the customer token was created
	CreatedAt time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the Token object and its
// children
func (s *Token) SetClient(c *ProcessOut) *Token {
	if s == nil {
		return s
	}
	s.client = c
	if s.Customer != nil {
		s.Customer.SetClient(c)
	}
	if s.GatewayConfiguration != nil {
		s.GatewayConfiguration.SetClient(c)
	}
	if s.Card != nil {
		s.Card.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Token) Prefill(c *Token) *Token {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Customer = c.Customer
	s.CustomerID = c.CustomerID
	s.GatewayConfiguration = c.GatewayConfiguration
	s.GatewayConfigurationID = c.GatewayConfigurationID
	s.Card = c.Card
	s.CardID = c.CardID
	s.Type = c.Type
	s.Metadata = c.Metadata
	s.IsSubscriptionOnly = c.IsSubscriptionOnly
	s.IsDefault = c.IsDefault
	s.CreatedAt = c.CreatedAt

	return s
}

// TokenFindParameters is the structure representing the
// additional parameters used to call Token.Find
type TokenFindParameters struct {
	*Options
	*Token
}

// Find allows you to find a customer's token by its ID.
func (s Token) Find(customerID, tokenID string, options ...TokenFindParameters) (*Token, error) {
	if s.client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TokenFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Token)

	type Response struct {
		Token   *Token `json:"token"`
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
	setupRequest(s.client, opt.Options, req)

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

	payload.Token.SetClient(s.client)
	return payload.Token, nil
}

// TokenCreateParameters is the structure representing the
// additional parameters used to call Token.Create
type TokenCreateParameters struct {
	*Options
	*Token
	Settings interface{} `json:"settings"`
	Target   interface{} `json:"target"`
}

// Create allows you to create a new token for the given customer ID.
func (s Token) Create(customerID, source string, options ...TokenCreateParameters) (*Token, error) {
	if s.client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TokenCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Token)

	type Response struct {
		Token   *Token `json:"token"`
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Metadata interface{} `json:"metadata"`
		Settings interface{} `json:"settings"`
		Target   interface{} `json:"target"`
		Source   interface{} `json:"source"`
	}{
		Options:  opt.Options,
		Metadata: s.Metadata,
		Settings: opt.Settings,
		Target:   opt.Target,
		Source:   source,
	}

	body, err := json.Marshal(data)
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
	setupRequest(s.client, opt.Options, req)

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

	payload.Token.SetClient(s.client)
	return payload.Token, nil
}

// TokenCreateFromRequestParameters is the structure representing the
// additional parameters used to call Token.CreateFromRequest
type TokenCreateFromRequestParameters struct {
	*Options
	*Token
	Settings interface{} `json:"settings"`
}

// CreateFromRequest allows you to create a new token for the given customer ID from an authorization request
func (s Token) CreateFromRequest(customerID, source, target string, options ...TokenCreateFromRequestParameters) (*Token, error) {
	if s.client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TokenCreateFromRequestParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Token)

	type Response struct {
		Token   *Token `json:"token"`
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Metadata interface{} `json:"metadata"`
		Settings interface{} `json:"settings"`
		Source   interface{} `json:"source"`
		Target   interface{} `json:"target"`
	}{
		Options:  opt.Options,
		Metadata: s.Metadata,
		Settings: opt.Settings,
		Source:   source,
		Target:   target,
	}

	body, err := json.Marshal(data)
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
	setupRequest(s.client, opt.Options, req)

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

	payload.Token.SetClient(s.client)
	return payload.Token, nil
}

// TokenDeleteParameters is the structure representing the
// additional parameters used to call Token.Delete
type TokenDeleteParameters struct {
	*Options
	*Token
}

// Delete allows you to delete a customer token
func (s Token) Delete(options ...TokenDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TokenDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Token)

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

	path := "/customers/" + url.QueryEscape(s.CustomerID) + "/tokens/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		g io.Reader
	}
	errors.New(nil, "", "")
}
