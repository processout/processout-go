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
	// Plan : Plan used to create this subscription
	Plan *Plan `json:"plan"`
	// Customer : Customer linked to the subscription
	Customer *Customer `json:"customer"`
	// Token : Token used to capture payments on this subscription
	Token *Token `json:"token"`
	// URL : URL to which you may redirect your customer to activate the subscription
	URL string `json:"url"`
	// Name : Name of the subscription
	Name string `json:"name"`
	// Amount : Amount to be paid at each billing cycle of the subscription
	Amount string `json:"amount"`
	// Currency : Currency of the subscription
	Currency string `json:"currency"`
	// Metadata : Metadata related to the subscription, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Interval : The subscription interval, formatted in the format "1d2w3m4y" (day, week, month, year)
	Interval string `json:"interval"`
	// TrialEndAt : Date at which the subscription trial should end. Can be null to set no trial
	TrialEndAt time.Time `json:"trial_end_at"`
	// Activated : Whether or not the subscription was activated. This field does not take into account whether or not the subscription was canceled. Used the active field to know if the subscription is currently active
	Activated bool `json:"activated"`
	// Active : Whether or not the subscription is currently active (ie activated and not cancelled)
	Active bool `json:"active"`
	// Canceled : Whether or not the subscription was canceled. The cancellation reason can be found in the cancellation_reason field
	Canceled bool `json:"canceled"`
	// CancellationReason : Reason as to why the subscription was cancelled
	CancellationReason string `json:"cancellation_reason"`
	// PendingCancellation : Wheither or not the subscription is pending cancellation (meaning a cancel_at date was set)
	PendingCancellation bool `json:"pending_cancellation"`
	// CancelAt : Date at which the subscription will automatically be canceled. Can be null
	CancelAt time.Time `json:"cancel_at"`
	// ReturnURL : URL where the customer will be redirected upon activation of the subscription
	ReturnURL string `json:"return_url"`
	// CancelURL : URL where the customer will be redirected if the subscription activation was canceled
	CancelURL string `json:"cancel_url"`
	// Sandbox : Define whether or not the subscription is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the subscription was created
	CreatedAt time.Time `json:"created_at"`
	// ActivatedAt : Date at which the subscription was activated. Null if the subscription hasn't been activated yet
	ActivatedAt time.Time `json:"activated_at"`
	// IterateAt : Next iteration date, corresponding to the next billing cycle start date
	IterateAt time.Time `json:"iterate_at"`
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

// Discounts : Get the discounts applied to the subscription.
func (s Subscriptions) Discounts(subscription *Subscription, options ...Options) ([]*Discount, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Discounts []*Discount `json:"discounts"`

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

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + "/discounts"

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

	return payload.Discounts, nil
}

// Transactions : Get the subscriptions past transactions.
func (s Subscriptions) Transactions(subscription *Subscription, options ...Options) ([]*Transaction, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Transactions []*Transaction `json:"transactions"`

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

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + "/transactions"

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

	return payload.Transactions, nil
}

// All : Get all the subscriptions.
func (s Subscriptions) All(options ...Options) ([]*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscriptions []*Subscription `json:"subscriptions"`

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

	path := "/subscriptions"

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

	return payload.Subscriptions, nil
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
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"cancel_at":    subscription.CancelAt,
		"name":         subscription.Name,
		"amount":       subscription.Amount,
		"currency":     subscription.Currency,
		"metadata":     subscription.Metadata,
		"interval":     subscription.Interval,
		"trial_end_at": subscription.TrialEndAt,
		"return_url":   subscription.ReturnURL,
		"cancel_url":   subscription.CancelURL,
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

	return payload.Subscription, nil
}

// CreateFromPlan : Create a new subscription for the customer from the given plan ID.
func (s Subscriptions) CreateFromPlan(subscription *Subscription, customerID, planID string, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"cancel_at":    subscription.CancelAt,
		"name":         subscription.Name,
		"amount":       subscription.Amount,
		"currency":     subscription.Currency,
		"metadata":     subscription.Metadata,
		"interval":     subscription.Interval,
		"trial_end_at": subscription.TrialEndAt,
		"return_url":   subscription.ReturnURL,
		"cancel_url":   subscription.CancelURL,
		"customer_id":  customerID,
		"plan_id":      planID,
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

	return payload.Subscription, nil
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
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
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

	return payload.Subscription, nil
}

// Update : Update the subscription.
func (s Subscriptions) Update(subscription *Subscription, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"trial_end_at": subscription.TrialEndAt,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + ""

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

	return payload.Subscription, nil
}

// UpdatePlan : Update the subscription's plan.
func (s Subscriptions) UpdatePlan(subscription *Subscription, planID, prorate bool, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"plan_id": planID,
		"prorate": prorate,
		"expand":  opt.Expand,
		"filter":  opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + ""

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

	return payload.Subscription, nil
}

// ApplySource : Apply a source to the subscription to activate or update the subscription's source.
func (s Subscriptions) ApplySource(subscription *Subscription, source string, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"source": source,
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + ""

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

	return payload.Subscription, nil
}

// Cancel : Cancel a subscription. The reason may be provided as well.
func (s Subscriptions) Cancel(subscription *Subscription, cancellationReason string, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"cancellation_reason": cancellationReason,
		"expand":              opt.Expand,
		"filter":              opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
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

	return payload.Subscription, nil
}

// CancelAt : Schedule the cancellation of the subscription. The reason may be provided as well.
func (s Subscriptions) CancelAt(subscription *Subscription, cancelAt, cancellationReason string, options ...Options) (*Subscription, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"cancel_at":           cancelAt,
		"cancellation_reason": cancellationReason,
		"expand":              opt.Expand,
		"filter":              opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/subscriptions/" + url.QueryEscape(subscription.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
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

	return payload.Subscription, nil
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
