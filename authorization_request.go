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
	// Token : Token linked to the authorization request, once authorized
	Token *Token `json:"token"`
	// URL : URL to which you may redirect your customer to proceed with the authorization
	URL string `json:"url"`
	// Authorized : Whether or not the authorization request was authorized
	Authorized bool `json:"authorized"`
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
func (s AuthorizationRequests) Customer(authorizationRequest *AuthorizationRequest, options ...Options) (*Customer, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Customer *Customer `json:"customer"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/authorization-requests/" + url.QueryEscape(authorizationRequest.ID) + "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, newError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, newError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, newError(err)
	}

	if !payload.Success {
		erri := newError(errors.New(payload.Message))
		erri.Code = payload.Code

		return nil, erri
	}

	return payload.Customer, nil
}

// Create : Create a new authorization request for the given customer ID.
func (s AuthorizationRequests) Create(authorizationRequest *AuthorizationRequest, customerID string, options ...Options) (*AuthorizationRequest, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		AuthorizationRequest *AuthorizationRequest `json:"authorization_request"`
		Success              bool                  `json:"success"`
		Message              string                `json:"message"`
		Code                 string                `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":        authorizationRequest.Name,
		"currency":    authorizationRequest.Currency,
		"return_url":  authorizationRequest.ReturnURL,
		"cancel_url":  authorizationRequest.CancelURL,
		"custom":      authorizationRequest.Custom,
		"customer_id": customerID,
		"expand":      opt.Expand,
		"filter":      opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/authorization-requests"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, newError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, newError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, newError(err)
	}

	if !payload.Success {
		erri := newError(errors.New(payload.Message))
		erri.Code = payload.Code

		return nil, erri
	}

	return payload.AuthorizationRequest, nil
}

// Find : Find an authorization request by its ID.
func (s AuthorizationRequests) Find(authorizationRequestID string, options ...Options) (*AuthorizationRequest, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		AuthorizationRequest *AuthorizationRequest `json:"authorization_request"`
		Success              bool                  `json:"success"`
		Message              string                `json:"message"`
		Code                 string                `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/authorization-requests/" + url.QueryEscape(authorizationRequestID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, newError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, newError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, newError(err)
	}

	if !payload.Success {
		erri := newError(errors.New(payload.Message))
		erri.Code = payload.Code

		return nil, erri
	}

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
	}
	errors.New("")
}
