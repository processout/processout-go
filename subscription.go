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

// Subscription represents the Subscription API object
type Subscription struct {
	// ID is the iD of the subscription
	ID *string `json:"id,omitempty"`
	// Project is the project to which the subscription belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the subscription belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Plan is the plan linked to this subscription, if any
	Plan *Plan `json:"plan,omitempty"`
	// PlanID is the iD of the plan linked to this subscription, if any
	PlanID *string `json:"plan_id,omitempty"`
	// Discounts is the list of the subscription discounts
	Discounts *[]*Discount `json:"discounts,omitempty"`
	// Addons is the list of the subscription addons
	Addons *[]*Addon `json:"addons,omitempty"`
	// Transactions is the list of the subscription transactions
	Transactions *[]*Transaction `json:"transactions,omitempty"`
	// Customer is the customer linked to the subscription
	Customer *Customer `json:"customer,omitempty"`
	// CustomerID is the iD of the customer linked to the subscription
	CustomerID *string `json:"customer_id,omitempty"`
	// Token is the token used to capture payments on this subscription
	Token *Token `json:"token,omitempty"`
	// TokenID is the iD of the token used to capture payments on this subscription
	TokenID *string `json:"token_id,omitempty"`
	// URL is the uRL to which you may redirect your customer to activate the subscription
	URL *string `json:"url,omitempty"`
	// Name is the name of the subscription
	Name *string `json:"name,omitempty"`
	// Amount is the base amount of the subscription
	Amount *string `json:"amount,omitempty"`
	// BillableAmount is the amount to be paid at each billing cycle of the subscription
	BillableAmount *string `json:"billable_amount,omitempty"`
	// DiscountedAmount is the amount discounted by discounts applied to the subscription
	DiscountedAmount *string `json:"discounted_amount,omitempty"`
	// AddonsAmount is the amount applied on top of the subscription base price with addons
	AddonsAmount *string `json:"addons_amount,omitempty"`
	// Currency is the currency of the subscription
	Currency *string `json:"currency,omitempty"`
	// Metadata is the metadata related to the subscription, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Interval is the the subscription interval, formatted in the format "1d2w3m4y" (day, week, month, year)
	Interval *string `json:"interval,omitempty"`
	// TrialEndAt is the date at which the subscription trial should end. Can be null to set no trial
	TrialEndAt *time.Time `json:"trial_end_at,omitempty"`
	// Activated is the whether or not the subscription was activated. This field does not take into account whether or not the subscription was canceled. Use the active field to know if the subscription is currently active
	Activated *bool `json:"activated,omitempty"`
	// Active is the whether or not the subscription is currently active (ie activated and not cancelled)
	Active *bool `json:"active,omitempty"`
	// CancelAt is the date at which the subscription will automatically be canceled. Can be null
	CancelAt *time.Time `json:"cancel_at,omitempty"`
	// Canceled is the whether or not the subscription was canceled. The cancellation reason can be found in the cancellation_reason field
	Canceled *bool `json:"canceled,omitempty"`
	// CancellationReason is the reason as to why the subscription was cancelled
	CancellationReason *string `json:"cancellation_reason,omitempty"`
	// PendingCancellation is the whether or not the subscription is pending cancellation (meaning a cancel_at date was set)
	PendingCancellation *bool `json:"pending_cancellation,omitempty"`
	// ReturnURL is the uRL where the customer will be redirected upon activation of the subscription
	ReturnURL *string `json:"return_url,omitempty"`
	// CancelURL is the uRL where the customer will be redirected if the subscription activation was canceled
	CancelURL *string `json:"cancel_url,omitempty"`
	// UnpaidState is the when the subscription has unpaid invoices, defines the dunning logic of the subscription (as specified in the project settings)
	UnpaidState *string `json:"unpaid_state,omitempty"`
	// Sandbox is the define whether or not the subscription is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the subscription was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ActivatedAt is the date at which the subscription was activated. Null if the subscription hasn't been activated yet
	ActivatedAt *time.Time `json:"activated_at,omitempty"`
	// IterateAt is the next iteration date, corresponding to the next billing cycle start date
	IterateAt *time.Time `json:"iterate_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Subscription) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Subscription object and its
// children
func (s *Subscription) SetClient(c *ProcessOut) *Subscription {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Plan != nil {
		s.Plan.SetClient(c)
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
func (s *Subscription) Prefill(c *Subscription) *Subscription {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Plan = c.Plan
	s.PlanID = c.PlanID
	s.Discounts = c.Discounts
	s.Addons = c.Addons
	s.Transactions = c.Transactions
	s.Customer = c.Customer
	s.CustomerID = c.CustomerID
	s.Token = c.Token
	s.TokenID = c.TokenID
	s.URL = c.URL
	s.Name = c.Name
	s.Amount = c.Amount
	s.BillableAmount = c.BillableAmount
	s.DiscountedAmount = c.DiscountedAmount
	s.AddonsAmount = c.AddonsAmount
	s.Currency = c.Currency
	s.Metadata = c.Metadata
	s.Interval = c.Interval
	s.TrialEndAt = c.TrialEndAt
	s.Activated = c.Activated
	s.Active = c.Active
	s.CancelAt = c.CancelAt
	s.Canceled = c.Canceled
	s.CancellationReason = c.CancellationReason
	s.PendingCancellation = c.PendingCancellation
	s.ReturnURL = c.ReturnURL
	s.CancelURL = c.CancelURL
	s.UnpaidState = c.UnpaidState
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt
	s.ActivatedAt = c.ActivatedAt
	s.IterateAt = c.IterateAt

	return s
}

// SubscriptionFetchAddonsParameters is the structure representing the
// additional parameters used to call Subscription.FetchAddons
type SubscriptionFetchAddonsParameters struct {
	*Options
	*Subscription
}

// FetchAddons allows you to get the addons applied to the subscription.
func (s Subscription) FetchAddons(options ...SubscriptionFetchAddonsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionFetchAddonsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Addons []*Addon `json:"addons"`

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

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/addons"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	addonsList := []Identifiable{}
	for _, o := range payload.Addons {
		addonsList = append(addonsList, o.SetClient(s.client))
	}
	addonsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    addonsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"addons"`
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
		hasMorePrev: true,
	}
	return addonsIterator, nil
}

// SubscriptionFindAddonParameters is the structure representing the
// additional parameters used to call Subscription.FindAddon
type SubscriptionFindAddonParameters struct {
	*Options
	*Subscription
}

// FindAddon allows you to find a subscription's addon by its ID.
func (s Subscription) FindAddon(addonID string, options ...SubscriptionFindAddonParameters) (*Addon, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionFindAddonParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Addon   *Addon `json:"addon"`
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

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/addons/" + url.QueryEscape(addonID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Addon.SetClient(s.client)
	return payload.Addon, nil
}

// SubscriptionDeleteAddonParameters is the structure representing the
// additional parameters used to call Subscription.DeleteAddon
type SubscriptionDeleteAddonParameters struct {
	*Options
	*Subscription
	Prorate       interface{} `json:"prorate"`
	ProrationDate interface{} `json:"proration_date"`
	Preview       interface{} `json:"preview"`
}

// DeleteAddon allows you to delete an addon applied to a subscription.
func (s Subscription) DeleteAddon(addonID string, options ...SubscriptionDeleteAddonParameters) error {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionDeleteAddonParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Prorate       interface{} `json:"prorate"`
		ProrationDate interface{} `json:"proration_date"`
		Preview       interface{} `json:"preview"`
	}{
		Options:       opt.Options,
		Prorate:       opt.Prorate,
		ProrationDate: opt.ProrationDate,
		Preview:       opt.Preview,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/addons/" + url.QueryEscape(addonID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
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

// SubscriptionFetchCustomerParameters is the structure representing the
// additional parameters used to call Subscription.FetchCustomer
type SubscriptionFetchCustomerParameters struct {
	*Options
	*Subscription
}

// FetchCustomer allows you to get the customer owning the subscription.
func (s Subscription) FetchCustomer(options ...SubscriptionFetchCustomerParameters) (*Customer, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionFetchCustomerParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

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

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Customer.SetClient(s.client)
	return payload.Customer, nil
}

// SubscriptionFetchDiscountsParameters is the structure representing the
// additional parameters used to call Subscription.FetchDiscounts
type SubscriptionFetchDiscountsParameters struct {
	*Options
	*Subscription
}

// FetchDiscounts allows you to get the discounts applied to the subscription.
func (s Subscription) FetchDiscounts(options ...SubscriptionFetchDiscountsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionFetchDiscountsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Discounts []*Discount `json:"discounts"`

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

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/discounts"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	discountsList := []Identifiable{}
	for _, o := range payload.Discounts {
		discountsList = append(discountsList, o.SetClient(s.client))
	}
	discountsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    discountsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"discounts"`
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
		hasMorePrev: true,
	}
	return discountsIterator, nil
}

// SubscriptionFindDiscountParameters is the structure representing the
// additional parameters used to call Subscription.FindDiscount
type SubscriptionFindDiscountParameters struct {
	*Options
	*Subscription
}

// FindDiscount allows you to find a subscription's discount by its ID.
func (s Subscription) FindDiscount(discountID string, options ...SubscriptionFindDiscountParameters) (*Discount, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionFindDiscountParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Discount *Discount `json:"discount"`
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

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/discounts/" + url.QueryEscape(discountID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Discount.SetClient(s.client)
	return payload.Discount, nil
}

// SubscriptionDeleteDiscountParameters is the structure representing the
// additional parameters used to call Subscription.DeleteDiscount
type SubscriptionDeleteDiscountParameters struct {
	*Options
	*Subscription
}

// DeleteDiscount allows you to delete a discount applied to a subscription.
func (s Subscription) DeleteDiscount(discountID string, options ...SubscriptionDeleteDiscountParameters) error {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionDeleteDiscountParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

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

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/discounts/" + url.QueryEscape(discountID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
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

// SubscriptionFetchTransactionsParameters is the structure representing the
// additional parameters used to call Subscription.FetchTransactions
type SubscriptionFetchTransactionsParameters struct {
	*Options
	*Subscription
}

// FetchTransactions allows you to get the subscriptions past transactions.
func (s Subscription) FetchTransactions(options ...SubscriptionFetchTransactionsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionFetchTransactionsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Transactions []*Transaction `json:"transactions"`

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

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + "/transactions"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	transactionsList := []Identifiable{}
	for _, o := range payload.Transactions {
		transactionsList = append(transactionsList, o.SetClient(s.client))
	}
	transactionsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    transactionsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"transactions"`
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
		hasMorePrev: true,
	}
	return transactionsIterator, nil
}

// SubscriptionAllParameters is the structure representing the
// additional parameters used to call Subscription.All
type SubscriptionAllParameters struct {
	*Options
	*Subscription
}

// All allows you to get all the subscriptions.
func (s Subscription) All(options ...SubscriptionAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Subscriptions []*Subscription `json:"subscriptions"`

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

	path := "/subscriptions"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	subscriptionsList := []Identifiable{}
	for _, o := range payload.Subscriptions {
		subscriptionsList = append(subscriptionsList, o.SetClient(s.client))
	}
	subscriptionsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    subscriptionsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"subscriptions"`
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
		hasMorePrev: true,
	}
	return subscriptionsIterator, nil
}

// SubscriptionCreateParameters is the structure representing the
// additional parameters used to call Subscription.Create
type SubscriptionCreateParameters struct {
	*Options
	*Subscription
	Source   interface{} `json:"source"`
	CouponID interface{} `json:"coupon_id"`
}

// Create allows you to create a new subscription for the given customer.
func (s Subscription) Create(options ...SubscriptionCreateParameters) (*Subscription, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	data := struct {
		*Options
		PlanID     interface{} `json:"plan_id"`
		CancelAt   interface{} `json:"cancel_at"`
		Name       interface{} `json:"name"`
		Amount     interface{} `json:"amount"`
		Currency   interface{} `json:"currency"`
		Metadata   interface{} `json:"metadata"`
		Interval   interface{} `json:"interval"`
		TrialEndAt interface{} `json:"trial_end_at"`
		CustomerID interface{} `json:"customer_id"`
		ReturnURL  interface{} `json:"return_url"`
		CancelURL  interface{} `json:"cancel_url"`
		Source     interface{} `json:"source"`
		CouponID   interface{} `json:"coupon_id"`
	}{
		Options:    opt.Options,
		PlanID:     s.PlanID,
		CancelAt:   s.CancelAt,
		Name:       s.Name,
		Amount:     s.Amount,
		Currency:   s.Currency,
		Metadata:   s.Metadata,
		Interval:   s.Interval,
		TrialEndAt: s.TrialEndAt,
		CustomerID: s.CustomerID,
		ReturnURL:  s.ReturnURL,
		CancelURL:  s.CancelURL,
		Source:     opt.Source,
		CouponID:   opt.CouponID,
	}

	body, err := json.Marshal(data)
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
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Subscription.SetClient(s.client)
	return payload.Subscription, nil
}

// SubscriptionFindParameters is the structure representing the
// additional parameters used to call Subscription.Find
type SubscriptionFindParameters struct {
	*Options
	*Subscription
}

// Find allows you to find a subscription by its ID.
func (s Subscription) Find(subscriptionID string, options ...SubscriptionFindParameters) (*Subscription, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
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

	path := "/subscriptions/" + url.QueryEscape(subscriptionID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Subscription.SetClient(s.client)
	return payload.Subscription, nil
}

// SubscriptionSaveParameters is the structure representing the
// additional parameters used to call Subscription.Save
type SubscriptionSaveParameters struct {
	*Options
	*Subscription
	CouponID      interface{} `json:"coupon_id"`
	Source        interface{} `json:"source"`
	Prorate       interface{} `json:"prorate"`
	ProrationDate interface{} `json:"proration_date"`
	Preview       interface{} `json:"preview"`
}

// Save allows you to save the updated subscription attributes.
func (s Subscription) Save(options ...SubscriptionSaveParameters) (*Subscription, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	data := struct {
		*Options
		PlanID        interface{} `json:"plan_id"`
		Name          interface{} `json:"name"`
		Amount        interface{} `json:"amount"`
		Interval      interface{} `json:"interval"`
		TrialEndAt    interface{} `json:"trial_end_at"`
		Metadata      interface{} `json:"metadata"`
		CouponID      interface{} `json:"coupon_id"`
		Source        interface{} `json:"source"`
		Prorate       interface{} `json:"prorate"`
		ProrationDate interface{} `json:"proration_date"`
		Preview       interface{} `json:"preview"`
	}{
		Options:       opt.Options,
		PlanID:        s.PlanID,
		Name:          s.Name,
		Amount:        s.Amount,
		Interval:      s.Interval,
		TrialEndAt:    s.TrialEndAt,
		Metadata:      s.Metadata,
		CouponID:      opt.CouponID,
		Source:        opt.Source,
		Prorate:       opt.Prorate,
		ProrationDate: opt.ProrationDate,
		Preview:       opt.Preview,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Subscription.SetClient(s.client)
	return payload.Subscription, nil
}

// SubscriptionCancelParameters is the structure representing the
// additional parameters used to call Subscription.Cancel
type SubscriptionCancelParameters struct {
	*Options
	*Subscription
	CancelAtEnd interface{} `json:"cancel_at_end"`
}

// Cancel allows you to cancel a subscription. The reason may be provided as well.
func (s Subscription) Cancel(options ...SubscriptionCancelParameters) (*Subscription, error) {
	if s.client == nil {
		panic("Please use the client.NewSubscription() method to create a new Subscription object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := SubscriptionCancelParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Subscription)

	type Response struct {
		Subscription *Subscription `json:"subscription"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	data := struct {
		*Options
		CancelAt           interface{} `json:"cancel_at"`
		CancellationReason interface{} `json:"cancellation_reason"`
		CancelAtEnd        interface{} `json:"cancel_at_end"`
	}{
		Options:            opt.Options,
		CancelAt:           s.CancelAt,
		CancellationReason: s.CancellationReason,
		CancelAtEnd:        opt.CancelAtEnd,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Subscription.SetClient(s.client)
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
