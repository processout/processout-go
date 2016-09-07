package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// AuthorizationRequests manages the AuthorizationRequest struct
type AuthorizationRequests struct {
	p *ProcessOut
}

type AuthorizationRequest struct {
	// ID : ID of the authorization
	ID string `json:"id"`
	// Project : Project to which the authorization request belongs
	Project *Project `json:"project"`
	// Customer : Customer linked to the authorization request
	Customer *Customer `json:"customer"`
	// URL : URL to which you may redirect your customer to proceed with the authorization
	URL string `json:"url"`
	// Name : Name of the authorization
	Name string `json:"name"`
	// Currency : Currency of the authorization
	Currency string `json:"currency"`
	// ReturnURL : URL where the customer will be redirected upon authorization
	ReturnURL string `json:"return_url"`
	// CancelURL : URL where the customer will be redirected if the authorization was canceled
	CancelURL string `json:"cancel_url"`
	// Custom : Custom variable passed along in the events/webhooks
	Custom string `json:"custom"`
	// Sandbox : Define whether or not the authorization is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the authorization was created
	CreatedAt time.Time `json:"created_at"`
}

// Customer : Get the customer linked to the authorization request.
func (s AuthorizationRequests) Customer(authorizationRequest *AuthorizationRequest, options ...Options) (*Customer, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/authorization-requests/" + url.QueryEscape(authorizationRequest.ID) + "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Customer, nil
}

// CustomerAction : Get the customer action needed to be continue the token authorization flow on the given gateway.
func (s AuthorizationRequests) CustomerAction(authorizationRequest *AuthorizationRequest, gatewayConfigurationID string, options ...Options) (*CustomerAction, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		CustomerAction `json:"customer_action"`
		Success        bool   `json:"success"`
		Message        string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/authorization-requests/" + url.QueryEscape(authorizationRequest.ID) + "/gateway-configurations/" + url.QueryEscape(gatewayConfigurationID) + "/customer-action"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.CustomerAction, nil
}

// Create : Create a new authorization request for the given customer ID.
func (s AuthorizationRequests) Create(authorizationRequest *AuthorizationRequest, customerID string, options ...Options) (*AuthorizationRequest, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		AuthorizationRequest `json:"authorization_request"`
		Success              bool   `json:"success"`
		Message              string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":        authorizationRequest.Name,
		"currency":    authorizationRequest.Currency,
		"return_url":  authorizationRequest.ReturnURL,
		"cancel_url":  authorizationRequest.CancelURL,
		"custom":      authorizationRequest.Custom,
		"customer_id": customerID,
		"expand":      opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/authorization-requests"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.AuthorizationRequest, nil
}

// Find : Find an authorization request by its ID.
func (s AuthorizationRequests) Find(authorizationRequestID string, options ...Options) (*AuthorizationRequest, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		AuthorizationRequest `json:"authorization_request"`
		Success              bool   `json:"success"`
		Message              string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/authorization-requests/" + url.QueryEscape(authorizationRequestID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.AuthorizationRequest, nil
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
	}
	errors.New("")
}
