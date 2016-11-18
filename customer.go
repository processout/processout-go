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

// Customer represents the Customer API object
type Customer struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the customer
	ID string `json:"id"`
	// Project is the project to which the customer belongs
	Project *Project `json:"project"`
	// Email is the email of the customer
	Email string `json:"email"`
	// FirstName is the first name of the customer
	FirstName string `json:"first_name"`
	// LastName is the last name of the customer
	LastName string `json:"last_name"`
	// Address1 is the address of the customer
	Address1 string `json:"address1"`
	// Address2 is the secondary address of the customer
	Address2 string `json:"address2"`
	// City is the city of the customer
	City string `json:"city"`
	// State is the state of the customer
	State string `json:"state"`
	// Zip is the zIP code of the customer
	Zip string `json:"zip"`
	// CountryCode is the country code of the customer
	CountryCode string `json:"country_code"`
	// Balance is the customer balance. Can be positive or negative
	Balance string `json:"balance"`
	// Currency is the currency of the customer balance. Once the currency is set it cannot be modified
	Currency string `json:"currency"`
	// Metadata is the metadata related to the customer, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// HasPin is the wether the customer has a PIN set or not
	HasPin bool `json:"has_pin"`
	// Sandbox is the define whether or not the customer is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt is the date at which the customer was created
	CreatedAt *time.Time `json:"created_at"`
}

// FetchSubscriptions allows you to get the subscriptions belonging to the customer.
func (s Customer) FetchSubscriptions(options ...Options) ([]*Subscription, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Subscriptions []*Subscription `json:"subscriptions"`

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

	path := "/customers/" + url.QueryEscape(s.ID) + "/subscriptions"

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

	for _, o := range payload.Subscriptions {
		o.Client = s.Client
	}
	return payload.Subscriptions, nil
}

// FetchTokens allows you to get the customer's tokens.
func (s Customer) FetchTokens(options ...Options) ([]*Token, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Tokens []*Token `json:"tokens"`

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

	path := "/customers/" + url.QueryEscape(s.ID) + "/tokens"

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

	for _, o := range payload.Tokens {
		o.Client = s.Client
	}
	return payload.Tokens, nil
}

// FindToken allows you to find a customer's token by its ID.
func (s Customer) FindToken(tokenID string, options ...Options) (*Token, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Token   *Token `json:"token"`
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

	path := "/customers/" + url.QueryEscape(s.ID) + "/tokens/" + url.QueryEscape(tokenID) + ""

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

	payload.Token.Client = s.Client
	return payload.Token, nil
}

// FetchTransactions allows you to get the transactions belonging to the customer.
func (s Customer) FetchTransactions(options ...Options) ([]*Transaction, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
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

	path := "/customers/" + url.QueryEscape(s.ID) + "/transactions"

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
		o.Client = s.Client
	}
	return payload.Transactions, nil
}

// All allows you to get all the customers.
func (s Customer) All(options ...Options) ([]*Customer, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Customers []*Customer `json:"customers"`

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

	path := "/customers"

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

	for _, o := range payload.Customers {
		o.Client = s.Client
	}
	return payload.Customers, nil
}

// Create allows you to create a new customer.
func (s Customer) Create(options ...Options) (*Customer, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Customer *Customer `json:"customer"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"balance":      s.Balance,
		"currency":     s.Currency,
		"email":        s.Email,
		"first_name":   s.FirstName,
		"last_name":    s.LastName,
		"address1":     s.Address1,
		"address2":     s.Address2,
		"city":         s.City,
		"state":        s.State,
		"zip":          s.Zip,
		"country_code": s.CountryCode,
		"metadata":     s.Metadata,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
		"limit":        opt.Limit,
		"page":         opt.Page,
		"end_before":   opt.EndBefore,
		"start_after":  opt.StartAfter,
	})
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

	payload.Customer.Client = s.Client
	return payload.Customer, nil
}

// Find allows you to find a customer by its ID.
func (s Customer) Find(customerID string, options ...Options) (*Customer, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Customer *Customer `json:"customer"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
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

	path := "/customers/" + url.QueryEscape(customerID) + ""

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

	payload.Customer.Client = s.Client
	return payload.Customer, nil
}

// Save allows you to save the updated customer attributes.
func (s Customer) Save(options ...Options) (*Customer, error) {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Customer *Customer `json:"customer"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"balance":      s.Balance,
		"email":        s.Email,
		"first_name":   s.FirstName,
		"last_name":    s.LastName,
		"address1":     s.Address1,
		"address2":     s.Address2,
		"city":         s.City,
		"state":        s.State,
		"zip":          s.Zip,
		"country_code": s.CountryCode,
		"metadata":     s.Metadata,
		"expand":       opt.Expand,
		"filter":       opt.Filter,
		"limit":        opt.Limit,
		"page":         opt.Page,
		"end_before":   opt.EndBefore,
		"start_after":  opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	payload.Customer.Client = s.Client
	return payload.Customer, nil
}

// Delete allows you to delete the customer.
func (s Customer) Delete(options ...Options) error {
	if s.Client == nil {
		panic("Please use the client.NewCustomer() method to create a new Customer object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
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
		return errors.New(err, "", "")
	}

	path := "/customers/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
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
	}
	errors.New(nil, "", "")
}
