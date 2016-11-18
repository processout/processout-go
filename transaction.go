package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Transaction represents the Transaction API object
type Transaction struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the transaction
	ID string `json:"id"`
	// Project is the project to which the transaction belongs
	Project *Project `json:"project"`
	// Customer is the customer that was linked to this transaction
	Customer *Customer `json:"customer"`
	// Subscription is the subscription to which this transaction belongs
	Subscription *Subscription `json:"subscription"`
	// Token is the token that was used to capture the payment of this transaction, if any
	Token *Token `json:"token"`
	// Card is the card that was used to capture the payment of this transaction, if any
	Card *Card `json:"card"`
	// Name is the name of the transaction
	Name string `json:"name"`
	// AuthorizedAmount is the amount that was successfully authorized on the transaction
	AuthorizedAmount string `json:"authorized_amount"`
	// CapturedAmount is the amount that was successfully captured on the transaction
	CapturedAmount string `json:"captured_amount"`
	// Currency is the currency of the transaction
	Currency string `json:"currency"`
	// Status is the status of the transaction
	Status string `json:"status"`
	// Authorized is the whether the transaction was authorized or not
	Authorized bool `json:"authorized"`
	// Captured is the whether the transaction was captured or not
	Captured bool `json:"captured"`
	// ProcessoutFee is the processOut fee applied on the transaction
	ProcessoutFee string `json:"processout_fee"`
	// Metadata is the metadata related to the transaction, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Sandbox is the define whether or not the transaction is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt is the date at which the transaction was created
	CreatedAt *time.Time `json:"created_at"`
}

// SetClient sets the client for the Transaction object and its
// children
func (s *Transaction) SetClient(c *ProcessOut) {
	s.Client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Customer != nil {
		s.Customer.SetClient(c)
	}
	if s.Subscription != nil {
		s.Subscription.SetClient(c)
	}
	if s.Token != nil {
		s.Token.SetClient(c)
	}
	if s.Card != nil {
		s.Card.SetClient(c)
	}
}

// FetchRefunds allows you to get the transaction's refunds.
func (s Transaction) FetchRefunds(options ...Options) ([]*Refund, error) {
	if s.Client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Refunds []*Refund `json:"refunds"`

		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/transactions/" + url.QueryEscape(s.ID) + "/refunds"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
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

	for _, o := range payload.Refunds {
		o.SetClient(s.Client)
	}
	return payload.Refunds, nil
}

// FindRefund allows you to find a transaction's refund by its ID.
func (s Transaction) FindRefund(refundID string, options ...Options) (*Refund, error) {
	if s.Client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Refund  *Refund `json:"refund"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/transactions/" + url.QueryEscape(s.ID) + "/refunds/" + url.QueryEscape(refundID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
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

	payload.Refund.SetClient(s.Client)
	return payload.Refund, nil
}

// All allows you to get all the transactions.
func (s Transaction) All(options ...Options) ([]*Transaction, error) {
	if s.Client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Transactions []*Transaction `json:"transactions"`

		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/transactions"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
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

	for _, o := range payload.Transactions {
		o.SetClient(s.Client)
	}
	return payload.Transactions, nil
}

// Find allows you to find a transaction by its ID.
func (s Transaction) Find(transactionID string, options ...Options) (*Transaction, error) {
	if s.Client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Transaction *Transaction `json:"transaction"`
		Success     bool         `json:"success"`
		Message     string       `json:"message"`
		Code        string       `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/transactions/" + url.QueryEscape(transactionID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

	res, err := http.DefaultClient.Do(req)
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

	payload.Transaction.SetClient(s.Client)
	return payload.Transaction, nil
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
	errors.New(nil, "", "")
}
