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

// Plans manages the Plan struct
type Plans struct {
	p *ProcessOut
}

type Plan struct {
	// ID : ID of the plan
	ID string `json:"id"`
	// Project : Project to which the plan belongs
	Project *Project `json:"project"`
	// Name : Name of the plan
	Name string `json:"name"`
	// Amount : Amount of the plan
	Amount string `json:"amount"`
	// Currency : Currency of the plan
	Currency string `json:"currency"`
	// Metadata : Metadata related to the plan, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Interval : The plan interval, formatted in the format "1d2w3m4y" (day, week, month, year)
	Interval string `json:"interval"`
	// TrialPeriod : The trial period. The customer will not be charged during this time span. Formatted in the format "1d2w3m4y" (day, week, month, year)
	TrialPeriod string `json:"trial_period"`
	// ReturnURL : URL where the customer will be redirected when he activates the subscription created using this plan
	ReturnURL string `json:"return_url"`
	// CancelURL : URL where the customer will be redirected when he cancelling the subscription created using this plan
	CancelURL string `json:"cancel_url"`
	// Sandbox : Define whether or not the plan is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the plan was created
	CreatedAt time.Time `json:"created_at"`
}

// All : Get all the plans.
func (s Plans) All(options ...Options) ([]*Plan, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Plans []*Plan `json:"plans"`

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

	path := "/plans"

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

	return payload.Plans, nil
}

// Create : Create a new plan.
func (s Plans) Create(plan *Plan, options ...Options) (*Plan, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Plan    *Plan  `json:"plan"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"id":           plan.ID,
		"name":         plan.Name,
		"amount":       plan.Amount,
		"currency":     plan.Currency,
		"interval":     plan.Interval,
		"trial_period": plan.TrialPeriod,
		"metadata":     plan.Metadata,
		"return_url":   plan.ReturnURL,
		"cancel_url":   plan.CancelURL,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/plans"

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

	return payload.Plan, nil
}

// Find : Find a plan by its ID.
func (s Plans) Find(planID string, options ...Options) (*Plan, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Plan    *Plan  `json:"plan"`
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

	path := "/plans/" + url.QueryEscape(planID) + ""

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

	return payload.Plan, nil
}

// Update : Update the plan. This action won't affect subscriptions already linked to this plan.
func (s Plans) Update(plan *Plan, options ...Options) (*Plan, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Plan    *Plan  `json:"plan"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":         plan.Name,
		"trial_period": plan.TrialPeriod,
		"metadata":     plan.Metadata,
		"return_url":   plan.ReturnURL,
		"cancel_url":   plan.CancelURL,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/plans/" + url.QueryEscape(plan.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	return payload.Plan, nil
}

// End : Delete a plan. Subscriptions linked to this plan won't be affected.
func (s Plans) End(plan *Plan, options ...Options) *Error {
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
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return newError(err)
	}

	path := "/plans/" + url.QueryEscape(plan.ID) + ""

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
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
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

// dummyPlan is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPlan() {
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
