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

// Customers manages the Customer struct
type Customers struct {
	p *ProcessOut
}

type Customer struct {
	// ID : ID of the customer
	ID string `json:"id"`
	// Project : Project to which the customer belongs
	Project *Project `json:"project"`
	// Email : Email of the customer
	Email string `json:"email"`
	// FirstName : First name of the customer
	FirstName string `json:"first_name"`
	// LastName : Last name of the customer
	LastName string `json:"last_name"`
	// Address1 : Address of the customer
	Address1 string `json:"address1"`
	// Address2 : Secondary address of the customer
	Address2 string `json:"address2"`
	// City : City of the customer
	City string `json:"city"`
	// State : State of the customer
	State string `json:"state"`
	// Zip : ZIP code of the customer
	Zip string `json:"zip"`
	// CountryCode : Country code of the customer
	CountryCode string `json:"country_code"`
	// Metadata : Metadata related to the customer, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// HasPin : Wether the customer has a PIN set or not
	HasPin bool `json:"has_pin"`
	// Sandbox : Define whether or not the customer is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the customer was created
	CreatedAt time.Time `json:"created_at"`
}

// RecurringInvoices : Get the recurring invoices linked to the customer.
func (s Customers) RecurringInvoices(customer *Customer, options ...Options) ([]*RecurringInvoice, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		RecurringInvoices []*RecurringInvoice `json:"recurring_invoices"`
		Success           bool                `json:"success"`
		Message           string              `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers/" + url.QueryEscape(customer.ID) + "/recurring-invoices"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return payload.RecurringInvoices, nil
}

// Tokens : Get the customer's tokens.
func (s Customers) Tokens(customer *Customer, options ...Options) ([]*Token, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Tokens  []*Token `json:"tokens"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers/" + url.QueryEscape(customer.ID) + "/tokens"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return payload.Tokens, nil
}

// Token : Get a specific customer's token by its ID.
func (s Customers) Token(customer *Customer, tokenID string, options ...Options) (*Token, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Token   `json:"token"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers/" + url.QueryEscape(customer.ID) + "/tokens/" + url.QueryEscape(tokenID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Token, nil
}

// All : Get all the customers.
func (s Customers) All(options ...Options) ([]*Customer, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Customers []*Customer `json:"customers"`
		Success   bool        `json:"success"`
		Message   string      `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return payload.Customers, nil
}

// Create : Create a new customer.
func (s Customers) Create(customer *Customer, options ...Options) (*Customer, error) {
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
	}

	body, err := json.Marshal(map[string]interface{}{
		"email":        customer.Email,
		"first_name":   customer.FirstName,
		"last_name":    customer.LastName,
		"address1":     customer.Address1,
		"address2":     customer.Address2,
		"city":         customer.City,
		"state":        customer.State,
		"zip":          customer.Zip,
		"country_code": customer.CountryCode,
		"metadata":     customer.Metadata,
		"expand":       opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Customer, nil
}

// Find : Find a customer by its ID.
func (s Customers) Find(customerID string, options ...Options) (*Customer, error) {
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
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers/" + url.QueryEscape(customerID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Customer, nil
}

// Save : Save the updated customer attributes.
func (s Customers) Save(customer *Customer, options ...Options) (*Customer, error) {
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
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers/" + url.QueryEscape(customer.ID) + ""

	req, err := http.NewRequest(
		"PUT",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Customer, nil
}

// Delete : Delete the customer.
func (s Customers) Delete(customer *Customer, options ...Options) error {
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
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return err
	}

	path := "/customers/" + url.QueryEscape(customer.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
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
		return err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return err
	}

	if !payload.Success {
		return errors.New(payload.Message)
	}
	return nil
}

// dummyCustomer is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCustomer() {
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
