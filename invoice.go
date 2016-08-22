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

// Invoices manages the Invoice struct
type Invoices struct {
	p *ProcessOut
}

type Invoice struct {
	// ID : ID of the invoice
	ID string `json:"id"`
	// Customer : Customer linked to the invoice, if any
	Customer *Customer `json:"customer"`
	// RecurringInvoice : Recurring invoice to which the invoice is linked to, if any
	RecurringInvoice *RecurringInvoice `json:"recurring_invoice"`
	// URL : URL to which you may redirect your customer to proceed with the payment
	URL string `json:"url"`
	// Name : Name of the invoice
	Name string `json:"name"`
	// Price : Price of the invoice
	Price string `json:"price"`
	// Currency : Currency of the invoice
	Currency string `json:"currency"`
	// Taxes : Taxes applied on the invoice (on top of the price)
	Taxes string `json:"taxes"`
	// Shipping : Shipping fees applied on the invoice (on top of the price)
	Shipping string `json:"shipping"`
	// RequestEmail : Choose whether or not to request the email during the checkout process
	RequestEmail bool `json:"request_email"`
	// RequestShipping : Choose whether or not to request the shipping address during the checkout process
	RequestShipping bool `json:"request_shipping"`
	// ReturnURL : URL where the customer will be redirected upon payment
	ReturnURL string `json:"return_url"`
	// CancelURL : URL where the customer will be redirected if the paymen was canceled
	CancelURL string `json:"cancel_url"`
	// Custom : Custom variable passed along in the events/webhooks
	Custom string `json:"custom"`
	// Sandbox : Define whether or not the authorization is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the invoice was created
	CreatedAt time.Time `json:"created_at"`
}

// Customer : Get the customer linked to the invoice.
func (s Invoices) Customer(invoice *Invoice) (*Customer, error) {

	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/invoices/" + url.QueryEscape(invoice.ID) + "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
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

// AssignCustomer : Assign a customer to the invoice.
func (s Invoices) AssignCustomer(invoice *Invoice, customerID string) (*Customer, error) {

	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"customer_id": customerID,
	})
	if err != nil {
		return nil, err
	}

	path := "/invoices/" + url.QueryEscape(invoice.ID) + "/customers"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
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

// Charge : Charge the invoice using the given customer token ID.
func (s Invoices) Charge(invoice *Invoice, tokenID string) error {

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return err
	}

	path := "/invoices/" + url.QueryEscape(invoice.ID) + "/tokens/" + url.QueryEscape(tokenID) + "/charges"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
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

// Tokens : Get all the customer tokens available on the current invoice.
func (s Invoices) Tokens(invoice *Invoice) ([]*Token, error) {

	type Response struct {
		Tokens  []*Token `json:"tokens"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/invoices/" + url.QueryEscape(invoice.ID) + "/tokens"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
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

// All : Get all the invoices.
func (s Invoices) All() ([]*Invoice, error) {

	type Response struct {
		Invoices []*Invoice `json:"invoices"`
		Success  bool       `json:"success"`
		Message  string     `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/invoices"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
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
	return payload.Invoices, nil
}

// Create : Create a new invoice.
func (s Invoices) Create(invoice *Invoice) (*Invoice, error) {

	type Response struct {
		Invoice `json:"invoice"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":             invoice.Name,
		"price":            invoice.Price,
		"taxes":            invoice.Taxes,
		"shipping":         invoice.Shipping,
		"currency":         invoice.Currency,
		"request_email":    invoice.RequestEmail,
		"request_shipping": invoice.RequestShipping,
		"return_url":       invoice.ReturnURL,
		"cancel_url":       invoice.CancelURL,
		"custom":           invoice.Custom,
	})
	if err != nil {
		return nil, err
	}

	path := "/invoices"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
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
	return &payload.Invoice, nil
}

// Find : Find an invoice by its ID.
func (s Invoices) Find(invoiceID string) (*Invoice, error) {

	type Response struct {
		Invoice `json:"invoice"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/invoices/" + url.QueryEscape(invoiceID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
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
	return &payload.Invoice, nil
}

// Lock : Lock the invoice so it can't be interacted with anymore.
func (s Invoices) Lock(invoice *Invoice) error {

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return err
	}

	path := "/invoices/" + url.QueryEscape(invoice.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
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

// dummyInvoice is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyInvoice() {
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