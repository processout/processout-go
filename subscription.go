package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Subscription represents the Subscription API object
type Subscription struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the subscription
	ID string `json:"id"`
	// Project is the project to which the subscription belongs
	Project *Project `json:"project"`
	// Plan is the plan used to create this subscription
	Plan *Plan `json:"plan"`
	// Customer is the customer linked to the subscription
	Customer *Customer `json:"customer"`
	// Token is the token used to capture payments on this subscription
	Token *Token `json:"token"`
	// URL is the uRL to which you may redirect your customer to activate the subscription
	URL string `json:"url"`
	// Name is the name of the subscription
	Name string `json:"name"`
	// Amount is the amount to be paid at each billing cycle of the subscription
	Amount string `json:"amount"`
	// Currency is the currency of the subscription
	Currency string `json:"currency"`
	// Metadata is the metadata related to the subscription, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Interval is the the subscription interval, formatted in the format "1d2w3m4y" (day, week, month, year)
	Interval string `json:"interval"`
	// TrialEndAt is the date at which the subscription trial should end. Can be null to set no trial
	TrialEndAt time.Time `json:"trial_end_at"`
	// Activated is the whether or not the subscription was activated. This field does not take into account whether or not the subscription was canceled. Used the active field to know if the subscription is currently active
	Activated bool `json:"activated"`
	// Active is the whether or not the subscription is currently active (ie activated and not cancelled)
	Active bool `json:"active"`
	// Canceled is the whether or not the subscription was canceled. The cancellation reason can be found in the cancellation_reason field
	Canceled bool `json:"canceled"`
	// CancellationReason is the reason as to why the subscription was cancelled
	CancellationReason string `json:"cancellation_reason"`
	// PendingCancellation is the wheither or not the subscription is pending cancellation (meaning a cancel_at date was set)
	PendingCancellation bool `json:"pending_cancellation"`
	// CancelAt is the date at which the subscription will automatically be canceled. Can be null
	CancelAt time.Time `json:"cancel_at"`
	// ReturnURL is the uRL where the customer will be redirected upon activation of the subscription
	ReturnURL string `json:"return_url"`
	// CancelURL is the uRL where the customer will be redirected if the subscription activation was canceled
	CancelURL string `json:"cancel_url"`
	// Sandbox is the define whether or not the subscription is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt is the date at which the subscription was created
	CreatedAt time.Time `json:"created_at"`
	// ActivatedAt is the date at which the subscription was activated. Null if the subscription hasn't been activated yet
	ActivatedAt time.Time `json:"activated_at"`
	// IterateAt is the next iteration date, corresponding to the next billing cycle start date
	IterateAt time.Time `json:"iterate_at"`
}

// GetCustomer allows you to get the customer owning the subscription.
func (s Subscription) GetCustomer(options ...Options) (*Customer, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
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

	path := "/subscriptions/" + url.QueryEscape(s.ID) + "/customers"

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

// GetDiscounts allows you to get the discounts applied to the subscription.
func (s Subscription) GetDiscounts(options ...Options) ([]*Discount, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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

	path := "/subscriptions/" + url.QueryEscape(s.ID) + "/discounts"

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

	return payload.Discounts, nil
}

// FindDiscount allows you to find a subscription's discount by its ID.
func (s Subscription) FindDiscount(discountID string, options ...Options) (*Discount, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Discount *Discount `json:"discount"`
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

	path := "/subscriptions/" + url.QueryEscape(s.ID) + "/discounts/" + url.QueryEscape(discountID) + ""

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

	return payload.Discount, nil
}

// RemoveDiscount allows you to remove a discount applied to a subscription.
func (s Subscription) RemoveDiscount(discountID string, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscription *Subscription `json:"discount"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
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

	path := "/subscriptions/" + url.QueryEscape(s.ID) + "/discounts/" + url.QueryEscape(discountID) + ""

	req, err := http.NewRequest(
		"DELETE",
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

	return payload.Subscription, nil
}

// GetTransactions allows you to get the subscriptions past transactions.
func (s Subscription) GetTransactions(options ...Options) ([]*Transaction, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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

	path := "/subscriptions/" + url.QueryEscape(s.ID) + "/transactions"

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

	return payload.Transactions, nil
}

// All allows you to get all the subscriptions.
func (s Subscription) All(options ...Options) ([]*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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

	path := "/subscriptions"

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

	return payload.Subscriptions, nil
}

// Create allows you to create a new subscription for the given customer.
func (s Subscription) Create(customerID string, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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
		"cancel_at":    s.CancelAt,
		"name":         s.Name,
		"amount":       s.Amount,
		"currency":     s.Currency,
		"metadata":     s.Metadata,
		"interval":     s.Interval,
		"trial_end_at": s.TrialEndAt,
		"return_url":   s.ReturnURL,
		"cancel_url":   s.CancelURL,
		"customer_id":  customerID,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
		"limit":        opt.Limit,
		"page":         opt.Page,
		"end_before":   opt.EndBefore,
		"start_after":  opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions"

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

	return payload.Subscription, nil
}

// CreateFromPlan allows you to create a new subscription for the customer from the given plan ID.
func (s Subscription) CreateFromPlan(customerID, planID string, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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
		"cancel_at":    s.CancelAt,
		"name":         s.Name,
		"amount":       s.Amount,
		"currency":     s.Currency,
		"metadata":     s.Metadata,
		"interval":     s.Interval,
		"trial_end_at": s.TrialEndAt,
		"return_url":   s.ReturnURL,
		"cancel_url":   s.CancelURL,
		"customer_id":  customerID,
		"plan_id":      planID,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
		"limit":        opt.Limit,
		"page":         opt.Page,
		"end_before":   opt.EndBefore,
		"start_after":  opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions"

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

	return payload.Subscription, nil
}

// Find allows you to find a subscription by its ID.
func (s Subscription) Find(subscriptionID string, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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

	path := "/subscriptions/" + url.QueryEscape(subscriptionID) + ""

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

	return payload.Subscription, nil
}

// Update allows you to update the subscription.
func (s Subscription) Update(prorate bool, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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
		"trial_end_at": s.TrialEndAt,
		"prorate":      prorate,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
		"limit":        opt.Limit,
		"page":         opt.Page,
		"end_before":   opt.EndBefore,
		"start_after":  opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	return payload.Subscription, nil
}

// UpdatePlan allows you to update the subscription's plan.
func (s Subscription) UpdatePlan(planID, prorate bool, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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
		"plan_id":     planID,
		"prorate":     prorate,
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

	path := "/subscriptions/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	return payload.Subscription, nil
}

// ApplySource allows you to apply a source to the subscription to activate or update the subscription's source.
func (s Subscription) ApplySource(source string, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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
		"source":      source,
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

	path := "/subscriptions/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	return payload.Subscription, nil
}

// Cancel allows you to cancel a subscription. The reason may be provided as well.
func (s Subscription) Cancel(cancellationReason string, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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
		"limit":               opt.Limit,
		"page":                opt.Page,
		"end_before":          opt.EndBefore,
		"start_after":         opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
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

	return payload.Subscription, nil
}

// CancelAtDate allows you to schedule the cancellation of the subscription. The reason may be provided as well.
func (s Subscription) CancelAtDate(cancelAt, cancellationReason string, options ...Options) (*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}

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
		"limit":               opt.Limit,
		"page":                opt.Page,
		"end_before":          opt.EndBefore,
		"start_after":         opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
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
	errors.New(nil, "", "")
}
