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

// RecurringInvoices manages the RecurringInvoice struct
type RecurringInvoices struct {
	p *ProcessOut
}

type RecurringInvoice struct {
	// ID : ID of the recurring invoice
	ID string `json:"id"`
	// Customer : Customer linked to the recurring invoice
	Customer *Customer `json:"customer"`
	// URL : URL to which you may redirect your customer to authorize the recurring invoice
	URL string `json:"url"`
	// Name : Name of the recurring invoice
	Name string `json:"name"`
	// Price : Price of the recurring invoice
	Price string `json:"price"`
	// Currency : Currency of the recurring invoice
	Currency string `json:"currency"`
	// Taxes : Taxes applied on the recurring invoice (on top of the price)
	Taxes string `json:"taxes"`
	// Shipping : Shipping fees applied on the recurring invoice (on top of the price)
	Shipping string `json:"shipping"`
	// ReturnURL : URL where the customer will be redirected upon payment
	ReturnURL string `json:"return_url"`
	// CancelURL : URL where the customer will be redirected if the paymen was canceled
	CancelURL string `json:"cancel_url"`
	// Interval : The recurring payment period, formatted in the format "1d2w3m4y" (day, week, month, year)
	Interval string `json:"interval"`
	// TrialPeriod : The trial period. The customer will not be charged during this time span. Formatted in the format "1d2w3m4y" (day, week, month, year)
	TrialPeriod string `json:"trial_period"`
	// Custom : Custom variable passed along in the events/webhooks
	Custom string `json:"custom"`
	// Ended : Weither or not the recurring invoice has ended (programmatically or canceled)
	Ended bool `json:"ended"`
	// EndedReason : Reason as to why the recurring invoice ended
	EndedReason string `json:"ended_reason"`
	// Sandbox : Define whether or not the authorization is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the invoice was created
	CreatedAt time.Time `json:"created_at"`
}


// Customer : Get the customer linked to the recurring invoice.
func (s RecurringInvoices) Customer(recurringInvoice *RecurringInvoice) (*Customer, error) {

	type Response struct {
		Customer `json:"customer"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/recurring-invoices/"+url.QueryEscape(recurringInvoice.ID)+"/customers"

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

// Invoice : Get the invoice corresponding to the last iteration of the recurring invoice.
func (s RecurringInvoices) Invoice(recurringInvoice *RecurringInvoice) (*Invoice, error) {

	type Response struct {
		Invoice `json:"invoice"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/recurring-invoices/"+url.QueryEscape(recurringInvoice.ID)+"/invoices"

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

// Create : Create a new recurring invoice for the given customer.
func (s RecurringInvoices) Create(recurringInvoice *RecurringInvoice, customerID string) (*RecurringInvoice, error) {

	type Response struct {
		RecurringInvoice `json:"recurring_invoice"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 body , err := json.Marshal(map[string]interface{}{
		"name": recurringInvoice.Name,
		"price": recurringInvoice.Price,
		"taxes": recurringInvoice.Taxes,
		"shipping": recurringInvoice.Shipping,
		"currency": recurringInvoice.Currency,
		"return_url": recurringInvoice.ReturnURL,
		"cancel_url": recurringInvoice.CancelURL,
		"custom": recurringInvoice.Custom,
		"interval": recurringInvoice.Interval,
		"trial_period": recurringInvoice.TrialPeriod,
		"ended_reason": recurringInvoice.EndedReason,
		"customer_id": customerID,

	})
	if err != nil {
		return nil, err
	}

	path := "/recurring-invoices"

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
	return &payload.RecurringInvoice, nil
}

// Find : Find a recurring invoice by its ID.
func (s RecurringInvoices) Find(recurringInvoiceID string) (*RecurringInvoice, error) {

	type Response struct {
		RecurringInvoice `json:"recurring_invoice"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/recurring-invoices/"+url.QueryEscape(recurringInvoiceID)+""

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
	return &payload.RecurringInvoice, nil
}

// End : End a recurring invoice. The reason may be provided as well.
func (s RecurringInvoices) End(recurringInvoice *RecurringInvoice, reason string) error {

	type Response struct {Success bool `json:"success"`
		Message string `json:"message"`
	}

	 body , err := json.Marshal(map[string]interface{}{
		"reason": reason,

	})
	if err != nil {
		return err
	}

	path := "/recurring-invoices/"+url.QueryEscape(recurringInvoice.ID)+""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
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


// dummyRecurringInvoice is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyRecurringInvoice() {
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
