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

// Token represents the Token API object
type Token struct {
	// ID is the iD of the customer token
	ID *string `json:"id,omitempty"`
	// Customer is the customer owning the token
	Customer *Customer `json:"customer,omitempty"`
	// CustomerID is the iD of the customer linked to the token
	CustomerID *string `json:"customer_id,omitempty"`
	// GatewayConfiguration is the gateway configuration to which the token is linked, if any
	GatewayConfiguration *GatewayConfiguration `json:"gateway_configuration,omitempty"`
	// GatewayConfigurationID is the iD of the gateway configuration to which the token is linked, if any
	GatewayConfigurationID *string `json:"gateway_configuration_id,omitempty"`
	// Card is the card used to create this token, if any
	Card *Card `json:"card,omitempty"`
	// CardID is the iD of the card used to create the token, if any
	CardID *string `json:"card_id,omitempty"`
	// Type is the type of the token. Can be card or gateway_token
	Type *string `json:"type,omitempty"`
	// Metadata is the metadata related to the token, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// IsSubscriptionOnly is the define whether or not the customer token is used on a recurring invoice
	IsSubscriptionOnly *bool `json:"is_subscription_only,omitempty"`
	// ReturnURL is the uRL where the customer will be redirected upon payment authentication (if required by tokenization method)
	ReturnURL *string `json:"return_url,omitempty"`
	// CancelURL is the uRL where the customer will be redirected if the tokenization was canceled (if required by tokenization method)
	CancelURL *string `json:"cancel_url,omitempty"`
	// IsDefault is the true if the token it the default token of the customer, false otherwise
	IsDefault *bool `json:"is_default,omitempty"`
	// IsChargeable is the true if the token is chargeable, false otherwise
	IsChargeable *bool `json:"is_chargeable,omitempty"`
	// CreatedAt is the date at which the customer token was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Token) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
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
	s.ReturnURL = c.ReturnURL
	s.CancelURL = c.CancelURL
	s.IsDefault = c.IsDefault
	s.IsChargeable = c.IsChargeable
	s.CreatedAt = c.CreatedAt

	return s
}

// TokenVerifyParameters is the structure representing the
// additional parameters used to call Token.Verify
type TokenVerifyParameters struct {
	*Options
	*Token
}

// Verify allows you to verify a customer token's card is valid.
func (s Token) Verify(options ...TokenVerifyParameters) error {
	if s.client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TokenVerifyParameters{}
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

	path := "/customers/" + url.QueryEscape(*s.CustomerID) + "/tokens/" + url.QueryEscape(*s.ID) + "/verify"

	req, err := http.NewRequest(
		"POST",
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

// TokenFetchCustomerTokensParameters is the structure representing the
// additional parameters used to call Token.FetchCustomerTokens
type TokenFetchCustomerTokensParameters struct {
	*Options
	*Token
}

// FetchCustomerTokens allows you to get the customer's tokens.
func (s Token) FetchCustomerTokens(customerID string, options ...TokenFetchCustomerTokensParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TokenFetchCustomerTokensParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Token)

	type Response struct {
		Tokens []*Token `json:"tokens"`

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

	path := "/customers/" + url.QueryEscape(customerID) + "/tokens"

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

	tokensList := []Identifiable{}
	for _, o := range payload.Tokens {
		tokensList = append(tokensList, o.SetClient(s.client))
	}
	tokensIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    tokensList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"tokens"`
				HasMore bool            `json:"has_more"`
			}{}
			if err := json.NewDecoder(b).Decode(&r); err != nil {
				return false, err
			}
			if err := json.Unmarshal(r.Data, i); err != nil {
				return false, err
			}
			return r.HasMore, nil
		},
		client:      s.client,
		hasMoreNext: payload.HasMore,
		hasMorePrev: false,
	}
	return tokensIterator, nil
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

	payload.Token.SetClient(s.client)
	return payload.Token, nil
}

// TokenCreateParameters is the structure representing the
// additional parameters used to call Token.Create
type TokenCreateParameters struct {
	*Options
	*Token
	Source         interface{} `json:"source"`
	Settings       interface{} `json:"settings"`
	Target         interface{} `json:"target"`
	Verify         interface{} `json:"verify"`
	VerifyMetadata interface{} `json:"verify_metadata"`
	SetDefault     interface{} `json:"set_default"`
}

// Create allows you to create a new token for the given customer ID.
func (s Token) Create(options ...TokenCreateParameters) (*Token, error) {
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
		Metadata       interface{} `json:"metadata"`
		ReturnURL      interface{} `json:"return_url"`
		CancelURL      interface{} `json:"cancel_url"`
		Source         interface{} `json:"source"`
		Settings       interface{} `json:"settings"`
		Target         interface{} `json:"target"`
		Verify         interface{} `json:"verify"`
		VerifyMetadata interface{} `json:"verify_metadata"`
		SetDefault     interface{} `json:"set_default"`
	}{
		Options:        opt.Options,
		Metadata:       s.Metadata,
		ReturnURL:      s.ReturnURL,
		CancelURL:      s.CancelURL,
		Source:         opt.Source,
		Settings:       opt.Settings,
		Target:         opt.Target,
		Verify:         opt.Verify,
		VerifyMetadata: opt.VerifyMetadata,
		SetDefault:     opt.SetDefault,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(*s.CustomerID) + "/tokens"

	req, err := http.NewRequest(
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

	payload.Token.SetClient(s.client)
	return payload.Token, nil
}

// TokenSaveParameters is the structure representing the
// additional parameters used to call Token.Save
type TokenSaveParameters struct {
	*Options
	*Token
}

// Save allows you to save the updated customer attributes.
func (s Token) Save(options ...TokenSaveParameters) error {
	if s.client == nil {
		panic("Please use the client.NewToken() method to create a new Token object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TokenSaveParameters{}
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

	path := "/customers/" + url.QueryEscape(*s.CustomerID) + "/tokens/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	path := "/customers/" + url.QueryEscape(*s.CustomerID) + "/tokens/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
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
