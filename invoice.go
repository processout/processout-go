package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// Invoices manages the Invoice struct
type Invoices struct {
	p *ProcessOut
}

type Invoice struct {
	// CancelURL : URL where to redirect the customer when the transaction has been canceled. Defaults to ProcessOut's landing page
	CancelURL string `json:"cancel_url"`
	// Currency : Currency of the item's price (ex: USD)
	Currency string `json:"currency"`
	// Custom : Custom value, can be anything. The value is sent back to notify_url
	Custom string `json:"custom"`
	// ID : Id of the created invoice
	ID string `json:"id"`
	// Metas : Contains a key value dictionary representing additional informations shown on the checkout page
	Metas map[string]string `json:"metas"`
	// Name : Name of the item
	Name string `json:"name"`
	// Price : Price of the item
	Price string `json:"price"`
	// RequestEmail : Determine if we want to ask the customer for his email
	RequestEmail bool `json:"request_email"`
	// RequestShipping : Determine if we want to ask the customer for its shipping address
	RequestShipping bool `json:"request_shipping"`
	// ReturnURL : URL where to redirect the customer once the payment has been placed. Defaults to ProcessOut's landing page
	ReturnURL string `json:"return_url"`
	// Shipping : Shipping price added on top of the item price
	Shipping string `json:"shipping"`
	// Taxes : Taxes price added on top of the item price
	Taxes string `json:"taxes"`
	// URL : URL to which you can redirect your customer in order to pay
	URL string `json:"url"`
}

// Customer : Get the customer associated with the current invoice.
func (i Invoices) Customer(invoice *Invoice) (*Customer, error) {
	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	path := "/invoices/{id}/customers"
	path = strings.Replace(path, "{id}", invoice.ID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(i.p.projectID, i.p.projectSecret)

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

// SetCustomer : Link a customer to the invoice.
func (i Invoices) SetCustomer(invoice *Invoice, customerID string) (*Customer, error) {

	type Request struct {
		CustomerID string `json:"customer_id"`
	}

	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	body, err := json.Marshal(&Request{
		CustomerID: customerID,
	})
	if err != nil {
		return nil, err
	}

	path := "/invoices/{id}/customers"
	path = strings.Replace(path, "{id}", invoice.ID, -1)

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", i.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(i.p.projectID, i.p.projectSecret)

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

// Create : Create an invoice.
func (i Invoices) Create(invoice *Invoice) (*Invoice, error) {
	type Response struct {
		Invoice `json:"invoice"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	body, err := json.Marshal(invoice)
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
	req.Header.Set("API-Version", i.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(i.p.projectID, i.p.projectSecret)

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

// Find : Get the invoice data.
func (i Invoices) Find(ID string) (*Invoice, error) {
	type Response struct {
		Invoice `json:"invoice"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/invoices/{id}"
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
	req.SetBasicAuth(i.p.projectID, i.p.projectSecret)

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

// Charge : Charge using a manually generated payment gateway token.
func (i Invoices) Charge(invoice *Invoice, token string) (*CustomerAction, error) {

	type Request struct {
		Token string `json:"token"`
	}

	type Response struct {
		CustomerAction `json:"customer_action"`
		Success        bool   `json:"success"`
		Message        string `json:"message"`
	}

	body, err := json.Marshal(&Request{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	path := "/invoices/{id}/gateways/{gateway_name}/charges"
	path = strings.Replace(path, "{id}", invoice.ID, -1)

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", i.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(i.p.projectID, i.p.projectSecret)

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
	return &payload.CustomerAction, nil
}

// ChargeWithToken : Charge using a customer token.
func (i Invoices) ChargeWithToken(invoice *Invoice, tokenID string) (*CustomerAction, error) {
	type Response struct {
		CustomerAction `json:"customer_action"`
		Success        bool   `json:"success"`
		Message        string `json:"message"`
	}

	path := "/invoices/{id}/tokens/{token_id}/charges"
	path = strings.Replace(path, "{id}", invoice.ID, -1)
	path = strings.Replace(path, "{token_id}", tokenID, -1)

	req, err := http.NewRequest(
		"POST",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", i.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(i.p.projectID, i.p.projectSecret)

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
	return &payload.CustomerAction, nil
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
	}
	errors.New("")
}
