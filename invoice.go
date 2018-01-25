package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v4/errors"
)

// Invoice represents the Invoice API object
type Invoice struct {
	// ID is the iD of the invoice
	ID *string `json:"id,omitempty"`
	// Project is the project to which the invoice belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the invoice belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Transaction is the transaction generated by the invoice
	Transaction *Transaction `json:"transaction,omitempty"`
	// TransactionID is the iD of the transaction generated by the invoice
	TransactionID *string `json:"transaction_id,omitempty"`
	// Customer is the customer linked to the invoice, if any
	Customer *Customer `json:"customer,omitempty"`
	// CustomerID is the iD of the customer linked to the invoice, if any
	CustomerID *string `json:"customer_id,omitempty"`
	// Subscription is the subscription to which the invoice is linked to, if any
	Subscription *Subscription `json:"subscription,omitempty"`
	// SubscriptionID is the iD of the subscription to which the invoice is linked to, if any
	SubscriptionID *string `json:"subscription_id,omitempty"`
	// Token is the token used to pay the invoice, if any
	Token *Token `json:"token,omitempty"`
	// TokenID is the iD of the token used to pay the invoice, if any
	TokenID *string `json:"token_id,omitempty"`
	// Details is the details of the invoice
	Details *[]*InvoiceDetail `json:"details,omitempty"`
	// URL is the uRL to which you may redirect your customer to proceed with the payment
	URL *string `json:"url,omitempty"`
	// Name is the name of the invoice
	Name *string `json:"name,omitempty"`
	// Amount is the amount to be paid
	Amount *string `json:"amount,omitempty"`
	// Currency is the currency of the invoice
	Currency *string `json:"currency,omitempty"`
	// StatementDescriptor is the statement to be shown on the bank statement of your customer
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// StatementDescriptorPhone is the support phone number shown on the customer's bank statement
	StatementDescriptorPhone *string `json:"statement_descriptor_phone,omitempty"`
	// StatementDescriptorCity is the city shown on the customer's bank statement
	StatementDescriptorCity *string `json:"statement_descriptor_city,omitempty"`
	// StatementDescriptorCompany is the your company name shown on the customer's bank statement
	StatementDescriptorCompany *string `json:"statement_descriptor_company,omitempty"`
	// StatementDescriptorURL is the uRL shown on the customer's bank statement
	StatementDescriptorURL *string `json:"statement_descriptor_url,omitempty"`
	// Metadata is the metadata related to the invoice, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// ReturnURL is the uRL where the customer will be redirected upon payment
	ReturnURL *string `json:"return_url,omitempty"`
	// CancelURL is the uRL where the customer will be redirected if the payment was canceled
	CancelURL *string `json:"cancel_url,omitempty"`
	// Sandbox is the define whether or not the invoice is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the invoice was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Invoice) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Invoice object and its
// children
func (s *Invoice) SetClient(c *ProcessOut) *Invoice {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Transaction != nil {
		s.Transaction.SetClient(c)
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

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Invoice) Prefill(c *Invoice) *Invoice {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Transaction = c.Transaction
	s.TransactionID = c.TransactionID
	s.Customer = c.Customer
	s.CustomerID = c.CustomerID
	s.Subscription = c.Subscription
	s.SubscriptionID = c.SubscriptionID
	s.Token = c.Token
	s.TokenID = c.TokenID
	s.Details = c.Details
	s.URL = c.URL
	s.Name = c.Name
	s.Amount = c.Amount
	s.Currency = c.Currency
	s.StatementDescriptor = c.StatementDescriptor
	s.StatementDescriptorPhone = c.StatementDescriptorPhone
	s.StatementDescriptorCity = c.StatementDescriptorCity
	s.StatementDescriptorCompany = c.StatementDescriptorCompany
	s.StatementDescriptorURL = c.StatementDescriptorURL
	s.Metadata = c.Metadata
	s.ReturnURL = c.ReturnURL
	s.CancelURL = c.CancelURL
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// InvoiceAuthorizeParameters is the structure representing the
// additional parameters used to call Invoice.Authorize
type InvoiceAuthorizeParameters struct {
	*Options
	*Invoice
	Synchronous                       interface{} `json:"synchronous"`
	PrioritizedGatewayConfigurationID interface{} `json:"prioritized_gateway_configuration_id"`
}

// Authorize allows you to authorize the invoice using the given source (customer or token)
func (s Invoice) Authorize(source string, options ...InvoiceAuthorizeParameters) (*Transaction, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceAuthorizeParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		Transaction *Transaction `json:"transaction"`
		HasMore     bool         `json:"has_more"`
		Success     bool         `json:"success"`
		Message     string       `json:"message"`
		Code        string       `json:"error_type"`
	}

	data := struct {
		*Options
		Synchronous                       interface{} `json:"synchronous"`
		PrioritizedGatewayConfigurationID interface{} `json:"prioritized_gateway_configuration_id"`
		Source                            interface{} `json:"source"`
	}{
		Options:                           opt.Options,
		Synchronous:                       opt.Synchronous,
		PrioritizedGatewayConfigurationID: opt.PrioritizedGatewayConfigurationID,
		Source: source,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/invoices/" + url.QueryEscape(*s.ID) + "/authorize"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

// InvoiceCaptureParameters is the structure representing the
// additional parameters used to call Invoice.Capture
type InvoiceCaptureParameters struct {
	*Options
	*Invoice
	AuthorizeOnly                     interface{} `json:"authorize_only"`
	Synchronous                       interface{} `json:"synchronous"`
	PrioritizedGatewayConfigurationID interface{} `json:"prioritized_gateway_configuration_id"`
}

// Capture allows you to capture the invoice using the given source (customer or token)
func (s Invoice) Capture(source string, options ...InvoiceCaptureParameters) (*Transaction, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceCaptureParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		Transaction *Transaction `json:"transaction"`
		HasMore     bool         `json:"has_more"`
		Success     bool         `json:"success"`
		Message     string       `json:"message"`
		Code        string       `json:"error_type"`
	}

	data := struct {
		*Options
		AuthorizeOnly                     interface{} `json:"authorize_only"`
		Synchronous                       interface{} `json:"synchronous"`
		PrioritizedGatewayConfigurationID interface{} `json:"prioritized_gateway_configuration_id"`
		Source                            interface{} `json:"source"`
	}{
		Options:                           opt.Options,
		AuthorizeOnly:                     opt.AuthorizeOnly,
		Synchronous:                       opt.Synchronous,
		PrioritizedGatewayConfigurationID: opt.PrioritizedGatewayConfigurationID,
		Source: source,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/invoices/" + url.QueryEscape(*s.ID) + "/capture"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

// InvoiceFetchCustomerParameters is the structure representing the
// additional parameters used to call Invoice.FetchCustomer
type InvoiceFetchCustomerParameters struct {
	*Options
	*Invoice
}

// FetchCustomer allows you to get the customer linked to the invoice.
func (s Invoice) FetchCustomer(options ...InvoiceFetchCustomerParameters) (*Customer, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceFetchCustomerParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		Customer *Customer `json:"customer"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
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

	path := "/invoices/" + url.QueryEscape(*s.ID) + "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	payload.Customer.SetClient(s.client)
	return payload.Customer, nil
}

// InvoiceAssignCustomerParameters is the structure representing the
// additional parameters used to call Invoice.AssignCustomer
type InvoiceAssignCustomerParameters struct {
	*Options
	*Invoice
}

// AssignCustomer allows you to assign a customer to the invoice.
func (s Invoice) AssignCustomer(customerID string, options ...InvoiceAssignCustomerParameters) (*Customer, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceAssignCustomerParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		Customer *Customer `json:"customer"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
	}

	data := struct {
		*Options
		CustomerID interface{} `json:"customer_id"`
	}{
		Options:    opt.Options,
		CustomerID: customerID,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/invoices/" + url.QueryEscape(*s.ID) + "/customers"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	payload.Customer.SetClient(s.client)
	return payload.Customer, nil
}

// InvoiceInitiateThreeDSParameters is the structure representing the
// additional parameters used to call Invoice.InitiateThreeDS
type InvoiceInitiateThreeDSParameters struct {
	*Options
	*Invoice
}

// InitiateThreeDS allows you to initiate a 3-D Secure authentication
func (s Invoice) InitiateThreeDS(source string, options ...InvoiceInitiateThreeDSParameters) (*CustomerAction, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceInitiateThreeDSParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		CustomerAction *CustomerAction `json:"customer_action"`
		HasMore        bool            `json:"has_more"`
		Success        bool            `json:"success"`
		Message        string          `json:"message"`
		Code           string          `json:"error_type"`
	}

	data := struct {
		*Options
		Source interface{} `json:"source"`
	}{
		Options: opt.Options,
		Source:  source,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/invoices/" + url.QueryEscape(*s.ID) + "/three-d-s"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	payload.CustomerAction.SetClient(s.client)
	return payload.CustomerAction, nil
}

// InvoiceFetchTransactionParameters is the structure representing the
// additional parameters used to call Invoice.FetchTransaction
type InvoiceFetchTransactionParameters struct {
	*Options
	*Invoice
}

// FetchTransaction allows you to get the transaction of the invoice.
func (s Invoice) FetchTransaction(options ...InvoiceFetchTransactionParameters) (*Transaction, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceFetchTransactionParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

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

	path := "/invoices/" + url.QueryEscape(*s.ID) + "/transactions"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

// InvoiceVoidParameters is the structure representing the
// additional parameters used to call Invoice.Void
type InvoiceVoidParameters struct {
	*Options
	*Invoice
}

// Void allows you to void the invoice
func (s Invoice) Void(options ...InvoiceVoidParameters) (*Transaction, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceVoidParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

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

	path := "/invoices/" + url.QueryEscape(*s.ID) + "/void"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

// InvoiceAllParameters is the structure representing the
// additional parameters used to call Invoice.All
type InvoiceAllParameters struct {
	*Options
	*Invoice
}

// All allows you to get all the invoices.
func (s Invoice) All(options ...InvoiceAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		Invoices []*Invoice `json:"invoices"`

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

	path := "/invoices"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	invoicesList := []Identifiable{}
	for _, o := range payload.Invoices {
		invoicesList = append(invoicesList, o.SetClient(s.client))
	}
	invoicesIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    invoicesList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"invoices"`
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
	return invoicesIterator, nil
}

// InvoiceCreateParameters is the structure representing the
// additional parameters used to call Invoice.Create
type InvoiceCreateParameters struct {
	*Options
	*Invoice
}

// Create allows you to create a new invoice.
func (s Invoice) Create(options ...InvoiceCreateParameters) (*Invoice, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		Invoice *Invoice `json:"invoice"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
	}

	data := struct {
		*Options
		CustomerID                 interface{} `json:"customer_id"`
		Name                       interface{} `json:"name"`
		Amount                     interface{} `json:"amount"`
		Currency                   interface{} `json:"currency"`
		Metadata                   interface{} `json:"metadata"`
		Details                    interface{} `json:"details"`
		StatementDescriptor        interface{} `json:"statement_descriptor"`
		StatementDescriptorPhone   interface{} `json:"statement_descriptor_phone"`
		StatementDescriptorCity    interface{} `json:"statement_descriptor_city"`
		StatementDescriptorCompany interface{} `json:"statement_descriptor_company"`
		StatementDescriptorURL     interface{} `json:"statement_descriptor_url"`
		ReturnURL                  interface{} `json:"return_url"`
		CancelURL                  interface{} `json:"cancel_url"`
	}{
		Options:                    opt.Options,
		CustomerID:                 s.CustomerID,
		Name:                       s.Name,
		Amount:                     s.Amount,
		Currency:                   s.Currency,
		Metadata:                   s.Metadata,
		Details:                    s.Details,
		StatementDescriptor:        s.StatementDescriptor,
		StatementDescriptorPhone:   s.StatementDescriptorPhone,
		StatementDescriptorCity:    s.StatementDescriptorCity,
		StatementDescriptorCompany: s.StatementDescriptorCompany,
		StatementDescriptorURL:     s.StatementDescriptorURL,
		ReturnURL:                  s.ReturnURL,
		CancelURL:                  s.CancelURL,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/invoices"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	payload.Invoice.SetClient(s.client)
	return payload.Invoice, nil
}

// InvoiceFindParameters is the structure representing the
// additional parameters used to call Invoice.Find
type InvoiceFindParameters struct {
	*Options
	*Invoice
}

// Find allows you to find an invoice by its ID.
func (s Invoice) Find(invoiceID string, options ...InvoiceFindParameters) (*Invoice, error) {
	if s.client == nil {
		panic("Please use the client.NewInvoice() method to create a new Invoice object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := InvoiceFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Invoice)

	type Response struct {
		Invoice *Invoice `json:"invoice"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/invoices/" + url.QueryEscape(invoiceID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	payload.Invoice.SetClient(s.client)
	return payload.Invoice, nil
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
