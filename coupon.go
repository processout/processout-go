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

// Coupon represents the Coupon API object
type Coupon struct {
	// ID is the iD of the coupon
	ID *string `json:"id,omitempty"`
	// Project is the project to which the coupon belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the coupon belongs
	ProjectID *string `json:"project_id,omitempty"`
	// AmountOff is the amount to be removed from the subscription price
	AmountOff *string `json:"amount_off,omitempty"`
	// PercentOff is the percent of the subscription amount to be removed (integer between 0 and 100)
	PercentOff *int `json:"percent_off,omitempty"`
	// Currency is the currency of the coupon amount_off
	Currency *string `json:"currency,omitempty"`
	// IterationCount is the number billing cycles the coupon will last when applied to a subscription. If 0, will last forever
	IterationCount *int `json:"iteration_count,omitempty"`
	// MaxRedemptions is the number of time the coupon can be redeemed. If 0, there's no limit
	MaxRedemptions *int `json:"max_redemptions,omitempty"`
	// ExpiresAt is the date at which the coupon will expire
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Metadata is the metadata related to the coupon, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// RedeemedNumber is the number of times the coupon was already redeemed
	RedeemedNumber *int `json:"redeemed_number,omitempty"`
	// Sandbox is the true if the coupon was created in the sandbox environment, false otherwise
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the coupon was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Coupon) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Coupon object and its
// children
func (s *Coupon) SetClient(c *ProcessOut) *Coupon {
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
func (s *Coupon) Prefill(c *Coupon) *Coupon {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.AmountOff = c.AmountOff
	s.PercentOff = c.PercentOff
	s.Currency = c.Currency
	s.IterationCount = c.IterationCount
	s.MaxRedemptions = c.MaxRedemptions
	s.ExpiresAt = c.ExpiresAt
	s.Metadata = c.Metadata
	s.RedeemedNumber = c.RedeemedNumber
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// CouponAllParameters is the structure representing the
// additional parameters used to call Coupon.All
type CouponAllParameters struct {
	*Options
	*Coupon
}

// All allows you to get all the coupons.
func (s Coupon) All(options ...CouponAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewCoupon() method to create a new Coupon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CouponAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Coupon)

	type Response struct {
		Coupons []*Coupon `json:"coupons"`

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

	path := "/coupons"

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

	couponsList := []Identifiable{}
	for _, o := range payload.Coupons {
		couponsList = append(couponsList, o.SetClient(s.client))
	}
	couponsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    couponsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"coupons"`
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
	return couponsIterator, nil
}

// CouponCreateParameters is the structure representing the
// additional parameters used to call Coupon.Create
type CouponCreateParameters struct {
	*Options
	*Coupon
}

// Create allows you to create a new coupon.
func (s Coupon) Create(options ...CouponCreateParameters) (*Coupon, error) {
	if s.client == nil {
		panic("Please use the client.NewCoupon() method to create a new Coupon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CouponCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Coupon)

	type Response struct {
		Coupon  *Coupon `json:"coupon"`
		HasMore bool    `json:"has_more"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
	}

	data := struct {
		*Options
		ID             interface{} `json:"id"`
		AmountOff      interface{} `json:"amount_off"`
		PercentOff     interface{} `json:"percent_off"`
		Currency       interface{} `json:"currency"`
		IterationCount interface{} `json:"iteration_count"`
		MaxRedemptions interface{} `json:"max_redemptions"`
		ExpiresAt      interface{} `json:"expires_at"`
		Metadata       interface{} `json:"metadata"`
	}{
		Options:        opt.Options,
		ID:             s.ID,
		AmountOff:      s.AmountOff,
		PercentOff:     s.PercentOff,
		Currency:       s.Currency,
		IterationCount: s.IterationCount,
		MaxRedemptions: s.MaxRedemptions,
		ExpiresAt:      s.ExpiresAt,
		Metadata:       s.Metadata,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/coupons"

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

	payload.Coupon.SetClient(s.client)
	return payload.Coupon, nil
}

// CouponFindParameters is the structure representing the
// additional parameters used to call Coupon.Find
type CouponFindParameters struct {
	*Options
	*Coupon
}

// Find allows you to find a coupon by its ID.
func (s Coupon) Find(couponID string, options ...CouponFindParameters) (*Coupon, error) {
	if s.client == nil {
		panic("Please use the client.NewCoupon() method to create a new Coupon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CouponFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Coupon)

	type Response struct {
		Coupon  *Coupon `json:"coupon"`
		HasMore bool    `json:"has_more"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
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

	path := "/coupons/" + url.QueryEscape(couponID) + ""

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

	payload.Coupon.SetClient(s.client)
	return payload.Coupon, nil
}

// CouponSaveParameters is the structure representing the
// additional parameters used to call Coupon.Save
type CouponSaveParameters struct {
	*Options
	*Coupon
}

// Save allows you to save the updated coupon attributes.
func (s Coupon) Save(options ...CouponSaveParameters) (*Coupon, error) {
	if s.client == nil {
		panic("Please use the client.NewCoupon() method to create a new Coupon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CouponSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Coupon)

	type Response struct {
		Coupon  *Coupon `json:"coupon"`
		HasMore bool    `json:"has_more"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
	}

	data := struct {
		*Options
		Metadata interface{} `json:"metadata"`
	}{
		Options:  opt.Options,
		Metadata: s.Metadata,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/coupons/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	payload.Coupon.SetClient(s.client)
	return payload.Coupon, nil
}

// CouponDeleteParameters is the structure representing the
// additional parameters used to call Coupon.Delete
type CouponDeleteParameters struct {
	*Options
	*Coupon
}

// Delete allows you to delete the coupon.
func (s Coupon) Delete(options ...CouponDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewCoupon() method to create a new Coupon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CouponDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Coupon)

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

	path := "/coupons/" + url.QueryEscape(*s.ID) + ""

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

// dummyCoupon is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCoupon() {
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
