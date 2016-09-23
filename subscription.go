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

// Subscriptions manages the Subscription struct
type Subscriptions struct {
	p *ProcessOut
}

type Subscription struct {
	// ID : ID of the subscription
	ID string `json:"id"`
	// Project : Project to which the subscription belongs
	Project *Project `json:"project"`
	// Customer : Customer linked to the subscription
	Customer *Customer `json:"customer"`
	// Token : Token linked to the subscription, once started
	Token *Token `json:"token"`
	// URL : URL to which you may redirect your customer to authorize the subscription
	URL string `json:"url"`
	// Name : Name of the subscription
	Name string `json:"name"`
	// Amount : Amount of the subscription
	Amount string `json:"amount"`
	// Currency : Currency of the subscription
	Currency string `json:"currency"`
	// Metadata : Metadata related to the subscription, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// ReturnURL : URL where the customer will be redirected when he activates the subscription
	ReturnURL string `json:"return_url"`
	// CancelURL : URL where the customer will be redirected when he canceles the subscription
	CancelURL string `json:"cancel_url"`
	// Interval : The subscription interval, formatted in the format "1d2w3m4y" (day, week, month, year)
	Interval string `json:"interval"`
	// TrialPeriod : The trial period. The customer will not be charged during this time span. Formatted in the format "1d2w3m4y" (day, week, month, year)
	TrialPeriod string `json:"trial_period"`
	// Activated : Weither or not the subscription is active
	Activated bool `json:"activated"`
	// Ended : Weither or not the subscription has ended (programmatically or canceled)
	Ended bool `json:"ended"`
	// EndedReason : Reason as to why the subscription ended
	EndedReason string `json:"ended_reason"`
	// Sandbox : Define whether or not the authorization is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the invoice was created
	CreatedAt time.Time `json:"created_at"`
}

// Customer : Get the customer owning the subscription.
func (s Subscriptions) Customer(subscription *Subscription, options ...Options) (*Customer, *Error) {
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
		Code     string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + "/customers"

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
	return &payload.Customer, nil
}

// CustomerAction : Get the customer action needed to be continue the subscription authorization flow on the given gateway.
func (s Subscriptions) CustomerAction(subscription *Subscription, gatewayConfigurationID string, options ...Options) (*CustomerAction, *Error) {
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
		Code           string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + "/gateway-configurations/" + url.QueryEscape(gatewayConfigurationID) + "/customer-action"

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
	return &payload.CustomerAction, nil
}

// Invoice : Get the invoice corresponding to the last iteration of the subscription.
func (s Subscriptions) Invoice(subscription *Subscription, options ...Options) (*Invoice, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Invoice `json:"invoice"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + "/invoices"

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
	return &payload.Invoice, nil
}

// Create : Create a new subscription for the given customer.
func (s Subscriptions) Create(subscription *Subscription, customerID string, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription `json:"subscription"`
		Success      bool   `json:"success"`
		Message      string `json:"message"`
		Code         string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":         subscription.Name,
		"amount":       subscription.Amount,
		"currency":     subscription.Currency,
		"metadata":     subscription.Metadata,
		"return_url":   subscription.ReturnURL,
		"cancel_url":   subscription.CancelURL,
		"interval":     subscription.Interval,
		"trial_period": subscription.TrialPeriod,
		"ended_reason": subscription.EndedReason,
		"customer_id":  customerID,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions"

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
	return &payload.Subscription, nil
}

// Find : Find a subscription by its ID.
func (s Subscriptions) Find(subscriptionID string, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription `json:"subscription"`
		Success      bool   `json:"success"`
		Message      string `json:"message"`
		Code         string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscriptionID) + ""

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
	return &payload.Subscription, nil
}

// Activate : Activate the subscription with the provided token.
func (s Subscriptions) Activate(subscription *Subscription, source string, options ...Options) *Error {
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
		"source": source,
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + ""

	req, err := http.NewRequest(
		"PUT",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return newError(err)
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
		return newError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return newError(err)
	}

	if !payload.Success {
		erri := newError(errors.New(payload.Message))
		erri.Code = payload.Code

		return erri
	}
	return nil
}

// End : End a subscription. The reason may be provided as well.
func (s Subscriptions) End(subscription *Subscription, reason string, options ...Options) *Error {
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
		"reason": reason,
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return newError(err)
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
		return newError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return newError(err)
	}

	if !payload.Success {
		erri := newError(errors.New(payload.Message))
		erri.Code = payload.Code

		return erri
	}
	return nil
}

// dummySubscription is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummySubscription() {
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
