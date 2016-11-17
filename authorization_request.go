package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// AuthorizationRequest represents the AuthorizationRequest API object
type AuthorizationRequest struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the authorization
	ID string `json:"id"`
	// Project is the project to which the authorization request belongs
	Project *Project `json:"project"`
	// Customer is the customer linked to the authorization request
	Customer *Customer `json:"customer"`
	// Token is the token linked to the authorization request, once authorized
	Token *Token `json:"token"`
	// URL is the uRL to which you may redirect your customer to proceed with the authorization
	URL string `json:"url"`
	// Authorized is the whether or not the authorization request was authorized
	Authorized bool `json:"authorized"`
	// Name is the name of the authorization
	Name string `json:"name"`
	// Currency is the currency of the authorization
	Currency string `json:"currency"`
	// ReturnURL is the uRL where the customer will be redirected upon authorization
	ReturnURL string `json:"return_url"`
	// CancelURL is the uRL where the customer will be redirected if the authorization was canceled
	CancelURL string `json:"cancel_url"`
	// Custom is the custom variable passed along in the events/webhooks
	Custom string `json:"custom"`
	// Sandbox is the define whether or not the authorization is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt is the date at which the authorization was created
	CreatedAt time.Time `json:"created_at"`
}

// Customer allows you to get the customer linked to the authorization request.
func (s AuthorizationRequest) Customer(options ...Options) (*Customer, error) {
	if s.Client == nil {
		panic("Please use the client.NewAuthorizationRequest() method to create a new AuthorizationRequest object")
	}

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

	path := "/authorization-requests/" + url.QueryEscape(s.ID) + "/customers"

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return nil, erri
	}

	return payload.Customer, nil
}

// Create allows you to create a new authorization request for the given customer ID.
func (s AuthorizationRequest) Create(customerID string, options ...Options) (*AuthorizationRequest, error) {
	if s.Client == nil {
		panic("Please use the client.NewAuthorizationRequest() method to create a new AuthorizationRequest object")
	}

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
		"name":        s.Name,
		"currency":    s.Currency,
		"return_url":  s.ReturnURL,
		"cancel_url":  s.CancelURL,
		"custom":      s.Custom,
		"customer_id": customerID,
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

	path := "/authorization-requests"

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return nil, erri
	}

	return payload.AuthorizationRequest, nil
}

// Find allows you to find an authorization request by its ID.
func (s AuthorizationRequest) Find(authorizationRequestID string, options ...Options) (*AuthorizationRequest, error) {
	if s.Client == nil {
		panic("Please use the client.NewAuthorizationRequest() method to create a new AuthorizationRequest object")
	}

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

	path := "/authorization-requests/" + url.QueryEscape(authorizationRequestID) + ""

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

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
