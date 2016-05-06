package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// TailoredInvoices manages the TailoredInvoice struct
type TailoredInvoices struct {
	p *ProcessOut
}

type TailoredInvoice struct {
	// CancelURL : URL where to redirect the customer when the transaction has been canceled. Defaults to ProcessOut's landing page
	CancelURL string `json:"cancel_url"`
	// Currency : Currency of the item's price (ex: USD)
	Currency string `json:"currency"`
	// ID : Id of the tailored invoice
	ID string `json:"id"`
	// Name : Name of the item
	Name string `json:"name"`
	// Price : Price of the item
	Price string `json:"price"`
	// ReturnURL : URL where to redirect the customer once the payment has been placed. Defaults to ProcessOut's landing page
	ReturnURL string `json:"return_url"`
	// Shipping : Shipping price added on top of the item price
	Shipping string `json:"shipping"`
	// Taxes : Taxes price added on top of the item price
	Taxes string `json:"taxes"`
}

// All : List all tailored invoices.
func (t TailoredInvoices) All() ([]*TailoredInvoice, error) {
	type Response struct {
		TailoredInvoices []*TailoredInvoice `json:"tailored_invoices"`
		Success          bool               `json:"success"`
		Message          string             `json:"message"`
	}

	path := "/tailored-invoices"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(t.p.projectID, t.p.projectSecret)

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
	return payload.TailoredInvoices, nil
}

// Create : Create a new tailored invoice.
func (t TailoredInvoices) Create(tailoredInvoice *TailoredInvoice) (*TailoredInvoice, error) {
	type Response struct {
		TailoredInvoice `json:"tailored_invoice"`
		Success         bool   `json:"success"`
		Message         string `json:"message"`
	}

	body, err := json.Marshal(tailoredInvoice)
	if err != nil {
		return nil, err
	}

	path := "/tailored-invoices"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", t.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(t.p.projectID, t.p.projectSecret)

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
	return &payload.TailoredInvoice, nil
}

// Invoice : Create an invoice from a tailored invoice.
func (t TailoredInvoices) Invoice(tailoredInvoice *TailoredInvoice) (*Invoice, error) {
	type Response struct {
		Invoice `json:"invoice"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/tailored-invoices/{id}/invoices"
	path = strings.Replace(path, "{id}", tailoredInvoice.ID, -1)

	req, err := http.NewRequest(
		"POST",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", t.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(t.p.projectID, t.p.projectSecret)

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

// Find : Get tailored invoice data.
func (t TailoredInvoices) Find(ID string) (*TailoredInvoice, error) {
	type Response struct {
		TailoredInvoice `json:"tailored_invoice"`
		Success         bool   `json:"success"`
		Message         string `json:"message"`
	}

	path := "/tailored-invoices/{id}"
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
	req.SetBasicAuth(t.p.projectID, t.p.projectSecret)

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
	return &payload.TailoredInvoice, nil
}

// Save : Update the tailored invoice data.
func (t TailoredInvoices) Save(tailoredInvoice *TailoredInvoice) (*TailoredInvoice, error) {
	type Response struct {
		TailoredInvoice `json:"tailored_invoice"`
		Success         bool   `json:"success"`
		Message         string `json:"message"`
	}

	body, err := json.Marshal(tailoredInvoice)
	if err != nil {
		return nil, err
	}

	path := "/tailored-invoices/{id}"
	path = strings.Replace(path, "{id}", tailoredInvoice.ID, -1)

	req, err := http.NewRequest(
		"PUT",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", t.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(t.p.projectID, t.p.projectSecret)

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
	return &payload.TailoredInvoice, nil
}

// Delete : Delete a tailored invoice.
func (t TailoredInvoices) Delete(tailoredInvoice *TailoredInvoice) error {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/tailored-invoices/{id}"
	path = strings.Replace(path, "{id}", tailoredInvoice.ID, -1)

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", t.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(t.p.projectID, t.p.projectSecret)

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

// dummyTailoredInvoice is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyTailoredInvoice() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
	}
	errors.New("")
}
