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

// Plan represents the Plan API object
type Plan struct {
	// ID is the iD of the plan
	ID *string `json:"id,omitempty"`
	// Project is the project to which the plan belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the plan belongs
	ProjectID *string `json:"project_id,omitempty"`
	// URL is the uRL to which you may redirect your customer to activate the subscription plan
	URL *string `json:"url,omitempty"`
	// Name is the name of the plan
	Name *string `json:"name,omitempty"`
	// Amount is the amount of the plan
	Amount *string `json:"amount,omitempty"`
	// Currency is the currency of the plan
	Currency *string `json:"currency,omitempty"`
	// Metadata is the metadata related to the plan, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Interval is the the plan interval, formatted in the format "1d2w3m4y" (day, week, month, year)
	Interval *string `json:"interval,omitempty"`
	// TrialPeriod is the the trial period. The customer will not be charged during this time span. Formatted in the format "1d2w3m4y" (day, week, month, year)
	TrialPeriod *string `json:"trial_period,omitempty"`
	// ReturnURL is the uRL where the customer will be redirected when activating the subscription created using this plan
	ReturnURL *string `json:"return_url,omitempty"`
	// CancelURL is the uRL where the customer will be redirected when cancelling the subscription created using this plan
	CancelURL *string `json:"cancel_url,omitempty"`
	// Sandbox is the define whether or not the plan is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the plan was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Plan) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Plan object and its
// children
func (s *Plan) SetClient(c *ProcessOut) *Plan {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Plan) Prefill(c *Plan) *Plan {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.URL = c.URL
	s.Name = c.Name
	s.Amount = c.Amount
	s.Currency = c.Currency
	s.Metadata = c.Metadata
	s.Interval = c.Interval
	s.TrialPeriod = c.TrialPeriod
	s.ReturnURL = c.ReturnURL
	s.CancelURL = c.CancelURL
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// PlanAllParameters is the structure representing the
// additional parameters used to call Plan.All
type PlanAllParameters struct {
	*Options
	*Plan
}

// All allows you to get all the plans.
func (s Plan) All(options ...PlanAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewPlan() method to create a new Plan object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PlanAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Plan)

	type Response struct {
		Plans []*Plan `json:"plans"`

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

	path := "/plans"

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

	plansList := []Identifiable{}
	for _, o := range payload.Plans {
		plansList = append(plansList, o.SetClient(s.client))
	}
	plansIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    plansList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"plans"`
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
	return plansIterator, nil
}

// PlanCreateParameters is the structure representing the
// additional parameters used to call Plan.Create
type PlanCreateParameters struct {
	*Options
	*Plan
}

// Create allows you to create a new plan.
func (s Plan) Create(options ...PlanCreateParameters) (*Plan, error) {
	if s.client == nil {
		panic("Please use the client.NewPlan() method to create a new Plan object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PlanCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Plan)

	type Response struct {
		Plan    *Plan  `json:"plan"`
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		ID          interface{} `json:"id"`
		Name        interface{} `json:"name"`
		Amount      interface{} `json:"amount"`
		Currency    interface{} `json:"currency"`
		Interval    interface{} `json:"interval"`
		TrialPeriod interface{} `json:"trial_period"`
		Metadata    interface{} `json:"metadata"`
		ReturnURL   interface{} `json:"return_url"`
		CancelURL   interface{} `json:"cancel_url"`
	}{
		Options:     opt.Options,
		ID:          s.ID,
		Name:        s.Name,
		Amount:      s.Amount,
		Currency:    s.Currency,
		Interval:    s.Interval,
		TrialPeriod: s.TrialPeriod,
		Metadata:    s.Metadata,
		ReturnURL:   s.ReturnURL,
		CancelURL:   s.CancelURL,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/plans"

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

	payload.Plan.SetClient(s.client)
	return payload.Plan, nil
}

// PlanFindParameters is the structure representing the
// additional parameters used to call Plan.Find
type PlanFindParameters struct {
	*Options
	*Plan
}

// Find allows you to find a plan by its ID.
func (s Plan) Find(planID string, options ...PlanFindParameters) (*Plan, error) {
	if s.client == nil {
		panic("Please use the client.NewPlan() method to create a new Plan object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PlanFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Plan)

	type Response struct {
		Plan    *Plan  `json:"plan"`
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

	path := "/plans/" + url.QueryEscape(planID) + ""

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

	payload.Plan.SetClient(s.client)
	return payload.Plan, nil
}

// PlanSaveParameters is the structure representing the
// additional parameters used to call Plan.Save
type PlanSaveParameters struct {
	*Options
	*Plan
}

// Save allows you to save the updated plan attributes. This action won't affect subscriptions already linked to this plan.
func (s Plan) Save(options ...PlanSaveParameters) (*Plan, error) {
	if s.client == nil {
		panic("Please use the client.NewPlan() method to create a new Plan object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PlanSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Plan)

	type Response struct {
		Plan    *Plan  `json:"plan"`
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Name        interface{} `json:"name"`
		TrialPeriod interface{} `json:"trial_period"`
		Metadata    interface{} `json:"metadata"`
		ReturnURL   interface{} `json:"return_url"`
		CancelURL   interface{} `json:"cancel_url"`
	}{
		Options:     opt.Options,
		Name:        s.Name,
		TrialPeriod: s.TrialPeriod,
		Metadata:    s.Metadata,
		ReturnURL:   s.ReturnURL,
		CancelURL:   s.CancelURL,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/plans/" + url.QueryEscape(*s.ID) + ""

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

	payload.Plan.SetClient(s.client)
	return payload.Plan, nil
}

// PlanEndParameters is the structure representing the
// additional parameters used to call Plan.End
type PlanEndParameters struct {
	*Options
	*Plan
}

// End allows you to delete a plan. Subscriptions linked to this plan won't be affected.
func (s Plan) End(options ...PlanEndParameters) error {
	if s.client == nil {
		panic("Please use the client.NewPlan() method to create a new Plan object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PlanEndParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Plan)

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

	path := "/plans/" + url.QueryEscape(*s.ID) + ""

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
		g io.Reader
	}
	errors.New(nil, "", "")
}
