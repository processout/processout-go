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

// Refund represents the Refund API object
type Refund struct {
	// ID is the iD of the refund
	ID *string `json:"id,omitempty"`
	// Transaction is the transaction to which the refund is applied
	Transaction *Transaction `json:"transaction,omitempty"`
	// TransactionID is the iD of the transaction to which the refund is applied
	TransactionID *string `json:"transaction_id,omitempty"`
	// Amount is the amount to be refunded. Must not be greater than the amount still available on the transaction
	Amount *string `json:"amount,omitempty"`
	// Reason is the reason for the refund. Either customer_request, duplicate or fraud
	Reason *string `json:"reason,omitempty"`
	// Information is the custom details regarding the refund
	Information *string `json:"information,omitempty"`
	// HasFailed is the true if the refund was asynchronously failed, false otherwise
	HasFailed *bool `json:"has_failed,omitempty"`
	// Metadata is the metadata related to the refund, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Sandbox is the define whether or not the refund is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the refund was done
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// InvoiceDetailIds is the list of invoice details ids to refund
	InvoiceDetailIds *[]string `json:"invoice_detail_ids,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Refund) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Refund object and its
// children
func (s *Refund) SetClient(c *ProcessOut) *Refund {
	if s == nil {
		return s
	}
	s.client = c
	if s.Transaction != nil {
		s.Transaction.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Refund) Prefill(c *Refund) *Refund {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Transaction = c.Transaction
	s.TransactionID = c.TransactionID
	s.Amount = c.Amount
	s.Reason = c.Reason
	s.Information = c.Information
	s.HasFailed = c.HasFailed
	s.Metadata = c.Metadata
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt
	s.InvoiceDetailIds = c.InvoiceDetailIds

	return s
}

// RefundCreateForInvoiceParameters is the structure representing the
// additional parameters used to call Refund.CreateForInvoice
type RefundCreateForInvoiceParameters struct {
	*Options
	*Refund
	Metadata interface{} `json:"metadata"`
}

// CreateForInvoice allows you to create a refund for an invoice.
func (s Refund) CreateForInvoice(invoiceID string, options ...RefundCreateForInvoiceParameters) error {
	if s.client == nil {
		panic("Please use the client.NewRefund() method to create a new Refund object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := RefundCreateForInvoiceParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Refund)

	type Response struct {
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Amount           interface{} `json:"amount"`
		Reason           interface{} `json:"reason"`
		Information      interface{} `json:"information"`
		InvoiceDetailIds interface{} `json:"invoice_detail_ids"`
		Metadata         interface{} `json:"metadata"`
	}{
		Options:          opt.Options,
		Amount:           s.Amount,
		Reason:           s.Reason,
		Information:      s.Information,
		InvoiceDetailIds: s.InvoiceDetailIds,
		Metadata:         opt.Metadata,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "/invoices/" + url.QueryEscape(invoiceID) + "/refunds"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return erri
	}

	return nil
}

// RefundFetchTransactionRefundsParameters is the structure representing the
// additional parameters used to call Refund.FetchTransactionRefunds
type RefundFetchTransactionRefundsParameters struct {
	*Options
	*Refund
}

// FetchTransactionRefunds allows you to get the transaction's refunds.
func (s Refund) FetchTransactionRefunds(transactionID string, options ...RefundFetchTransactionRefundsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewRefund() method to create a new Refund object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := RefundFetchTransactionRefundsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Refund)

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

	path := "/transactions/" + url.QueryEscape(transactionID) + "/refunds"

	req, err := http.NewRequest(
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

// RefundFindParameters is the structure representing the
// additional parameters used to call Refund.Find
type RefundFindParameters struct {
	*Options
	*Refund
}

// Find allows you to find a transaction's refund by its ID.
func (s Refund) Find(transactionID, refundID string, options ...RefundFindParameters) (*Refund, error) {
	if s.client == nil {
		panic("Please use the client.NewRefund() method to create a new Refund object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := RefundFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Refund)

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

	path := "/transactions/" + url.QueryEscape(transactionID) + "/refunds/" + url.QueryEscape(refundID) + ""

	req, err := http.NewRequest(
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

// RefundCreateParameters is the structure representing the
// additional parameters used to call Refund.Create
type RefundCreateParameters struct {
	*Options
	*Refund
	Metadata interface{} `json:"metadata"`
}

// Create allows you to create a refund for a transaction.
func (s Refund) Create(options ...RefundCreateParameters) error {
	if s.client == nil {
		panic("Please use the client.NewRefund() method to create a new Refund object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := RefundCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Refund)

	type Response struct {
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Amount           interface{} `json:"amount"`
		Reason           interface{} `json:"reason"`
		Information      interface{} `json:"information"`
		InvoiceDetailIds interface{} `json:"invoice_detail_ids"`
		Metadata         interface{} `json:"metadata"`
	}{
		Options:          opt.Options,
		Amount:           s.Amount,
		Reason:           s.Reason,
		Information:      s.Information,
		InvoiceDetailIds: s.InvoiceDetailIds,
		Metadata:         opt.Metadata,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "/transactions/" + url.QueryEscape(*s.TransactionID) + "/refunds"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return erri
	}

	return nil
}

// dummyRefund is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyRefund() {
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
