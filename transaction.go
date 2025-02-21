package processout

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
)

// Transaction represents the Transaction API object
type Transaction struct {
	// ID is the iD of the transaction
	ID *string `json:"id,omitempty"`
	// Project is the project to which the transaction belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the transaction belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Invoice is the invoice used to generate this transaction, if any
	Invoice *Invoice `json:"invoice,omitempty"`
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
	// GatewayConfiguration is the gateway Configuration is the last gateway configuration that was used to process the payment, if any
	GatewayConfiguration *GatewayConfiguration `json:"gateway_configuration,omitempty"`
	// ExternalThreeDSGatewayConfiguration is the external ThreeDS Gateway Configuration is the gateway configuration that was used to authenticate the payment, if configured
	ExternalThreeDSGatewayConfiguration *GatewayConfiguration `json:"external_three_d_s_gateway_configuration,omitempty"`
	// GatewayConfigurationID is the iD of the last gateway configuration that was used to process the payment, if any
	GatewayConfigurationID *string `json:"gateway_configuration_id,omitempty"`
	// Operations is the operations linked to the transaction
	Operations *[]*TransactionOperation `json:"operations,omitempty"`
	// Refunds is the list of the transaction refunds
	Refunds *[]*Refund `json:"refunds,omitempty"`
	// Name is the name of the transaction
	Name *string `json:"name,omitempty"`
	// Amount is the amount requested when creating the transaction
	Amount *string `json:"amount,omitempty"`
	// AmountLocal is the amount requested when creating the transaction, in the currency of the project
	AmountLocal *string `json:"amount_local,omitempty"`
	// AuthorizedAmount is the amount that was successfully authorized on the transaction
	AuthorizedAmount *string `json:"authorized_amount,omitempty"`
	// AuthorizedAmountLocal is the amount that was successfully authorized on the transaction, in the currency of the project
	AuthorizedAmountLocal *string `json:"authorized_amount_local,omitempty"`
	// CapturedAmount is the amount that was successfully captured on the transaction
	CapturedAmount *string `json:"captured_amount,omitempty"`
	// CapturedAmountLocal is the amount that was successfully captured on the transaction, in the currency of the project
	CapturedAmountLocal *string `json:"captured_amount_local,omitempty"`
	// RefundedAmount is the amount that was successfully refunded on the transaction
	RefundedAmount *string `json:"refunded_amount,omitempty"`
	// RefundedAmountLocal is the amount that was successfully refunded on the transaction, in the currency of the project
	RefundedAmountLocal *string `json:"refunded_amount_local,omitempty"`
	// AvailableAmount is the amount available on the transaction (captured - refunded)
	AvailableAmount *string `json:"available_amount,omitempty"`
	// AvailableAmountLocal is the amount available on the transaction (captured - refunded), in the currency of the project
	AvailableAmountLocal *string `json:"available_amount_local,omitempty"`
	// VoidedAmount is the amount that was voided on the transaction
	VoidedAmount *string `json:"voided_amount,omitempty"`
	// VoidedAmountLocal is the amount that was voided on the transaction, in the currency of the project
	VoidedAmountLocal *string `json:"voided_amount_local,omitempty"`
	// Currency is the currency of the transaction
	Currency *string `json:"currency,omitempty"`
	// ErrorCode is the error code of the transaction, when the payment has failed
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorMessage is the error message of the transaction, when the payment has failed
	ErrorMessage *string `json:"error_message,omitempty"`
	// AcquirerName is the name of the merchant acquirer
	AcquirerName *string `json:"acquirer_name,omitempty"`
	// GatewayName is the name of the last gateway the transaction was attempted on (successfully or not). Use the operations list to get the full transaction's history
	GatewayName *string `json:"gateway_name,omitempty"`
	// ThreeDSStatus is the status of the potential 3-D Secure authentication
	ThreeDSStatus *string `json:"three_d_s_status,omitempty"`
	// Status is the status of the transaction
	Status *string `json:"status,omitempty"`
	// Authorized is the whether the transaction was authorized or not
	Authorized *bool `json:"authorized,omitempty"`
	// Captured is the whether the transaction was captured or not
	Captured *bool `json:"captured,omitempty"`
	// Voided is the whether the transaction was voided or not
	Voided *bool `json:"voided,omitempty"`
	// Refunded is the whether the transaction was refunded or not
	Refunded *bool `json:"refunded,omitempty"`
	// Chargedback is the whether the transaction was charged back or not
	Chargedback *bool `json:"chargedback,omitempty"`
	// ReceivedFraudNotification is the whether the transaction received a fraud notification event or not
	ReceivedFraudNotification *bool `json:"received_fraud_notification,omitempty"`
	// ReceivedRetrievalRequest is the whether the transaction received a retrieval request event or not
	ReceivedRetrievalRequest *bool `json:"received_retrieval_request,omitempty"`
	// ProcessoutFee is the processOut fee applied on the transaction
	ProcessoutFee *string `json:"processout_fee,omitempty"`
	// EstimatedFee is the gateway fee estimated before processing the payment
	EstimatedFee *string `json:"estimated_fee,omitempty"`
	// GatewayFee is the fee taken by the payment gateway to process the payment
	GatewayFee *string `json:"gateway_fee,omitempty"`
	// GatewayFeeLocal is the fee taken by the payment gateway to process the payment, in the currency of the project
	GatewayFeeLocal *string `json:"gateway_fee_local,omitempty"`
	// CurrencyFee is the currency of the fee taken on the transaction (field `gateway_fee`)
	CurrencyFee *string `json:"currency_fee,omitempty"`
	// Metadata is the metadata related to the transaction, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Sandbox is the define whether or not the transaction is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the transaction was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ChargedbackAt is the date at which the transaction was charged back
	ChargedbackAt *time.Time `json:"chargedback_at,omitempty"`
	// RefundedAt is the date at which the transaction was refunded
	RefundedAt *time.Time `json:"refunded_at,omitempty"`
	// AuthorizedAt is the date at which the transaction was authorized
	AuthorizedAt *time.Time `json:"authorized_at,omitempty"`
	// CapturedAt is the date at which the transaction was captured
	CapturedAt *time.Time `json:"captured_at,omitempty"`
	// VoidedAt is the date at which the transaction was voided
	VoidedAt *time.Time `json:"voided_at,omitempty"`
	// ThreeDS is the 3DS data of a transaction if it was authenticated
	ThreeDS *ThreeDS `json:"three_d_s,omitempty"`
	// CvcCheck is the cVC check done during the transaction
	CvcCheck *string `json:"cvc_check,omitempty"`
	// AvsCheck is the aVS check done during the transaction
	AvsCheck *string `json:"avs_check,omitempty"`
	// InitialSchemeTransactionID is the initial scheme ID that was referenced in the request
	InitialSchemeTransactionID *string `json:"initial_scheme_transaction_id,omitempty"`
	// SchemeID is the the ID assigned to the transaction by the scheme in the last successful authorization
	SchemeID *string `json:"scheme_id,omitempty"`
	// PaymentType is the payment type of the transaction
	PaymentType *string `json:"payment_type,omitempty"`
	// Eci is the the Electronic Commerce Indicator
	Eci *string `json:"eci,omitempty"`
	// NativeApm is the native APM response data
	NativeApm *NativeAPMResponse `json:"native_apm,omitempty"`
	// ExternalDetails is the additional data about the transaction, originating from a PSP, for example customer shipping address
	ExternalDetails interface{} `json:"external_details,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Transaction) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
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
	if s.GatewayConfiguration != nil {
		s.GatewayConfiguration.SetClient(c)
	}
	if s.ExternalThreeDSGatewayConfiguration != nil {
		s.ExternalThreeDSGatewayConfiguration.SetClient(c)
	}
	if s.ThreeDS != nil {
		s.ThreeDS.SetClient(c)
	}
	if s.NativeApm != nil {
		s.NativeApm.SetClient(c)
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
	s.GatewayConfiguration = c.GatewayConfiguration
	s.ExternalThreeDSGatewayConfiguration = c.ExternalThreeDSGatewayConfiguration
	s.GatewayConfigurationID = c.GatewayConfigurationID
	s.Operations = c.Operations
	s.Refunds = c.Refunds
	s.Name = c.Name
	s.Amount = c.Amount
	s.AmountLocal = c.AmountLocal
	s.AuthorizedAmount = c.AuthorizedAmount
	s.AuthorizedAmountLocal = c.AuthorizedAmountLocal
	s.CapturedAmount = c.CapturedAmount
	s.CapturedAmountLocal = c.CapturedAmountLocal
	s.RefundedAmount = c.RefundedAmount
	s.RefundedAmountLocal = c.RefundedAmountLocal
	s.AvailableAmount = c.AvailableAmount
	s.AvailableAmountLocal = c.AvailableAmountLocal
	s.VoidedAmount = c.VoidedAmount
	s.VoidedAmountLocal = c.VoidedAmountLocal
	s.Currency = c.Currency
	s.ErrorCode = c.ErrorCode
	s.ErrorMessage = c.ErrorMessage
	s.AcquirerName = c.AcquirerName
	s.GatewayName = c.GatewayName
	s.ThreeDSStatus = c.ThreeDSStatus
	s.Status = c.Status
	s.Authorized = c.Authorized
	s.Captured = c.Captured
	s.Voided = c.Voided
	s.Refunded = c.Refunded
	s.Chargedback = c.Chargedback
	s.ReceivedFraudNotification = c.ReceivedFraudNotification
	s.ReceivedRetrievalRequest = c.ReceivedRetrievalRequest
	s.ProcessoutFee = c.ProcessoutFee
	s.EstimatedFee = c.EstimatedFee
	s.GatewayFee = c.GatewayFee
	s.GatewayFeeLocal = c.GatewayFeeLocal
	s.CurrencyFee = c.CurrencyFee
	s.Metadata = c.Metadata
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt
	s.ChargedbackAt = c.ChargedbackAt
	s.RefundedAt = c.RefundedAt
	s.AuthorizedAt = c.AuthorizedAt
	s.CapturedAt = c.CapturedAt
	s.VoidedAt = c.VoidedAt
	s.ThreeDS = c.ThreeDS
	s.CvcCheck = c.CvcCheck
	s.AvsCheck = c.AvsCheck
	s.InitialSchemeTransactionID = c.InitialSchemeTransactionID
	s.SchemeID = c.SchemeID
	s.PaymentType = c.PaymentType
	s.Eci = c.Eci
	s.NativeApm = c.NativeApm
	s.ExternalDetails = c.ExternalDetails

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
	return s.FetchRefundsWithContext(context.Background(), options...)
}

// FetchRefunds allows you to get the transaction's refunds., passes the provided context to the request
func (s Transaction) FetchRefundsWithContext(ctx context.Context, options ...TransactionFetchRefundsParameters) (*Iterator, error) {
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

	path := "/transactions/" + url.QueryEscape(*s.ID) + "/refunds"

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
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
		hasMorePrev: false,
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
	return s.FindRefundWithContext(context.Background(), refundID, options...)
}

// FindRefund allows you to find a transaction's refund by its ID., passes the provided context to the request
func (s Transaction) FindRefundWithContext(ctx context.Context, refundID string, options ...TransactionFindRefundParameters) (*Refund, error) {
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

	path := "/transactions/" + url.QueryEscape(*s.ID) + "/refunds/" + url.QueryEscape(refundID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
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
	return s.AllWithContext(context.Background(), options...)
}

// All allows you to get all the transactions., passes the provided context to the request
func (s Transaction) AllWithContext(ctx context.Context, options ...TransactionAllParameters) (*Iterator, error) {
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

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
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
		hasMorePrev: false,
	}
	return transactionsIterator, nil
}

// TransactionListParameters is the structure representing the
// additional parameters used to call Transaction.List
type TransactionListParameters struct {
	*Options
	*Transaction
}

// List allows you to get full transactions data for specified list of ids.
func (s Transaction) List(options ...TransactionListParameters) (*Iterator, error) {
	return s.ListWithContext(context.Background(), options...)
}

// List allows you to get full transactions data for specified list of ids., passes the provided context to the request
func (s Transaction) ListWithContext(ctx context.Context, options ...TransactionListParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewTransaction() method to create a new Transaction object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := TransactionListParameters{}
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

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
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
		hasMorePrev: false,
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
	return s.FindWithContext(context.Background(), transactionID, options...)
}

// Find allows you to find a transaction by its ID., passes the provided context to the request
func (s Transaction) FindWithContext(ctx context.Context, transactionID string, options ...TransactionFindParameters) (*Transaction, error) {
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

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
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
