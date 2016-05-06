package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// RecurringInvoices manages the RecurringInvoice struct
type RecurringInvoices struct {
	p *ProcessOut
}

type RecurringInvoice struct {
	// Currency : Currency of the item's price (ex: USD)
	Currency string `json:"currency"`
	// Ended : Whether or not the recurring invoice has ended. True if expired is true
	Ended bool `json:"ended"`
	// EndedReason : The reason why the recurring invoice ended
	EndedReason string `json:"ended_reason"`
	// ID : Id of the recurring invoice
	ID string `json:"id"`
	// Name : Name of the item
	Name string `json:"name"`
	// Price : Price of the item
	Price string `json:"price"`
	// RecurringDays : The recurring payment period, in days. ProcessOut will make sure to collect your payments at the end of each period.
	RecurringDays int `json:"recurring_days"`
	// Shipping : Shipping price added on top of the item price
	Shipping string `json:"shipping"`
	// Taxes : Taxes price added on top of the item price
	Taxes string `json:"taxes"`
	// TrialDays : The recurring trial period, in days.
	TrialDays int `json:"trial_days"`
	// URL : URL to which you can redirect your customer in order to activate the recurring invoice
	URL string `json:"url"`
}

// Customer : Get the customer linked to the recurring invoice.
func (r RecurringInvoices) Customer(recurringInvoice *RecurringInvoice) (*Customer, error) {
	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	path := "/recurring-invoices/{id}/customers"
	path = strings.Replace(path, "{id}", recurringInvoice.ID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(r.p.projectID, r.p.projectSecret)

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

// Create : Create a new recurring invoice.
func (r RecurringInvoices) Create(recurringInvoice *RecurringInvoice, customerID string) (*RecurringInvoice, error) {
	type Response struct {
		RecurringInvoice `json:"recurring_invoice"`
		Success          bool   `json:"success"`
		Message          string `json:"message"`
	}

	body, err := json.Marshal(recurringInvoice)
	if err != nil {
		return nil, err
	}

	path := "/customers/{customer_id}/recurring-invoices"
	path = strings.Replace(path, "{customer_id}", customerID, -1)

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", r.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(r.p.projectID, r.p.projectSecret)

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

// Find : Get the recurring invoice data.
func (r RecurringInvoices) Find(ID string) (*RecurringInvoice, error) {
	type Response struct {
		RecurringInvoice `json:"recurring_invoice"`
		Success          bool   `json:"success"`
		Message          string `json:"message"`
	}

	path := "/recurring-invoices/{id}"
	path = strings.Replace(path, "{id}", ID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(r.p.projectID, r.p.projectSecret)

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

// End : End a recurring invoice.
func (r RecurringInvoices) End(recurringInvoice *RecurringInvoice, reason string) error {

	type Request struct {
		Reason string `json:"reason"`
	}

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	body, err := json.Marshal(&Request{
		Reason: reason,
	})
	if err != nil {
		return err
	}

	path := "/recurring-invoices/{id}"
	path = strings.Replace(path, "{id}", recurringInvoice.ID, -1)

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", r.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(r.p.projectID, r.p.projectSecret)

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

// Invoice : Get the invoice representing the new recurring invoice iteration.
func (r RecurringInvoices) Invoice(recurringInvoice *RecurringInvoice) (*Invoice, error) {
	type Response struct {
		Invoice `json:"invoice"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/recurring-invoices/{id}/invoices"
	path = strings.Replace(path, "{id}", recurringInvoice.ID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(r.p.projectID, r.p.projectSecret)

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
	}
	errors.New("")
}
