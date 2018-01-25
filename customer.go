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

// Customer represents the Customer API object
type Customer struct {
	// ID is the iD of the customer
	ID *string `json:"id,omitempty"`
	// Project is the project to which the customer belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the customer belongs
	ProjectID *string `json:"project_id,omitempty"`
	// DefaultToken is the default token of the customer
	DefaultToken *Token `json:"default_token,omitempty"`
	// DefaultTokenID is the iD of the default token of the customer
	DefaultTokenID *string `json:"default_token_id,omitempty"`
	// Tokens is the list of the customer tokens
	Tokens *[]*Token `json:"tokens,omitempty"`
	// Subscriptions is the list of the customer subscriptions
	Subscriptions *[]*Subscription `json:"subscriptions,omitempty"`
	// Transactions is the list of the customer transactions
	Transactions *[]*Transaction `json:"transactions,omitempty"`
	// Balance is the customer balance. Can be positive or negative
	Balance *string `json:"balance,omitempty"`
	// Currency is the currency of the customer balance. Once the currency is set it cannot be modified
	Currency *string `json:"currency,omitempty"`
	// Email is the email of the customer
	Email *string `json:"email,omitempty"`
	// FirstName is the first name of the customer
	FirstName *string `json:"first_name,omitempty"`
	// LastName is the last name of the customer
	LastName *string `json:"last_name,omitempty"`
	// Address1 is the address of the customer
	Address1 *string `json:"address1,omitempty"`
	// Address2 is the secondary address of the customer
	Address2 *string `json:"address2,omitempty"`
	// City is the city of the customer
	City *string `json:"city,omitempty"`
	// State is the state of the customer
	State *string `json:"state,omitempty"`
	// Zip is the zIP code of the customer
	Zip *string `json:"zip,omitempty"`
	// CountryCode is the country code of the customer (ISO-3166, 2 characters format)
	CountryCode *string `json:"country_code,omitempty"`
	// TransactionsCount is the number of transactions processed by the customer
	TransactionsCount *int `json:"transactions_count,omitempty"`
	// SubscriptionsCount is the number of active subscriptions linked to the customer
	SubscriptionsCount *int `json:"subscriptions_count,omitempty"`
	// MrrLocal is the mRR provided by the customer, converted to the currency of the Project
	MrrLocal *float64 `json:"mrr_local,omitempty"`
	// TotalRevenueLocal is the total revenue provided by the customer, converted to the currency of the Project
	TotalRevenueLocal *float64 `json:"total_revenue_local,omitempty"`
	// Metadata is the metadata related to the customer, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Sandbox is the define whether or not the customer is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the customer was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Customer) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Customer object and its
// children
func (s *Customer) SetClient(c *ProcessOut) *Customer {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.DefaultToken != nil {
		s.DefaultToken.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Customer) Prefill(c *Customer) *Customer {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.DefaultToken = c.DefaultToken
	s.DefaultTokenID = c.DefaultTokenID
	s.Tokens = c.Tokens
	s.Subscriptions = c.Subscriptions
	s.Transactions = c.Transactions
	s.Balance = c.Balance
	s.Currency = c.Currency
	s.Email = c.Email
	s.FirstName = c.FirstName
	s.LastName = c.LastName
	s.Address1 = c.Address1
	s.Address2 = c.Address2
	s.City = c.City
	s.State = c.State
	s.Zip = c.Zip
	s.CountryCode = c.CountryCode
	s.TransactionsCount = c.TransactionsCount
	s.SubscriptionsCount = c.SubscriptionsCount
	s.MrrLocal = c.MrrLocal
	s.TotalRevenueLocal = c.TotalRevenueLocal
	s.Metadata = c.Metadata
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// CustomerFetchSubscriptionsParameters is the structure representing the
// additional parameters used to call Customer.FetchSubscriptions
type CustomerFetchSubscriptionsParameters struct {
	*Options
	*Customer
}

// FetchSubscriptions allows you to get the subscriptions belonging to the customer.
func (s Customer) FetchSubscriptions(options ...CustomerFetchSubscriptionsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerFetchSubscriptionsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
		Subscriptions []*Subscription `json:"subscriptions"`

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

	path := "/customers/" + url.QueryEscape(*s.ID) + "/subscriptions"

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

	subscriptionsList := []Identifiable{}
	for _, o := range payload.Subscriptions {
		subscriptionsList = append(subscriptionsList, o.SetClient(s.client))
	}
	subscriptionsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    subscriptionsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"subscriptions"`
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
	return subscriptionsIterator, nil
}

// CustomerFetchTokensParameters is the structure representing the
// additional parameters used to call Customer.FetchTokens
type CustomerFetchTokensParameters struct {
	*Options
	*Customer
}

// FetchTokens allows you to get the customer's tokens.
func (s Customer) FetchTokens(options ...CustomerFetchTokensParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerFetchTokensParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
		Tokens []*Token `json:"tokens"`

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

	path := "/customers/" + url.QueryEscape(*s.ID) + "/tokens"

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

	tokensList := []Identifiable{}
	for _, o := range payload.Tokens {
		tokensList = append(tokensList, o.SetClient(s.client))
	}
	tokensIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    tokensList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"tokens"`
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
	return tokensIterator, nil
}

// CustomerFindTokenParameters is the structure representing the
// additional parameters used to call Customer.FindToken
type CustomerFindTokenParameters struct {
	*Options
	*Customer
}

// FindToken allows you to find a customer's token by its ID.
func (s Customer) FindToken(tokenID string, options ...CustomerFindTokenParameters) (*Token, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerFindTokenParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
		Token   *Token `json:"token"`
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

	path := "/customers/" + url.QueryEscape(*s.ID) + "/tokens/" + url.QueryEscape(tokenID) + ""

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

	payload.Token.SetClient(s.client)
	return payload.Token, nil
}

// CustomerDeleteTokenParameters is the structure representing the
// additional parameters used to call Customer.DeleteToken
type CustomerDeleteTokenParameters struct {
	*Options
	*Customer
}

// DeleteToken allows you to delete a customer's token by its ID.
func (s Customer) DeleteToken(tokenID string, options ...CustomerDeleteTokenParameters) error {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerDeleteTokenParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
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
		return errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(*s.ID) + "/tokens/" + url.QueryEscape(tokenID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
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

// CustomerFetchTransactionsParameters is the structure representing the
// additional parameters used to call Customer.FetchTransactions
type CustomerFetchTransactionsParameters struct {
	*Options
	*Customer
}

// FetchTransactions allows you to get the transactions belonging to the customer.
func (s Customer) FetchTransactions(options ...CustomerFetchTransactionsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerFetchTransactionsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

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

	path := "/customers/" + url.QueryEscape(*s.ID) + "/transactions"

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

// CustomerAllParameters is the structure representing the
// additional parameters used to call Customer.All
type CustomerAllParameters struct {
	*Options
	*Customer
}

// All allows you to get all the customers.
func (s Customer) All(options ...CustomerAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
		Customers []*Customer `json:"customers"`

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

	path := "/customers"

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

	customersList := []Identifiable{}
	for _, o := range payload.Customers {
		customersList = append(customersList, o.SetClient(s.client))
	}
	customersIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    customersList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"customers"`
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
	return customersIterator, nil
}

// CustomerCreateParameters is the structure representing the
// additional parameters used to call Customer.Create
type CustomerCreateParameters struct {
	*Options
	*Customer
}

// Create allows you to create a new customer.
func (s Customer) Create(options ...CustomerCreateParameters) (*Customer, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
		Customer *Customer `json:"customer"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
	}

	data := struct {
		*Options
		Balance     interface{} `json:"balance"`
		Currency    interface{} `json:"currency"`
		Email       interface{} `json:"email"`
		FirstName   interface{} `json:"first_name"`
		LastName    interface{} `json:"last_name"`
		Address1    interface{} `json:"address1"`
		Address2    interface{} `json:"address2"`
		City        interface{} `json:"city"`
		State       interface{} `json:"state"`
		Zip         interface{} `json:"zip"`
		CountryCode interface{} `json:"country_code"`
		Metadata    interface{} `json:"metadata"`
	}{
		Options:     opt.Options,
		Balance:     s.Balance,
		Currency:    s.Currency,
		Email:       s.Email,
		FirstName:   s.FirstName,
		LastName:    s.LastName,
		Address1:    s.Address1,
		Address2:    s.Address2,
		City:        s.City,
		State:       s.State,
		Zip:         s.Zip,
		CountryCode: s.CountryCode,
		Metadata:    s.Metadata,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/customers"

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

// CustomerFindParameters is the structure representing the
// additional parameters used to call Customer.Find
type CustomerFindParameters struct {
	*Options
	*Customer
}

// Find allows you to find a customer by its ID.
func (s Customer) Find(customerID string, options ...CustomerFindParameters) (*Customer, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

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

	path := "/customers/" + url.QueryEscape(customerID) + ""

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

// CustomerSaveParameters is the structure representing the
// additional parameters used to call Customer.Save
type CustomerSaveParameters struct {
	*Options
	*Customer
}

// Save allows you to save the updated customer attributes.
func (s Customer) Save(options ...CustomerSaveParameters) (*Customer, error) {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
		Customer *Customer `json:"customer"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
	}

	data := struct {
		*Options
		Balance        interface{} `json:"balance"`
		DefaultTokenID interface{} `json:"default_token_id"`
		Email          interface{} `json:"email"`
		FirstName      interface{} `json:"first_name"`
		LastName       interface{} `json:"last_name"`
		Address1       interface{} `json:"address1"`
		Address2       interface{} `json:"address2"`
		City           interface{} `json:"city"`
		State          interface{} `json:"state"`
		Zip            interface{} `json:"zip"`
		CountryCode    interface{} `json:"country_code"`
		Metadata       interface{} `json:"metadata"`
	}{
		Options:        opt.Options,
		Balance:        s.Balance,
		DefaultTokenID: s.DefaultTokenID,
		Email:          s.Email,
		FirstName:      s.FirstName,
		LastName:       s.LastName,
		Address1:       s.Address1,
		Address2:       s.Address2,
		City:           s.City,
		State:          s.State,
		Zip:            s.Zip,
		CountryCode:    s.CountryCode,
		Metadata:       s.Metadata,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

// CustomerDeleteParameters is the structure representing the
// additional parameters used to call Customer.Delete
type CustomerDeleteParameters struct {
	*Options
	*Customer
}

// Delete allows you to delete the customer.
func (s Customer) Delete(options ...CustomerDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CustomerDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Customer)

	type Response struct {
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
		return errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
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

// dummyCustomer is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCustomer() {
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
