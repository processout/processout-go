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

// AuthorizationRequest represents the AuthorizationRequest API object
type AuthorizationRequest struct {
	// ID is the iD of the authorization
	ID *string `json:"id,omitempty"`
	// Project is the project to which the authorization request belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the authorization request belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Customer is the customer linked to the authorization request
	Customer *Customer `json:"customer,omitempty"`
	// CustomerID is the iD of the customer linked to the authorization request
	CustomerID *string `json:"customer_id,omitempty"`
	// Token is the token linked to the authorization request, once authorized
	Token *Token `json:"token,omitempty"`
	// TokenID is the iD of the token linked to the authorization request, once authorized
	TokenID *string `json:"token_id,omitempty"`
	// Name is the name of the authorization
	Name *string `json:"name,omitempty"`
	// Currency is the currency of the authorization
	Currency *string `json:"currency,omitempty"`
	// ReturnURL is the uRL where the customer will be redirected upon authorization
	ReturnURL *string `json:"return_url,omitempty"`
	// CancelURL is the uRL where the customer will be redirected if the authorization was canceled
	CancelURL *string `json:"cancel_url,omitempty"`
	// Authorized is the whether or not the authorization request was authorized
	Authorized *bool `json:"authorized,omitempty"`
	// Sandbox is the define whether or not the authorization is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// URL is the uRL to which you may redirect your customer to proceed with the authorization
	URL *string `json:"url,omitempty"`
	// CreatedAt is the date at which the authorization was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *AuthorizationRequest) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the AuthorizationRequest object and its
// children
func (s *AuthorizationRequest) SetClient(c *ProcessOut) *AuthorizationRequest {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Customer != nil {
		s.Customer.SetClient(c)
	}
	if s.Token != nil {
		s.Token.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *AuthorizationRequest) Prefill(c *AuthorizationRequest) *AuthorizationRequest {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Customer = c.Customer
	s.CustomerID = c.CustomerID
	s.Token = c.Token
	s.TokenID = c.TokenID
	s.Name = c.Name
	s.Currency = c.Currency
	s.ReturnURL = c.ReturnURL
	s.CancelURL = c.CancelURL
	s.Authorized = c.Authorized
	s.Sandbox = c.Sandbox
	s.URL = c.URL
	s.CreatedAt = c.CreatedAt

	return s
}

// AuthorizationRequestFetchCustomerParameters is the structure representing the
// additional parameters used to call AuthorizationRequest.FetchCustomer
type AuthorizationRequestFetchCustomerParameters struct {
	*Options
	*AuthorizationRequest
}

// FetchCustomer allows you to get the customer linked to the authorization request.
func (s AuthorizationRequest) FetchCustomer(options ...AuthorizationRequestFetchCustomerParameters) (*Customer, error) {
	if s.client == nil {
		panic("Please use the client.NewAuthorizationRequest() method to create a new AuthorizationRequest object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AuthorizationRequestFetchCustomerParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.AuthorizationRequest)

	type Response struct {
		Customer *Customer `json:"customer"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
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

	path := "/authorization-requests/" + url.QueryEscape(*s.ID) + "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
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

	payload.Customer.SetClient(s.client)
	return payload.Customer, nil
}

// AuthorizationRequestCreateParameters is the structure representing the
// additional parameters used to call AuthorizationRequest.Create
type AuthorizationRequestCreateParameters struct {
	*Options
	*AuthorizationRequest
}

// Create allows you to create a new authorization request for the given customer ID.
func (s AuthorizationRequest) Create(options ...AuthorizationRequestCreateParameters) (*AuthorizationRequest, error) {
	if s.client == nil {
		panic("Please use the client.NewAuthorizationRequest() method to create a new AuthorizationRequest object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AuthorizationRequestCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.AuthorizationRequest)

	type Response struct {
		AuthorizationRequest *AuthorizationRequest `json:"authorization_request"`
		HasMore              bool                  `json:"has_more"`
		Success              bool                  `json:"success"`
		Message              string                `json:"message"`
		Code                 string                `json:"error_type"`
	}

	data := struct {
		*Options
		CustomerID interface{} `json:"customer_id"`
		Name       interface{} `json:"name"`
		Currency   interface{} `json:"currency"`
		ReturnURL  interface{} `json:"return_url"`
		CancelURL  interface{} `json:"cancel_url"`
	}{
		Options:    opt.Options,
		CustomerID: s.CustomerID,
		Name:       s.Name,
		Currency:   s.Currency,
		ReturnURL:  s.ReturnURL,
		CancelURL:  s.CancelURL,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/authorization-requests"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
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

	payload.AuthorizationRequest.SetClient(s.client)
	return payload.AuthorizationRequest, nil
}

// AuthorizationRequestFindParameters is the structure representing the
// additional parameters used to call AuthorizationRequest.Find
type AuthorizationRequestFindParameters struct {
	*Options
	*AuthorizationRequest
}

// Find allows you to find an authorization request by its ID.
func (s AuthorizationRequest) Find(authorizationRequestID string, options ...AuthorizationRequestFindParameters) (*AuthorizationRequest, error) {
	if s.client == nil {
		panic("Please use the client.NewAuthorizationRequest() method to create a new AuthorizationRequest object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AuthorizationRequestFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.AuthorizationRequest)

	type Response struct {
		AuthorizationRequest *AuthorizationRequest `json:"authorization_request"`
		HasMore              bool                  `json:"has_more"`
		Success              bool                  `json:"success"`
		Message              string                `json:"message"`
		Code                 string                `json:"error_type"`
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

	path := "/authorization-requests/" + url.QueryEscape(authorizationRequestID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
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

	payload.AuthorizationRequest.SetClient(s.client)
	return payload.AuthorizationRequest, nil
}

// dummyAuthorizationRequest is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyAuthorizationRequest() {
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
