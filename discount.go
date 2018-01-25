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

// Discount represents the Discount API object
type Discount struct {
	// ID is the iD of the discount
	ID *string `json:"id,omitempty"`
	// Project is the project to which the discount belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the discount belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Subscription is the subscription to which the discount belongs
	Subscription *Subscription `json:"subscription,omitempty"`
	// SubscriptionID is the iD of the subscription to which the addon belongs
	SubscriptionID *string `json:"subscription_id,omitempty"`
	// Coupon is the coupon used to create the discount, if any
	Coupon *Coupon `json:"coupon,omitempty"`
	// CouponID is the iD of the coupon used to create the discount, if any
	CouponID *string `json:"coupon_id,omitempty"`
	// Name is the name of the discount
	Name *string `json:"name,omitempty"`
	// Amount is the amount discounted
	Amount *string `json:"amount,omitempty"`
	// Percent is the percentage discounted
	Percent *int `json:"percent,omitempty"`
	// ExpiresAt is the date at which the discount will expire
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Metadata is the metadata related to the discount, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Sandbox is the define whether or not the discount is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the discount was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Discount) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Discount object and its
// children
func (s *Discount) SetClient(c *ProcessOut) *Discount {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Subscription != nil {
		s.Subscription.SetClient(c)
	}
	if s.Coupon != nil {
		s.Coupon.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Discount) Prefill(c *Discount) *Discount {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Subscription = c.Subscription
	s.SubscriptionID = c.SubscriptionID
	s.Coupon = c.Coupon
	s.CouponID = c.CouponID
	s.Name = c.Name
	s.Amount = c.Amount
	s.Percent = c.Percent
	s.ExpiresAt = c.ExpiresAt
	s.Metadata = c.Metadata
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// DiscountFetchSubscriptionDiscountsParameters is the structure representing the
// additional parameters used to call Discount.FetchSubscriptionDiscounts
type DiscountFetchSubscriptionDiscountsParameters struct {
	*Options
	*Discount
}

// FetchSubscriptionDiscounts allows you to get the discounts applied to the subscription.
func (s Discount) FetchSubscriptionDiscounts(subscriptionID string, options ...DiscountFetchSubscriptionDiscountsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewDiscount() method to create a new Discount object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := DiscountFetchSubscriptionDiscountsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Discount)

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

	path := "/subscriptions/" + url.QueryEscape(subscriptionID) + "/discounts"

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
		hasMorePrev: false,
	}
	return discountsIterator, nil
}

// DiscountCreateParameters is the structure representing the
// additional parameters used to call Discount.Create
type DiscountCreateParameters struct {
	*Options
	*Discount
}

// Create allows you to create a new discount for the given subscription ID.
func (s Discount) Create(options ...DiscountCreateParameters) (*Discount, error) {
	if s.client == nil {
		panic("Please use the client.NewDiscount() method to create a new Discount object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := DiscountCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Discount)

	type Response struct {
		Discount *Discount `json:"discount"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
	}

	data := struct {
		*Options
		CouponID  interface{} `json:"coupon_id"`
		Name      interface{} `json:"name"`
		Amount    interface{} `json:"amount"`
		ExpiresAt interface{} `json:"expires_at"`
		Metadata  interface{} `json:"metadata"`
	}{
		Options:   opt.Options,
		CouponID:  s.CouponID,
		Name:      s.Name,
		Amount:    s.Amount,
		ExpiresAt: s.ExpiresAt,
		Metadata:  s.Metadata,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(*s.SubscriptionID) + "/discounts"

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

// DiscountFindParameters is the structure representing the
// additional parameters used to call Discount.Find
type DiscountFindParameters struct {
	*Options
	*Discount
}

// Find allows you to find a subscription's discount by its ID.
func (s Discount) Find(subscriptionID, discountID string, options ...DiscountFindParameters) (*Discount, error) {
	if s.client == nil {
		panic("Please use the client.NewDiscount() method to create a new Discount object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := DiscountFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Discount)

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

	path := "/subscriptions/" + url.QueryEscape(subscriptionID) + "/discounts/" + url.QueryEscape(discountID) + ""

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

// DiscountDeleteParameters is the structure representing the
// additional parameters used to call Discount.Delete
type DiscountDeleteParameters struct {
	*Options
	*Discount
}

// Delete allows you to delete a discount applied to a subscription.
func (s Discount) Delete(options ...DiscountDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewDiscount() method to create a new Discount object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := DiscountDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Discount)

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

	path := "/subscriptions/" + url.QueryEscape(*s.SubscriptionID) + "/discounts/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

// dummyDiscount is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyDiscount() {
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
