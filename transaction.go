package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Transaction represents the Transaction API object
type Transaction struct {
	Identifier

	// Project is the project to which the transaction belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the transaction belongs
	ProjectID string `json:"project_id,omitempty"`
	// Invoice is the invoice used to generate this transaction, if any
	Invoice *Customer `json:"invoice,omitempty"`
	// InvoiceID is the iD of the invoice used to generate this transaction, if any
	InvoiceID *string `json:"invoice_id,omitempty"`
	// Customer is the customer that was linked to this transaction, if any
	Customer *Customer `json:"customer,omitempty"`
	// CustomerID is the iD of the customer that was linked to the transaction, if any
	CustomerID *string `json:"customer_id,omitempty"`
	// Subscription is the subscription to which this transaction belongs
	Subscription *Subscription `json:"subscription,omitempty"`
	// SubscriptionID is the iD of the subscription to which the transaction belongs, if any
	SubscriptionID *string `json:"subscription_id,omitempty"`
	// Token is the token that was used to capture the payment of the transaction, if any
	Token *Token `json:"token,omitempty"`
	// TokenID is the iD of the token was used to capture the payment of the transaction, if any
	TokenID *string `json:"token_id,omitempty"`
	// Card is the card that was used to capture the payment of the transaction, if any
	Card *Card `json:"card,omitempty"`
	// CardID is the iD of the card that was used to capture the payment of the transaction, if any
	CardID *string `json:"card_id,omitempty"`
	// Operations is the operations linked to the transaction
	Operations []*TransactionOperation `json:"operations,omitempty"`
	// Refunds is the list of the transaction refunds
	Refunds []*Refund `json:"refunds,omitempty"`
	// Name is the name of the transaction
	Name string `json:"name,omitempty"`
	// AuthorizedAmount is the amount that was successfully authorized on the transaction
	AuthorizedAmount string `json:"authorized_amount,omitempty"`
	// CapturedAmount is the amount that was successfully captured on the transaction
	CapturedAmount string `json:"captured_amount,omitempty"`
	// Currency is the currency of the transaction
	Currency string `json:"currency,omitempty"`
	// Status is the status of the transaction
	Status string `json:"status,omitempty"`
	// Authorized is the whether the transaction was authorized or not
	Authorized bool `json:"authorized,omitempty"`
	// Captured is the whether the transaction was captured or not
	Captured bool `json:"captured,omitempty"`
	// ProcessoutFee is the processOut fee applied on the transaction
	ProcessoutFee string `json:"processout_fee,omitempty"`
	// EstimatedFee is the gateway fee estimated before processing the payment
	EstimatedFee string `json:"estimated_fee,omitempty"`
	// GatewayFee is the fee taken by the payment gateway to process the payment
	GatewayFee string `json:"gateway_fee,omitempty"`
	// Metadata is the metadata related to the transaction, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata,omitempty"`
	// Sandbox is the define whether or not the transaction is in sandbox environment
	Sandbox bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the transaction was created
	CreatedAt time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the Transaction object and its
// children
func (s *Transaction) SetClient(c *ProcessOut) *Transaction {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Invoice != nil {
		s.Invoice.SetClient(c)
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

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Transaction) Prefill(c *Transaction) *Transaction {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Invoice = c.Invoice
	s.InvoiceID = c.InvoiceID
	s.Customer = c.Customer
	s.CustomerID = c.CustomerID
	s.Subscription = c.Subscription
	s.SubscriptionID = c.SubscriptionID
	s.Token = c.Token
	s.TokenID = c.TokenID
	s.Card = c.Card
	s.CardID = c.CardID
	s.Operations = c.Operations
	s.Refunds = c.Refunds
	s.Name = c.Name
	s.AuthorizedAmount = c.AuthorizedAmount
	s.CapturedAmount = c.CapturedAmount
	s.Currency = c.Currency
	s.Status = c.Status
	s.Authorized = c.Authorized
	s.Captured = c.Captured
	s.ProcessoutFee = c.ProcessoutFee
	s.EstimatedFee = c.EstimatedFee
	s.GatewayFee = c.GatewayFee
	s.Metadata = c.Metadata
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// TransactionFetchRefundsParameters is the structure representing the
// additional parameters used to call Transaction.FetchRefunds
type TransactionFetchRefundsParameters struct {
	*Options
	*Transaction
}

// FetchRefunds allows you to get the transaction's refunds.
func (s Transaction) FetchRefunds(options ...TransactionFetchRefundsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TransactionFetchRefundsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Transaction)

	type Response struct {
		Refunds []*Refund `json:"refunds"`

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

	path := "/transactions/" + url.QueryEscape(s.ID) + "/refunds"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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

	refundsList := []Identifiable{}
	for _, o := range payload.Refunds {
		refundsList = append(refundsList, o.SetClient(s.client))
	}
	refundsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    refundsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"refunds"`
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
		hasMorePrev: true,
	}
	return refundsIterator, nil
}

// TransactionFindRefundParameters is the structure representing the
// additional parameters used to call Transaction.FindRefund
type TransactionFindRefundParameters struct {
	*Options
	*Transaction
}

// FindRefund allows you to find a transaction's refund by its ID.
func (s Transaction) FindRefund(refundID string, options ...TransactionFindRefundParameters) (*Refund, error) {
	if s.client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TransactionFindRefundParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Transaction)

	type Response struct {
		Refund  *Refund `json:"refund"`
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

	path := "/transactions/" + url.QueryEscape(s.ID) + "/refunds/" + url.QueryEscape(refundID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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

	payload.Refund.SetClient(s.client)
	return payload.Refund, nil
}

// TransactionAllParameters is the structure representing the
// additional parameters used to call Transaction.All
type TransactionAllParameters struct {
	*Options
	*Transaction
}

// All allows you to get all the transactions.
func (s Transaction) All(options ...TransactionAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TransactionAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Transaction)

	type Response struct {
		Transactions []*Transaction `json:"transactions"`

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

	path := "/transactions"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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

	transactionsList := []Identifiable{}
	for _, o := range payload.Transactions {
		transactionsList = append(transactionsList, o.SetClient(s.client))
	}
	transactionsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    transactionsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"transactions"`
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
		hasMorePrev: true,
	}
	return transactionsIterator, nil
}

// TransactionFindParameters is the structure representing the
// additional parameters used to call Transaction.Find
type TransactionFindParameters struct {
	*Options
	*Transaction
}

// Find allows you to find a transaction by its ID.
func (s Transaction) Find(transactionID string, options ...TransactionFindParameters) (*Transaction, error) {
	if s.client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TransactionFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Transaction)

	type Response struct {
		Transaction *Transaction `json:"transaction"`
		HasMore     bool         `json:"has_more"`
		Success     bool         `json:"success"`
		Message     string       `json:"message"`
		Code        string       `json:"error_type"`
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

	path := "/transactions/" + url.QueryEscape(transactionID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

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

	payload.Transaction.SetClient(s.client)
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
