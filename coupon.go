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

// Coupons manages the Coupon struct
type Coupons struct {
	p *ProcessOut
}

type Coupon struct {
	// ID : ID of the coupon
	ID string `json:"id"`
	// Project : Project to which the coupon belongs
	Project *Project `json:"project"`
	// Name : Name of the coupon
	Name string `json:"name"`
	// AmountOff : Amount to be removed from the subscription price
	AmountOff string `json:"amount_off"`
	// PercentOff : Percent of the subscription amount to be removed (integer between 0 and 100)
	PercentOff int `json:"percent_off"`
	// Currency : Currency of the coupon amount_off
	Currency string `json:"currency"`
	// MaxRedemptions : Number of time the coupon can be redeemed. If 0, there's no limit
	MaxRedemptions int `json:"max_redemptions"`
	// ExpiresAt : Date at which the coupon will expire
	ExpiresAt time.Time `json:"expires_at"`
	// Metadata : Metadata related to the coupon, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// IterationCount : Number billing cycles the coupon will last when applied to a subscription. If 0, will last forever
	IterationCount int `json:"iteration_count"`
	// RedeemedNumber : Number of time the coupon was redeemed
	RedeemedNumber int `json:"redeemed_number"`
	// Sandbox : Define whether or not the plan is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the plan was created
	CreatedAt time.Time `json:"created_at"`
}

// All : Get all the coupons.
func (s Coupons) All(options ...Options) ([]*Coupon, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Coupons []*Coupon `json:"coupons"`

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

	path := "/coupons"

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

	return payload.Coupons, nil
}

// Create : Create a new coupon.
func (s Coupons) Create(coupon *Coupon, options ...Options) (*Coupon, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Coupon  *Coupon `json:"coupon"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"id":              coupon.ID,
		"amount_off":      coupon.AmountOff,
		"percent_off":     coupon.PercentOff,
		"currency":        coupon.Currency,
		"iteration_count": coupon.IterationCount,
		"max_redemptions": coupon.MaxRedemptions,
		"expires_at":      coupon.ExpiresAt,
		"metadata":        coupon.Metadata,
		"expand":          opt.Expand,
		"filter":          opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/coupons"

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

	return payload.Coupon, nil
}

// Find : Find a coupon by its ID.
func (s Coupons) Find(couponID string, options ...Options) (*Coupon, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Coupon  *Coupon `json:"coupon"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/coupons/" + url.QueryEscape(couponID) + ""

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

	return payload.Coupon, nil
}

// Save : Save the updated coupon attributes.
func (s Coupons) Save(coupon *Coupon, options ...Options) (*Coupon, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Coupon  *Coupon `json:"coupon"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"metadata": coupon.Metadata,
		"expand":   opt.Expand,
		"filter":   opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/coupons/" + url.QueryEscape(coupon.ID) + ""

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

	return payload.Coupon, nil
}

// Delete : Delete the coupon.
func (s Coupons) Delete(coupon *Coupon, options ...Options) *Error {
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

	path := "/coupons/" + url.QueryEscape(coupon.ID) + ""

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
	}
	errors.New("")
}
