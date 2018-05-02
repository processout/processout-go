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

// Payout represents the Payout API object
type Payout struct {
	// ID is the iD of the payout
	ID *string `json:"id,omitempty"`
	// Project is the project to which the payout belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the payout belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Status is the status of the payout
	Status *string `json:"status,omitempty"`
	// Amount is the amount of the payout
	Amount *string `json:"amount,omitempty"`
	// Currency is the currency of the payout
	Currency *string `json:"currency,omitempty"`
	// Metadata is the metadata related to the payout, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// BankName is the name of the bank to which the payout was issued, if available
	BankName *string `json:"bank_name,omitempty"`
	// BankSummary is the summary of the bank to which the payout was issued, if available
	BankSummary *string `json:"bank_summary,omitempty"`
	// SalesTransactions is the number of completed transactions linked to the payout, if available
	SalesTransactions *int `json:"sales_transactions,omitempty"`
	// SalesVolume is the volume of completed transactions linked to the payout, if available
	SalesVolume *string `json:"sales_volume,omitempty"`
	// RefundsTransactions is the number of refunded transactions linked to the payout, if available
	RefundsTransactions *int `json:"refunds_transactions,omitempty"`
	// RefundsVolume is the volume of refunded transactions linked to the payout, if available
	RefundsVolume *string `json:"refunds_volume,omitempty"`
	// ChargebacksTransactions is the number of chargebacked transactions linked to the payout, if available
	ChargebacksTransactions *int `json:"chargebacks_transactions,omitempty"`
	// ChargebacksVolume is the volume of chargebacked transactions linked to the payout, if available
	ChargebacksVolume *string `json:"chargebacks_volume,omitempty"`
	// Fees is the fees linked to the payout, if available
	Fees *string `json:"fees,omitempty"`
	// Adjustments is the adjustments linked to the payout, if available
	Adjustments *string `json:"adjustments,omitempty"`
	// Reserve is the reserve kept on the payout, if available
	Reserve *string `json:"reserve,omitempty"`
	// CreatedAt is the date at which the payout was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Payout) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Payout object and its
// children
func (s *Payout) SetClient(c *ProcessOut) *Payout {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Payout) Prefill(c *Payout) *Payout {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Status = c.Status
	s.Amount = c.Amount
	s.Currency = c.Currency
	s.Metadata = c.Metadata
	s.BankName = c.BankName
	s.BankSummary = c.BankSummary
	s.SalesTransactions = c.SalesTransactions
	s.SalesVolume = c.SalesVolume
	s.RefundsTransactions = c.RefundsTransactions
	s.RefundsVolume = c.RefundsVolume
	s.ChargebacksTransactions = c.ChargebacksTransactions
	s.ChargebacksVolume = c.ChargebacksVolume
	s.Fees = c.Fees
	s.Adjustments = c.Adjustments
	s.Reserve = c.Reserve
	s.CreatedAt = c.CreatedAt

	return s
}

// PayoutFetchItemsParameters is the structure representing the
// additional parameters used to call Payout.FetchItems
type PayoutFetchItemsParameters struct {
	*Options
	*Payout
}

// FetchItems allows you to get all the items linked to the payout.
func (s Payout) FetchItems(options ...PayoutFetchItemsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewPayout() method to create a new Payout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PayoutFetchItemsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Payout)

	type Response struct {
		Items []*PayoutItem `json:"items"`

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

	path := "/payouts/" + url.QueryEscape(*s.ID) + "/items"

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

	itemsList := []Identifiable{}
	for _, o := range payload.Items {
		itemsList = append(itemsList, o.SetClient(s.client))
	}
	itemsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    itemsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"items"`
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
	return itemsIterator, nil
}

// PayoutAllParameters is the structure representing the
// additional parameters used to call Payout.All
type PayoutAllParameters struct {
	*Options
	*Payout
}

// All allows you to get all the payouts.
func (s Payout) All(options ...PayoutAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewPayout() method to create a new Payout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PayoutAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Payout)

	type Response struct {
		Payouts []*Payout `json:"payouts"`

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

	path := "/payouts"

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

	payoutsList := []Identifiable{}
	for _, o := range payload.Payouts {
		payoutsList = append(payoutsList, o.SetClient(s.client))
	}
	payoutsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    payoutsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"payouts"`
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
	return payoutsIterator, nil
}

// PayoutFindParameters is the structure representing the
// additional parameters used to call Payout.Find
type PayoutFindParameters struct {
	*Options
	*Payout
}

// Find allows you to find a payout by its ID.
func (s Payout) Find(payoutID string, options ...PayoutFindParameters) (*Payout, error) {
	if s.client == nil {
		panic("Please use the client.NewPayout() method to create a new Payout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := PayoutFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Payout)

	type Response struct {
		Payout  *Payout `json:"payout"`
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

	path := "/payouts/" + url.QueryEscape(payoutID) + ""

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

	payload.Payout.SetClient(s.client)
	return payload.Payout, nil
}

// dummyPayout is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPayout() {
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
