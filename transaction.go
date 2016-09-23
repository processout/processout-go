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

// Transactions manages the Transaction struct
type Transactions struct {
	p *ProcessOut
}

type Transaction struct {
	// ID : ID of the transaction
	ID string `json:"id"`
	// Name : Name of the transaction
	Name string `json:"name"`
	// Amount : Amount of the transaction
	Amount string `json:"amount"`
	// Currency : Currency of the transaction
	Currency string `json:"currency"`
	// Status : Status of the transaction
	Status string `json:"status"`
	// ProcessoutFee : ProcessOut fee applied on the transaction
	ProcessoutFee string `json:"processout_fee"`
	// Metadata : Metadata related to the transaction, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Sandbox : Define whether or not the transaction is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the transaction was created
	CreatedAt time.Time `json:"created_at"`
}

// Refunds : Get the transaction's refunds.
func (s Transactions) Refunds(transaction *Transaction, options ...Options) ([]*Refund, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Refunds []*Refund `json:"refunds"`
		Success bool      `json:"success"`
		Message string    `json:"message"`
		Code    string    `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/transactions/" + url.QueryEscape(transaction.ID) + "/refunds"

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
	return payload.Refunds, nil
}

// All : Get all the transactions.
func (s Transactions) All(options ...Options) ([]*Transaction, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Transactions []*Transaction `json:"transactions"`
		Success      bool           `json:"success"`
		Message      string         `json:"message"`
		Code         string         `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/transactions"

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
	return payload.Transactions, nil
}

// Find : Find a transaction by its ID.
func (s Transactions) Find(transactionID string, options ...Options) (*Transaction, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Transaction `json:"transaction"`
		Success     bool   `json:"success"`
		Message     string `json:"message"`
		Code        string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
		"filter": opt.Filter,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/transactions/" + url.QueryEscape(transactionID) + ""

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
	return &payload.Transaction, nil
}

// dummyTransaction is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyTransaction() {
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
